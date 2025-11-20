import { useUser } from '../contexts/UserContext';
import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import Header from './Layout/Header';
import Footer from './Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/create.css';

const CreateCulturalPage = () => {
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

  const [selectedFile, setSelectedFile] = useState(null);
  const [culturalType, setCulturalType] = useState('');
  const [error, setError] = useState('');

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

  // Estado único para gerenciar todos os campos do formulário
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    location: '',
    price: 'R$0,00',
    isAccessible: false,
    startDate: '',
    endDate: '',
  });

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

  const closeConfirmModal = () => {
    setConfirmModalOpen(false);
  };

  const handleConfirmLogout = () => {
    navigate('/');
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setError('');

    const submissionFormData = new FormData();

    if (selectedFile) {
      submissionFormData.append('image', selectedFile);
    }

    let finalPayload = {
      // Use as chaves exatas do seu `json tag` no Go
      title: formData.title,
      type: culturalType,
      description: formData.description,
      price: formData.price,
      isAccessible: formData.isAccessible,
      organizerID: userID,
      location: formData.location,
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
        method: 'POST',
        body: submissionFormData,
        credentials: 'include',
      });

      if (response.ok) {
        const result = await response.json();
        navigate(`/card/${result.type}/${result.id}`);
      } else {
        const errorData = await response.json();
        console.error('Erro da API:', errorData);
        setError(errorData.message || 'Falha ao criar o item.');
        alert(`Falha ao criar o item: ${errorData.message}`);
      }
    } catch (err) {
      console.error('Erro de rede:', err);
      setError('Não foi possível criar seu cultural no momento, tente novamente mais tarde :(');
    }
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
            <h2>Crie Seu Cultural</h2>
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
                <label htmlFor="tourist_attraction">Ponto Turístico</label>
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
                      Horário de Funcionamento
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
                    <label htmlFor="working-hours" className="required">
                      Horário de Funcionamento
                    </label>
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

                <p className="button-container">
                  <button type="submit" className="btn">
                    Criar
                  </button>
                </p>
              </>
            )}
          </form>
        </div>
        <Footer userType={userType} />
      </section>
    </>
  );
};

export default CreateCulturalPage;
