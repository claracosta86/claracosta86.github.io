import { useUser } from '../contexts/UserContext';
import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import './styles/create.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import userIcon from '../assets/user-icon.png';
import homeIcon from '../assets/home-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import searchIcon from '../assets/search-icon.png';

const CreateCulturalPage = () => {
  const navigate = useNavigate();

  const { user } = useUser();
  const userID = user.userID;
  const userType = user.type;

  const [selectedFile, setSelectedFile] = useState(null);
  const [culturalType, setCulturalType] = useState('');
  const [error, setError] = useState('');

  // Estado único para gerenciar todos os campos do formulário
  const [formData, setFormData] = useState({
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
      title: formData.title,
      type: culturalType,
      description: formData.description,
      price: formData.price,
      is_accessible: formData.is_accessible,
      organizer: {
        id: userID,
      },
    };

    if (culturalType === 'event') {
      finalPayload.location = formData.location;
      finalPayload.event = {
        start_date: formData.start_date,
        end_date: formData.end_date,
        working_hours: formData.working_hours,
      };
    } else if (culturalType === 'attraction') {
      finalPayload.location = formData.location;
      finalPayload.tourist_attraction = {
        open_days: formData.open_days.join(', '),
        open_time: formData.open_time,
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
      <section className="screen" id="tela-create">
        <header className="top-bar">
          <img src={logo} alt="Logo Cultural" className="logo-tiny" />
          <div className="right-section">
            <Link to="/">
              <img src={logoutIcon} alt="Log-out" className="icon" />
            </Link>
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
                      placeholder="ex: 09:00 às 17:00 ou 24 horas"
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

                <p className="button-container">
                  <button type="submit" className="btn">
                    Criar
                  </button>
                </p>
              </>
            )}
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

export default CreateCulturalPage;
