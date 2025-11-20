import { useUser } from '../contexts/UserContext';
import { useEffect, useState } from 'react';
import { Link, useParams, useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import './styles/create.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import userIcon from '../assets/user-icon.png';
import homeIcon from '../assets/home-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import searchIcon from '../assets/search-icon.png';

const EditCulturalPage = () => {
  const navigate = useNavigate();

  const { user } = useUser();
  const userID = user.userID;
  const userType = user.type;

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

  const [selectedFile, setSelectedFile] = useState(null);

   const { id, culturalType } = useParams();

  const [error, setError] = useState('');

  const [culturalDetails, setCulturalDetails] = useState(null);

  useEffect(() => {
    const fetchCulturalDetails = async () => {
      try {
        console.log(`Fetching cultural details for culturalType: ${culturalType}, id: ${id}`);
        const culturalResponse = await fetch(`http://localhost:8080/culturais/${culturalType}/${id}`);
        const cultural = await culturalResponse.json();
        console.log('Cultural details received:', cultural);
        setCulturalDetails(cultural);
      } catch (error) {
        console.error('Erro ao buscar detalhes do cultural:', error);
      }
    };
    fetchCulturalDetails();
  }, [culturalType]);

  // Estado único para gerenciar todos os campos do formulário
  const [formData, setFormData] = useState({
    title: culturalDetails ? culturalDetails.title : '',
    description: culturalDetails ? culturalDetails.description : '',
    location: culturalDetails ? culturalDetails.location : '',
    price: culturalDetails ? culturalDetails.price : 'R$0,00',
    is_accessible: culturalDetails ? culturalDetails.is_accessible : false,
    start_date: culturalDetails ? culturalDetails.start_date : '',
    end_date: culturalDetails ? culturalDetails.end_date : '',
    working_hours: culturalDetails ? culturalDetails.working_hours : '',
    open_days: culturalDetails ? culturalDetails.open_days : [],
    open_time: culturalDetails ? culturalDetails.open_time : '',
  });

  useEffect(() => {
    // Apenas preenche o formulário se culturalDetails não for nulo
    if (culturalDetails) {
        setFormData({
        title: culturalDetails.title || '',
        description: culturalDetails.description || '',
        location: culturalDetails.location || '',
        price: culturalDetails.price || 'R$0,00',
        is_accessible: culturalDetails.is_accessible || false,
        
        // Lida com dados aninhados de eventos
        start_date: formatDateForInput(culturalDetails.event ? culturalDetails.event.start_date : ''),
        end_date: formatDateForInput(culturalDetails.event ? culturalDetails.event.end_date : ''),
        working_hours: culturalDetails.event ? culturalDetails.event.working_hours : '',
        
        // Lida com dados aninhados de pontos turísticos
        // A API envia como string "Segunda, Terça", então transformamos em array
        open_days: culturalDetails.tourist_attraction ? culturalDetails.tourist_attraction.open_days.split(', ') : [],
        open_time: culturalDetails.tourist_attraction ? culturalDetails.tourist_attraction.open_time : '',
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
      is_accessible: false,
      start_date: '',
      end_date: '',
      working_hours: '',
      open_days: [],
      open_time: '',
    });
  };

  const handleInputChange = (event) => {
    const { name, value, type, checked } = event.target;
    const finalValue = type === 'checkbox' ? checked : value;
    setFormData((prevData) => ({ ...prevData, [name]: finalValue }));
  };

  const handleCheckboxChange = (event) => {
    const { value, checked } = event.target;
    setFormData((prevData) => {
      const currentDays = prevData.open_days;
      if (checked) {
        return { ...prevData, open_days: [...currentDays, value] };
      } else {
        return { ...prevData, open_days: currentDays.filter((day) => day !== value) };
      }
    });
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
      id: parseInt(id, 10),
      title: formData.title,
      type: culturalType,
      description: formData.description,
      price: formData.price,
      is_accessible: formData.is_accessible,
      location: formData.location,
      organizer_id: userID,
    };

    if (culturalType === 'event') {
      finalPayload.event = {
        start_date: formData.start_date,
        end_date: formData.end_date,
        working_hours: formData.working_hours,
      };
    } else if (culturalType === 'tourist_attraction') {
      finalPayload.tourist_attraction = {
        open_days: formData.open_days.join(', '),
        open_time: formData.open_time,
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

      if (response.ok) {
        const result = await response.json();
        navigate(`/card/${culturalType}/${id}`);
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

  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);

  const handleNotificationIconClick = async () => {
    const fetchNewNotifications = async () => {
      try {
        const response = await fetch(`http://localhost:8080/notifications/${userID}`, {
          credentials: 'include',
        });
        if (response.ok) {
          const data = await response.json();
          setNotification(data.culturals);
          console.log('Notificações recebidas:', data.culturals);
        }
      } catch (error) {
        console.error('Erro ao buscar por novas notificações:', error);
      }
    };
    fetchNewNotifications();
    setNotificationModalOpen(true, notification);
  };

  const handleNotificationCloseClick = async () => {
    const setNotificationsAsSeen = async () => {
      try {
        const response = await fetch(`http://localhost:8080/notifications/${userID}/seen`, {
          method: 'PATCH',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ notificationIDs: notification.map((notif) => notif.ID) }),
          credentials: 'include',
        });
        if (response.ok) {
          console.log('Notificações marcadas como vistas com sucesso.');
        }
      } catch (error) {
        console.error('Erro ao atualizar favorito:', error);
      }
    };
    setNotificationsAsSeen();
    setNotificationModalOpen(false);
  };
  
  const closeConfirmModal = () => {
    setConfirmModalOpen(false);
  };

  const handleConfirmLogout = () => {
    navigate('/');
  };


 
  const daysOfWeek = ['Domingo', 'Segunda', 'Terça', 'Quarta', 'Quinta', 'Sexta', 'Sábado'];

  return (
    <>
      <NotificationModal
        isOpen={isNotificationModalOpen}
        onClose={() => handleNotificationCloseClick()}
        notifications={notification}
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
        <header className="top-bar">
          <img src={logo} alt="Logo Cultural" className="logo-tiny" />
          <div className="right-section">
            <div onClick={() => setConfirmModalOpen(true)} className="icon-button-container">
              <img src={logoutIcon} alt="Log-out" className="icon" />
            </div>
            <div onClick={handleNotificationIconClick} className="icon-button-container">
              <img
                src={notificationsIcon}
                id="notifications-icon"
                alt="Notificações"
                className="icon"
              />
            </div>
          </div>
        </header>

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
                  id="tourist_attraction"
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
                    <label htmlFor="start_date" className="required">
                      Data e Hora de Início
                    </label>
                    <input
                      type="datetime-local"
                      id="start_date"
                      name="start_date"
                      value={formData.start_date}
                      onChange={handleInputChange}
                      required
                    />

                    <label htmlFor="end_date" className="required">
                      Data e Hora de Fim
                    </label>
                    <input
                      type="datetime-local"
                      id="end_date"
                      name="end_date"
                      value={formData.end_date}
                      onChange={handleInputChange}
                      required
                    />

                    <label htmlFor="working_hours" className="required">
                      Horário de Funcionamento
                    </label>
                    <input
                      type="text"
                      id="working_hours"
                      name="working_hours"
                      placeholder="HH:MM"
                      value={formData.working_hours}
                      onChange={handleInputChange}
                      required
                    />
                  </>
                ) : (
                  <>
                    <label className="required">Dias de Funcionamento</label>
                    <div className="checkbox-group">
                      {daysOfWeek.map((day) => (
                        <div key={day} className="checkbox-option">
                          <input
                            type="checkbox"
                            id={day}
                            value={day}
                            checked={formData.open_days.includes(day)}
                            onChange={handleCheckboxChange}
                          />
                          <label htmlFor={day}>{day}</label>
                        </div>
                      ))}
                    </div>

                    <label htmlFor="open_time" className="required">
                      Horário de Funcionamento
                    </label>
                    <textarea
                      id="open_time"
                      name="open_time"
                      placeholder="ex: Dom - Sab 09:00 às 17:00"
                      value={formData.open_time}
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
                    id="is_accessible"
                    name="is_accessible"
                    checked={formData.is_accessible}
                    onChange={handleInputChange}
                  />
                  <label htmlFor="is_accessible">Possui estrutura de acessibilidade</label>
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
        <footer className="footer">
          <Link to={`/home`} state={{ userID, userType }}>
            <img src={homeIcon} alt="Logo Cultural" />
          </Link>
          <Link to={`/search`} state={{ userID, userType }}>
            <img src={searchIcon} alt="Buscar" />
          </Link>
          {userType === 'organizer' && (
            <Link to={`/create-cultural`} state={{ userID, userType }}>
              <img src={addIcon} alt="Adicionar" className="mostImportantButton" />
            </Link>
          )}
          <Link to={`/user/favorites`} state={{ userID, userType }}>
            <img src={favoriteIcon} alt="Favoritos" />
          </Link>
          <Link to={`/user/profile`} state={{ userID, userType }}>
            <img src={userIcon} alt="Usuário" />
          </Link>
        </footer>
      </section>
    </>
  );
};

export default EditCulturalPage;
