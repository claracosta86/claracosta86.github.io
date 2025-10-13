// src/components/ManageCulturalPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './styles/profile.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import userIcon from '../assets/user-icon.png';
import homeIcon from '../assets/home-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import searchIcon from '../assets/search-icon.png';

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

const ManageCulturalPage = () => {
  const navigate = useNavigate();

  const { user } = useUser();
  const userID = user.userID;
  const userType = user.type;

  useEffect(() => {
    if (userID) {
      console.log('UserID recebido:', userID);
    }
    if (userType) {
      console.log('UserType recebido:', userType);
    }
  }, [userID, userType]);

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

  const [culturais, setCulturais] = useState([]);
  useEffect(() => {
    const fetchCulturais = async () => {
      if (!userID) return;
      try {
        const response = await fetch(`http://localhost:8080/users/${userID}/culturais`);
        if (response.ok) {
          const data = await response.json();
          const culturaisWithDetails = await Promise.all(
            data.map(async (cult) => {
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
                  };
                }
              }
              return cult;
            })
          );
          setCulturais(culturaisWithDetails || []);
        } else {
          console.error('Falha ao buscar culturais.');
          setCulturais([]);
        }
      } catch (error) {
        console.error('Erro ao buscar culturais:', error);
      }
    };
    fetchCulturais();
  }, [userID]);

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
            <h2>Meus Culturais</h2>
          </div>
          <div className="favorites-list">
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
                          <p>{cult.Type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
                          <span>{cult.Location}</span>
                          <span className="price">
                            {cult.Price === 'R$0,00' || cult.Price === 'Gratuito'
                              ? 'Gratuito'
                              : `${cult.Price}`}
                          </span>
                        </div>
                      </Link>
                      <button className="remove-btn">Remover</button>
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

export default ManageCulturalPage;
