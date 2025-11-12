// src/components/HomePage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './styles/home.css';
import logo from '../assets/logo.png';
import homeIcon from '../assets/home-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';
import searchIcon from '../assets/search-icon.png';
import bienalEvent from '../assets/thumb-size/bienal-event.png';
import mcrEvent from '../assets/thumb-size/mcr-event.png';
import dccWeekEvent from '../assets/thumb-size/dccweek-event.png';
import iwnbEvent from '../assets/thumb-size/iwnb-event.png';
import cruEvent from '../assets/thumb-size/cru-event.png';
import liberdadeAttraction from '../assets/thumb-size/liberdade-attraction.png';
import igrejinhaAttraction from '../assets/thumb-size/igrejinha-attraction.png';
import pseteAttraction from '../assets/thumb-size/psete-attraction.png';
import mercadoAttraction from '../assets/thumb-size/mercado-attraction.png';
import mangabeirasAttraction from '../assets/thumb-size/mangabeiras-attraction.png';

const NotificationModal = ({ isOpen, onClose, notifications, navigate, userID, userType }) => {
  if (!isOpen) return null;

  const handleLinkClick = (culturalID, culturalType) => async () => {
    let isEvent = culturalType === 'event' ? true : false;
    navigate(`/card/${culturalID}`, { state: { userID, userType, event: isEvent } });
  };

  const renderNotificationContent = (notif) => {
    switch (notif.notificationType) {
      case 'updated':
        return (
          <p>
            Veja as atualizações de{' '}
            <button onClick={handleLinkClick(notif.id, notif.type)}>{notif.title}</button>
          </p>
        );
      case 'canceled':
        return (
          <p>
            O cultural{' '}
            <button onClick={handleLinkClick(notif.id, notif.type)}>{notif.title}</button> foi
            cancelado.
          </p>
        );
      case 'closed':
        return (
          <p>
            O cultural{' '}
            <button onClick={handleLinkClick(notif.id, notif.type)}>{notif.title}</button> foi
            encerrado.
          </p>
        );
      case 'commented':
        return (
          <p>
            Veja os novos comentários de{' '}
            <button onClick={handleLinkClick(notif.id, notif.type)}>{notif.title}</button>.
          </p>
        );
      default:
        return null;
    }
  };

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={(e) => e.stopPropagation()}>
        <h2 className="modal-title">Notificações</h2>

        <div className="modal-content">
          {notifications.length === 0 ? (
            <p>Você não tem novas notificações.</p>
          ) : (
            notifications.map((notif) => (
              <div key={notif.id} className="notification-item">
                {renderNotificationContent(notif)}
              </div>
            ))
          )}
        </div>

        <div className="modal-actions">
          <button onClick={onClose} className="modal-close-btn">
            Fechar
          </button>
        </div>
      </div>
    </div>
  );
};

const HomePage = () => {
  const navigate = useNavigate();

  const iwnbEventID = '1';
  const bienalEventID = '2';
  const mcrEventID = '3';
  const cruEventID = '4';
  const dccWeekEventID = '5';
  const liberdadeAttractionID = '11';
  const mercadoAttractionID = '12';
  const igrejinhaAttractionID = '13';
  const pseteAttractionID = '15';
  const mangabeirasAttractionID = '14';

  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);

  const { user } = useUser();

  const userID = user.userID;
  const userType = user.type;

  useEffect(() => {
    if (userID) {
      console.log('UserID recebido da página de login:', userID);
    }
    if (userType) {
      console.log('UserType recebido da página de login:', userType);
    }
  }, [user]);

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
      <section className="screen" id="tela-home">
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

        <main className="home-container">
          <section className="background-container">
            <div className="category-box">
              <h2>Principais Eventos</h2>
              <div className="card-events">
                <Link to={`/card/event/${bienalEventID}`} state={{ event: true }}>
                  <div className="card">
                    <p className="title">Bienal do Livro</p>
                    <img src={bienalEvent} alt="Bienal do livro" />
                  </div>
                </Link>
                <Link to={`/card/event/${mcrEventID}`} state={{ event: true }}>
                  <div className="card">
                    <p className="title">My Chemical Romance Ao Vivo</p>
                    <img src={mcrEvent} alt="MCR Ao Vivo" />
                  </div>
                </Link>
                <Link to={`/card/event/${dccWeekEventID}`} state={{ event: true }}>
                  <div className="card">
                    <p className="title">DCC Week</p>
                    <img src={dccWeekEvent} alt="DCC Week" />
                  </div>
                </Link>
                <Link to={`/card/event/${iwnbEventID}`} state={{ event: true }}>
                  <div className="card">
                    <p className="title">I Wanna Be Tour</p>
                    <img src={iwnbEvent} alt="I Wanna Be Tour" />
                  </div>
                </Link>
                <Link to={`/card/event/${cruEventID}`} state={{ event: true }}>
                  <div className="card">
                    <p className="title">Jogo do Cruzeiro</p>
                    <img src={cruEvent} alt="Jogo do Cruzeiro" />
                  </div>
                </Link>
              </div>
              <h2>Principais Pontos Turísticos</h2>
              <div className="card-attractions">
                <Link
                  to={`/card/tourist_attraction/${liberdadeAttractionID}`}
                  state={{ event: false }}
                >
                  <div className="card">
                    <p className="title">Praça Liberdade</p>
                    <img src={liberdadeAttraction} alt="Praça da Liberdade" />
                  </div>
                </Link>
                <Link
                  to={`/card/tourist_attraction/${igrejinhaAttractionID}`}
                  state={{ event: false }}
                >
                  <div className="card">
                    <p className="title">Igreja da Pampulha</p>
                    <img src={igrejinhaAttraction} alt="Igreja da Pampulha" />
                  </div>
                </Link>
                <Link to={`/card/tourist_attraction/${pseteAttractionID}`} state={{ event: false }}>
                  <div className="card">
                    <p className="title">Pirulito da Praça Sete</p>
                    <img src={pseteAttraction} alt="Pirulito da Praça Sete" />
                  </div>
                </Link>
                <Link
                  to={`/card/tourist_attraction/${mercadoAttractionID}`}
                  state={{ event: false }}
                >
                  <div className="card">
                    <p className="title">Mercado Central</p>
                    <img src={mercadoAttraction} alt="Mercado Central" />
                  </div>
                </Link>
                <Link
                  to={`/card/tourist_attraction/${mangabeirasAttractionID}`}
                  state={{ event: false }}
                >
                  <div className="card">
                    <p className="title">Parque das Mangabeiras</p>
                    <img src={mangabeirasAttraction} alt="Parque das Mangabeiras" />
                  </div>
                </Link>
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
