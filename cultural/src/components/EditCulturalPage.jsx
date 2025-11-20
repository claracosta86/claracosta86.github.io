import { useUser } from '../contexts/UserContext';
import { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import Header from './Layout/Header';
import Footer from './Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/create.css';

const EditCulturalPage = () => {
  const navigate = useNavigate();

  const { user } = useUser();
  const userID = user.userID;
  const userType = user.type;

  const {
    notifications,
    isNotificationModalOpen,
    setNotificationModalOpen,
    fetchNotifications,
    markNotificationsAsSeen,
  } = useNotifications(userID);

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

  const [selectedFile, setSelectedFile] = useState(null);

  const { id, culturalType: initialCulturalType } = useParams();
  const [culturalType, setCulturalType] = useState(initialCulturalType);

  const [error, setError] = useState('');

  const [culturalDetails, setCulturalDetails] = useState(null);

  useEffect(() => {
    const fetchCulturalDetails = async () => {
      try {
        console.log(`Fetching cultural details for culturalType: ${culturalType}, id: ${id}`);
        const culturalResponse = await fetch(
          `http://localhost:8080/culturais/${culturalType}/${id}`
        );
        const cultural = await culturalResponse.json();
        console.log('Cultural details received:', cultural);
        setCulturalDetails(cultural);
      } catch (error) {
        console.error('Erro ao buscar detalhes do cultural:', error);
      }
    };
    fetchCulturalDetails();
  }, [culturalType, id]);

  // Estado único para gerenciar todos os campos do formulário
  const [formData, setFormData] = useState({
    title: culturalDetails ? culturalDetails.title : '',
    description: culturalDetails ? culturalDetails.description : '',
    location: culturalDetails ? culturalDetails.location : '',
    price: culturalDetails ? culturalDetails.price : 'R$0,00',
    isAccessible: culturalDetails ? culturalDetails.isAccessible : false,
    startDate: culturalDetails ? culturalDetails.startDate : '',
    endDate: culturalDetails ? culturalDetails.endDate : '',
    durationHours: culturalDetails ? culturalDetails.durationHours : '',
    workingHours: culturalDetails ? culturalDetails.workingHours : '',
  });

  // Converte uma string de data para o formato 'YYYY-MM-DDTHH:MM'
  const formatDateForInput = (dateString) => {
    if (!dateString) return ''; // Retorna vazio se a data não existir

    const date = new Date(dateString);
    if (isNaN(date.getTime())) return ''; // Retorna vazio se a data for inválida

    // Pega os componentes da data
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Mês é 0-indexado
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    // Monta a string no formato correto
    return `${year}-${month}-${day}T${hours}:${minutes}`;
  };

  useEffect(() => {
    // Apenas preenche o formulário se culturalDetails não for nulo
    if (culturalDetails) {
      setFormData({
        title: culturalDetails.title || '',
        description: culturalDetails.description || '',
        location: culturalDetails.location || '',
        price: culturalDetails.price || 'R$0,00',
        isAccessible: culturalDetails.isAccessible || false,

        startDate: formatDateForInput(culturalDetails.event ? culturalDetails.event.startDate : ''),
        endDate: formatDateForInput(culturalDetails.event ? culturalDetails.event.endDate : ''),
        durationHours: culturalDetails.event ? culturalDetails.event.durationHours : '',

        workingHours: culturalDetails.touristAttraction
          ? culturalDetails.touristAttraction.workingHours
          : '',
        image: culturalDetails.image || '',
      });
    }
  }, [culturalDetails]);

  const handleTypeChange = (event) => {
    setCulturalType(event.target.value);
    setFormData({
      title: '',
      description: '',
      location: '',
      price: 'R$0,00',
      isAccessible: false,
      startDate: '',
      endDate: '',
      durationHours: '',
      workingHours: '',
    });
  };

  const handleInputChange = (event) => {
    const { name, value, type, checked } = event.target;
    const finalValue = type === 'checkbox' ? checked : value;
    setFormData((prevData) => ({ ...prevData, [name]: finalValue }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setError('');

    const submissionFormData = new FormData();

    if (selectedFile) {
      submissionFormData.append('image', selectedFile);
    }

    let finalPayload = {
      id: parseInt(id, 10),
      title: formData.title,
      type: culturalType,
      description: formData.description,
      price: formData.price,
      isAccessible: formData.isAccessible,
      location: formData.location,
      organizerID: userID,
      image: formData.image || (culturalDetails ? culturalDetails.image : ''),
    };

    if (culturalType === 'event') {
      finalPayload.event = {
        startDate: formData.startDate,
        endDate: formData.endDate,
        durationHours: formData.durationHours,
      };
    } else if (culturalType === 'tourist_attraction') {
      finalPayload.touristAttraction = {
        workingHours: formData.workingHours,
      };
    }

    submissionFormData.append('data', JSON.stringify(finalPayload));

    console.log('Enviando para a API:', JSON.stringify(finalPayload, null, 2));

    try {
      const response = await fetch('http://localhost:8080/culturais/', {
        method: 'PATCH',
        body: submissionFormData,
        credentials: 'include',
      });

      if (response.status === 204) {
        navigate(`/card/${culturalType}/${id}`);
      } else {
        const errorData = await response.json();
        console.error('Erro da API:', errorData);
        setError(errorData.message || 'Falha ao criar o item.');
        alert(`Falha ao criar o item: ${errorData.message}`);
      }
    } catch (err) {
      console.error('Erro de rede:', err);
      setError('Não foi possível atualizar seu cultural no momento, tente novamente mais tarde :(');
    }
  };

  const closeConfirmModal = () => {
    setConfirmModalOpen(false);
  };

  const handleConfirmLogout = () => {
    navigate('/');
  };

  return (
    <>
      <NotificationModal
        isOpen={isNotificationModalOpen}
        onClose={markNotificationsAsSeen}
        notifications={notifications}
        navigate={navigate}
        userID={userID}
        userType={userType}
      />
      <ConfirmModal
        isOpen={isConfirmModalOpen}
        onClose={closeConfirmModal}
        onConfirm={handleConfirmLogout}
      />

      <section className="screen" id="tela-home">
        <Header
          onLogoutClick={() => setConfirmModalOpen(true)}
          onNotificationClick={fetchNotifications}
        />

        <div className="create-box">
          <div className="header-title">
            <h2>Edite Seu Cultural</h2>
          </div>

          <form id="createCulturalForm" onSubmit={handleSubmit}>
            <fieldset className="selection-box">
              <legend>Selecione seu tipo de cultural:</legend>
              <div className="radio-option">
                <input
                  type="radio"
                  id="event"
                  name="culturalType"
                  value="event"
                  checked={culturalType === 'event'}
                  onChange={handleTypeChange}
                />
                <label htmlFor="event">Evento</label>
              </div>
              <div className="radio-option">
                <input
                  type="radio"
                  id="tourist-attraction"
                  name="culturalType"
                  value="tourist_attraction"
                  checked={culturalType === 'tourist_attraction'}
                  onChange={handleTypeChange}
                />
                <label htmlFor="tourist-attraction">Ponto Turístico</label>
              </div>
            </fieldset>

            {culturalType && (
              <>
                <label htmlFor="title" className="required">
                  Nome
                </label>
                <input
                  id="title"
                  name="title"
                  type="text"
                  placeholder="Nome do evento ou local"
                  value={formData.title}
                  onChange={handleInputChange}
                  required
                />

                <label htmlFor="location" className="required">
                  Endereço
                </label>
                <input
                  id="location"
                  name="location"
                  type="text"
                  placeholder="Rua, número, bairro..."
                  value={formData.location}
                  onChange={handleInputChange}
                  required
                />

                {culturalType === 'event' ? (
                  <>
                    <label htmlFor="start-date" className="required">
                      Data e Hora de Início
                    </label>
                    <input
                      type="datetime-local"
                      id="start-date"
                      name="startDate"
                      value={formData.startDate}
                      onChange={handleInputChange}
                      required
                    />

                    <label htmlFor="end-date" className="required">
                      Data e Hora de Fim
                    </label>
                    <input
                      type="datetime-local"
                      id="end-date"
                      name="endDate"
                      value={formData.endDate}
                      onChange={handleInputChange}
                      required
                    />

                    <label htmlFor="duration-hours" className="required">
                      Horário de Duração
                    </label>
                    <input
                      type="text"
                      id="duration-hours"
                      name="durationHours"
                      placeholder="HH:MM às HH:MM"
                      value={formData.durationHours}
                      onChange={handleInputChange}
                      required
                    />
                  </>
                ) : (
                  <>
                    <label className="required">Horário de Funcionamento</label>
                    <textarea
                      id="working-hours"
                      name="workingHours"
                      placeholder="ex:  Domingo	Fechado
                                        Segunda-feira	11:00–15:00
                                        Terça-feira	11:00–23:00
                                        Quarta-feira	11:00–23:00
                                        Quinta-feira	11:00–23:00
                                        Sexta-feira	11:00–23:00
                                        Sábado 24 horas"
                      value={formData.workingHours}
                      onChange={handleInputChange}
                      required
                    />
                  </>
                )}

                <label htmlFor="description">Descrição</label>
                <textarea
                  id="description"
                  name="description"
                  placeholder="Descreva seu cultural :)"
                  value={formData.description}
                  onChange={handleInputChange}
                />

                <label htmlFor="price">Preço</label>
                <input
                  type="text"
                  id="price"
                  name="price"
                  value={formData.price}
                  onChange={handleInputChange}
                  required
                />

                <label htmlFor="accessibility">Acessibilidade</label>
                <div className="checkbox-option accessibility-option">
                  <input
                    type="checkbox"
                    id="is-accessible"
                    name="isAccessible"
                    checked={formData.isAccessible}
                    onChange={handleInputChange}
                  />
                  <label htmlFor="is-accessible">Possui estrutura de acessibilidade</label>
                </div>

                <label htmlFor="image">Imagem do Cultural</label>
                <input
                  type="file"
                  id="image"
                  name="image"
                  onChange={(event) => setSelectedFile(event.target.files[0])}
                />

                {error && <span className="error">{error}</span>}
              </>
            )}
            <div className="button-container">
              <div className="button-container-row">
                <button type="button" className="btn" onClick={() => navigate(-1)}>
                  Cancelar
                </button>
                <button type="submit" className="btn">
                  Atualizar
                </button>
              </div>
            </div>
          </form>
        </div>
        <Footer userType={userType} />
      </section>
    </>
  );
};

export default EditCulturalPage;
