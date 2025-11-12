// src/components/FavoritesPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import './styles/favorites.css';
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

const RemoveFavoriteModal = ({ isOpen, onClose, onConfirm }) => {
  if (!isOpen) return null;

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={(e) => e.stopPropagation()}>
        <h2 className="modal-title">Confirmar Remoção</h2>
        <div className="modal-content">
          <p>
            Você realmente deseja remover o evento/ponto turístico solicitado da sua lista de
            favoritos?
          </p>
        </div>
        <div className="modal-actions">
          <button onClick={onClose} className="modal-close-btn">
            Cancelar
          </button>
          <button onClick={onConfirm} className="modal-confirm-btn">
            Remover
          </button>
        </div>
      </div>
    </div>
  );
};

const FavoritesPage = () => {
  const navigate = useNavigate();

  const { user } = useUser();
  const userID = user.userID;
  const userType = user.type;

  const [favorites, setFavorites] = useState([]);
  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);

  const [isRemoveFavoriteModalOpen, setRemoveFavoriteModalOpen] = useState(false);
  const [favoriteToRemove, setFavoriteToRemove] = useState(null);

  useEffect(() => {
    const fetchFavorites = async () => {
      if (!userID) return;
      try {
        const response = await fetch(`http://localhost:8080/users/${userID}/favorites`);
        if (response.ok) {
          const data = await response.json();
          const favoritesWithDetails = await Promise.all(
            data.map(async (fav) => {
              if (fav.type && fav.id) {
                const detailResponse = await fetch(
                  `http://localhost:8080/culturais/${fav.type}/${fav.id}`
                );
                if (detailResponse.ok) {
                  const detailData = await detailResponse.json();
                  return {
                    ...fav,
                    Title: detailData.title,
                    Image: detailData.image,
                    Price: fav.type === 'event' ? detailData.price : 'Gratuito',
                    Event: fav.type === 'event' ? detailData.event : null,
                    TouristAttraction: fav.type === 'tourist_attraction' ? detailData.tourist_attraction : null,
                  };
                }
              }
              return fav;
            })
          );
          setFavorites(favoritesWithDetails || []);
        } else {
          console.error('Falha ao buscar favoritos.');
          setFavorites([]);
        }
      } catch (error) {
        console.error('Erro ao buscar favoritos:', error);
      }
    };
    fetchFavorites();
  }, [userID]);

  const openRemoveModal = (favorite) => {
    setFavoriteToRemove(favorite);
    setRemoveFavoriteModalOpen(true);
  };

  const closeRemoveModal = () => {
    setRemoveFavoriteModalOpen(false);
    setFavoriteToRemove(null);
  };

  const handleConfirmRemove = async () => {
    if (!favoriteToRemove) return;

    try {
      const { id, type } = favoriteToRemove;
      const response = await fetch(`http://localhost:8080/users/${userID}/favorites`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ isFavorite: false, culturalType: type, culturalID: id }),
        credentials: 'include',
      });

      if (response.ok) {
        setFavorites((prevFavorites) => prevFavorites.filter((fav) => fav.id !== id));
      } else {
        alert('Não foi possível remover o favorito. Tente novamente.');
      }
    } catch (error) {
      console.error('Erro de rede ao remover favorito:', error);
    } finally {
      closeRemoveModal();
    }
  };

  const handleNotificationIconClick = async () => {
    const fetchNewNotifications = async () => {
      try {
        const response = await fetch(`http://localhost:8080/notifications/${userID}`, {
          credentials: 'include',
        });
        if (response.ok) {
          const data = await response.json();
          setNotification(data.culturals);
        }
      } catch (error) {
        console.error('Erro ao buscar por novas notificações:', error);
      }
    };
    fetchNewNotifications();
    setNotificationModalOpen(true);
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
      <RemoveFavoriteModal
        isOpen={isRemoveFavoriteModalOpen}
        onClose={closeRemoveModal}
        onConfirm={handleConfirmRemove}
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

        <div className="favorites-box">
          <div className="header-title">
            <h2>Meus Favoritos</h2>
          </div>

          <div className="favorites-list">
            {favorites.length > 0 ? (
              favorites.map(
                (fav) =>
                  fav.Title &&
                  fav.Image &&
                  (fav.Event || fav.TouristAttraction) && (
                    <div key={fav.id} className="favorite-card">
                      <Link
                        to={`/card/${fav.type}/${fav.id}`}
                        state={{ userID, userType, isEvent: fav.type === 'event' }}
                        className="card-link"
                      >
                        <img
                          src={`http://localhost:8080/static/culturalthumbs/${fav.Image}`}
                          alt={fav.Title}
                          className="favorite-img"
                        />
                        <div className="favorite-details">
                          <h3>{fav.Title}</h3>
                          <p>{fav.type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
                          <div className ="details-box">
                            <span>
                              {fav.type === 'event' && fav.Event && (
                                fav.Event.end_date === "" 
                                  ? ` ${fav.Event.start_date}, de ${fav.Event.working_hours}`
                                  : ` ${fav.Event.start_date} - ${fav.Event.end_date}, de ${fav.Event.working_hours}`
                              )}
                              {fav.type === 'tourist_attraction' &&
                                fav.TouristAttraction &&
                                ` ${fav.TouristAttraction.open_days}, ${fav.TouristAttraction.open_time}`}
                            </span>

                            <span className="price">
                              {fav.Price === 'R$0,00' || fav.Price === 'Gratuito'
                                ? 'Gratuito'
                                : `${fav.Price}`}
                            </span>
                          </div>
                        </div>
                      </Link>
                      <button onClick={() => openRemoveModal(fav)} className="remove-btn">
                        Remover
                      </button>
                    </div>
                  )
              )
            ) : (
              <p>
                <em>Você ainda não adicionou nenhum favorito.</em>
              </p>
            )}
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

export default FavoritesPage;
