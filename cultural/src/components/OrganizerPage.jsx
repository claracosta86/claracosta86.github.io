// src/components/OrganizerPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import './styles/organizer.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import userIcon from '../assets/user-icon.png';
import homeIcon from '../assets/home-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import searchIcon from '../assets/search-icon.png';

const OrganizerPage = () => {
  const navigate = useNavigate();

  const { user } = useUser();
  const userID = user.userID;
  const userType = user.type;

  const { id: organizerID } = useParams();

  const [organizerData, setOrganizerData] = useState({
    name: '',
    email: '',
    organizer_since: '',
    id: organizerID,
  });
  const [culturais, setCulturais] = useState([]);

  useEffect(() => {
    const fetchOrganizerData = async () => {
            if (!userID) return;

            try {
            const response = await fetch(`http://localhost:8080/users/${organizerID}/culturais/organizer`);
            if (!response.ok) {
                console.error('Falha ao buscar dados do organizador.');
                return; 
            }

            const data = await response.json();
            setOrganizerData({
                name: data.name,
                email: data.email,
                organizer_since: data.organizer_since,
                id: data.id,
            });

            if (data.cultural_items && data.cultural_items.length > 0) {
                const culturaisWithDetails = await Promise.all(
                data.cultural_items.map(async (cult) => {
                    if (cult.type && cult.id) {
                    const detailResponse = await fetch(
                        `http://localhost:8080/culturais/${cult.type}/${cult.id}`
                    );
                    if (detailResponse.ok) {
                        const detailData = await detailResponse.json();
                        return {
                        ...cult, 
                        Title: detailData.title,
                        Image: detailData.image,
                        Location: detailData.location,
                        Price: detailData.price,
                        Event: cult.type === 'event' ? detailData.event : null,
                        TouristAttraction: cult.type === 'tourist_attraction' ? detailData.tourist_attraction : null,
                        };
                    }
                    }
                    return cult;
                })
                );
                setCulturais(culturaisWithDetails);
            } else {
                setCulturais([]);
            }

            } catch (error) {
            console.error('Erro na requisição:', error);
            }
        };

        fetchOrganizerData();
    }, [userID]);

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

  const handleGoBackClick = () => {
    navigate(-1);
  };

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
      <section className="screen" id="tela-profile">
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
        <div className="profile-box">
          <div className="header-title">
            <h2>Conheça o Organizador</h2>
          </div>

          <div key={organizerData.id} className="organizer-info">
            <p><b>Nome:</b> {organizerData.name}</p>
            <p><b>Contato:</b> {organizerData.email}</p>
            <p><b>Tempo na plataforma:</b> {organizerData.organizer_since}</p>
          </div>
          <p className="separator-unique"></p>
          <div className="favorites-list">
            <h3>Culturais do Organizador</h3>
            {culturais.length > 0 ? (
              culturais.map(
                (cult) =>
                  cult.Title &&
                  cult.Image &&
                  cult.Location && (
                    <div key={cult.id} className="favorite-card">
                      <Link
                        to={`/card/${cult.type}/${cult.id}`}
                        state={{ userID, userType, isEvent: cult.type === 'event' }}
                        className="card-link"
                      >
                        <img
                          src={`/thumb-size/${cult.Image}`}
                          alt={cult.Title}
                          className="favorite-img"
                        />
                        <div className="favorite-details">
                          <h3>{cult.Title}</h3>
                          <p>{cult.type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
                          <div className="details-box">
                            <span id='Working Hours'>
                                {cult.type === 'event' && cult.Event && (
                                  cult.Event.end_date === "" 
                                    ? ` ${cult.Event.start_date}, de ${cult.Event.working_hours}`
                                    : ` ${cult.Event.start_date} - ${cult.Event.end_date}, de ${cult.Event.working_hours}`
                                )}
                                {cult.type !== 'event' &&
                                  cult.TouristAttraction &&
                                  ` ${cult.TouristAttraction.open_days}, ${cult.TouristAttraction.open_time}`}
                            </span>
                            <span className="price">
                              {cult.Price === 'R$0,00' || cult.Price === 'Gratuito'
                                ? 'Gratuito'
                                : `${cult.Price}`}
                            </span>
                          </div>
                        </div>
                      </Link>
                    </div>
                  )
              )
            ) : (
              <p>
                <em>Você ainda não adicionou nenhum cultural.</em>
              </p>
            )}
          </div>
          <div className="profile-actions">
            <div className="profile-actions-row">
              <button onClick={handleGoBackClick} className="profile-btn">
                Voltar
              </button>
            </div>
          </div>
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

export default OrganizerPage;
