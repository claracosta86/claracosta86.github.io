// src/components/HomePage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import './styles/home.css';
import logo from '../assets/logo.png';
import homeIcon from '../assets/home-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';
import searchIcon from '../assets/search-icon.png';

const HomePage = () => {
  const navigate = useNavigate();

  const [events, setEvents] = useState([]);
  const [attractions, setAttractions] = useState([]);
  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

  const { user } = useUser();

  const userID = user.userID;
  const userType = user.type;

  useEffect(() => {
    const fetchCulturals = async () => {
      try {
        const response = await fetch('http://localhost:8080/culturais/home');
        const data = await response.json();
        setEvents(data.events || []);
        setAttractions(data.tourist_attractions || []);
      } catch (error) {
        console.error('Erro ao buscar culturais:', error);
      }
    };
    fetchCulturals();
  }, []);

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

        <main className="home-container">
          <section className="background-container">
            <div className="category-box">
              <h2>Principais Eventos</h2>
              <div className="card-events">
                {events.map((event) => (
                  <Link to={`/card/event/${event.id}`} state={{ event: true }} key={event.id}>
                    <div className="card">
                      <p className="title">{event.title}</p>
                      <img
                        src={`http://localhost:8080/static/culturalthumbs/${event.image}`}
                        alt={event.title}
                      />
                    </div>
                  </Link>
                ))}
              </div>
              <h2>Principais Pontos Turísticos</h2>
              <div className="card-attractions">
                {attractions.map((attraction) => (
                  <Link
                    to={`/card/tourist_attraction/${attraction.id}`}
                    state={{ event: false }}
                    key={attraction.id}
                  >
                    <div className="card">
                      <p className="title">{attraction.title}</p>
                      <img
                        src={`http://localhost:8080/static/culturalthumbs/${attraction.image}`}
                        alt={attraction.title}
                      />
                    </div>
                  </Link>
                ))}
              </div>
            </div>
          </section>
        </main>
        <footer className="footer">
          <Link to={`/home`}>
            <img src={homeIcon} alt="Logo Cultural" />
          </Link>
          <Link to={`/search`}>
            <img src={searchIcon} alt="Buscar" />
          </Link>
          {userType === 'organizer' && (
            <Link to={`/create-cultural`}>
              <img src={addIcon} alt="Adicionar" className="mostImportantButton" />
            </Link>
          )}
          <Link to={`/user/favorites`}>
            <img src={favoriteIcon} alt="Favoritos" />
          </Link>
          <Link to={`/user/profile`}>
            <img src={userIcon} alt="Usuário" />
          </Link>
        </footer>
      </section>
    </>
  );
};

export default HomePage;
