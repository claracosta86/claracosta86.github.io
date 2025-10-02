// src/components/FavoritesPage.jsx
import { useState, useEffect } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import './styles/profile.css';
import './styles/favorites.css'; 
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';
import gobackIcon from '../assets/goback.png';

const NotificationModal = ({ isOpen, onClose, notifications, navigate, userID, userType }) => {
  if (!isOpen) return null;

  const handleLinkClick = (culturalID, culturalType, ) =>  async () => {
    let isEvent = culturalType === 'event' ? true : false
    navigate(`/card/${culturalID}`, { state: {userID, userType, "event": isEvent } });
  };

  const renderNotificationContent = (notif) => {
  switch (notif.notificationType) {
    case "updated":
      return <p>Veja as atualizações de <button onClick={handleLinkClick(notif.id, notif.type)}>{notif.title}</button></p>;
    case "canceled":
      return <p>O cultural <button onClick={handleLinkClick(notif.id, notif.type)}>{notif.title}</button> foi cancelado.</p>;
    case "closed":
      return <p>O cultural <button onClick={handleLinkClick(notif.id, notif.type)}>{notif.title}</button> foi encerrado.</p>;
    case "commented":
      return <p>Veja os novos comentários de <button onClick={handleLinkClick(notif.id, notif.type)}>{notif.title}</button>.</p>;
    default:
      return null;
  }
};

  
  return (
     <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={e => e.stopPropagation()}>
        <h2 className="modal-title">
          Notificações
        </h2>

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
          <button
            onClick={onClose}
            className="modal-close-btn"
          >
            Fechar
          </button>
        </div>
      </div>
    </div>
  );
};



const FavoritesPage = () => {
    const navigate = useNavigate();
    const location = useLocation();

    const userID = location.state?.userID || '';
    const userType = location.state?.userType || 'common';

    const [favorites, setFavorites] = useState([]);
    const [notification, setNotification] = useState([]);
    const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);

    useEffect(() => {
        const fetchFavorites = async () => {
            if (!userID) return;
            try {
                const response = await fetch(`http://localhost:8080/users/${userID}/favorites`);
                if (response.ok) {
                    const data = await response.json();
                    setFavorites(data || []); 
                } else {
                    console.error("Falha ao buscar favoritos.");
                    setFavorites([]);
                }
            } catch (error) {
                console.error("Erro ao buscar favoritos:", error);
            }
        };
        fetchFavorites();
    }, [userID]);

    const handleRemoveFavorite = async (culturalID, culturalType) => {
        try {
             const response = await fetch(`http://localhost:8080/users/${userID}/favorites`, {
                method: 'PATCH',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ isFavorite: false, culturalType: culturalType, culturalID: culturalID }),
                credentials: 'include'
            });

            if (response.ok) {
                setFavorites(prevFavorites => prevFavorites.filter(fav => fav.ID !== culturalID));
            } else {
                alert('Não foi possível remover o favorito. Tente novamente.');
            }
        } catch (error) {
            console.error("Erro de rede ao remover favorito:", error);
        }
    };

    const handleNotificationIconClick = async () => {
        const fetchNewNotifications = async () => {
            try {
                const response = await fetch(`http://localhost:8080/notifications/${userID}`, {
                    credentials: 'include'
                });
                if (response.ok) {
                    const data = await response.json();
                    setNotification(data.culturals);
                }
            } catch (error) {
                console.error("Erro ao buscar por novas notificações:", error);
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
            body: JSON.stringify({ notificationIDs: notification.map(notif => notif.ID) }),
            credentials: 'include'
            });
            if (response.ok) {
            console.log("Notificações marcadas como vistas com sucesso.");
            }
        } catch (error) {
            console.error("Erro ao atualizar favorito:", error);
        }
    };
        setNotificationsAsSeen();
        setNotificationModalOpen(false);
    };

    const handleUserIconClick = async () => {
        navigate('/user/profile', { state: { userID, userType } });
    };

    const headerClass = userType === 'organizer' ? 'top-bar-organizer' : 'top-bar-common';

    return (
        <> 
        <NotificationModal isOpen={isNotificationModalOpen} onClose={() => handleNotificationCloseClick()} notifications={notification} navigate={navigate} userID={userID} userType={userType}/>
            <section className="screen" id="tela-profile">
                <header className={headerClass}>
                    <div className="logo-container">
                        <Link to="/home" state={{ userID, userType }}>
                            <img src={logo} alt="Logo Cultural" className="logo-tiny" />
                        </Link>
                    </div>
                    <div className="right-section">
                        <div className="icons">
                            {userType === 'organizer' && (
                                <Link to="/create-cultural" state={{ userID, userType }} className="add-btn">Adicionar Cultural</Link>
                            )}
                            <div onClick={handleNotificationIconClick} className="icon-button-container">
                                <img src={notificationsIcon} id="notifications-icon" alt="Notificações" className="icon" />
                            </div>
                            <div onClick={handleUserIconClick} className="icon-button-container">
                                <img src={userIcon} id="user-icon" alt="Usuário" className="icon" />
                            </div>
                        </div>
                    </div>
                </header>
                <div className="profile-box">
                    <div className="header-title">
                        <button onClick={() => navigate(-1)} className="goback-btn" style={{background: 'none', border: 'none', cursor: 'pointer'}}>
                            <img src={gobackIcon} alt="Go Back Arrow" className="goback-img" />
                        </button>
                        <h2>Meus Favoritos</h2>
                    </div>
                    
                    <div className="favorites-list">
                        {favorites.length > 0 ? (
                            favorites.map(fav => (
                                <div key={fav.ID} className="favorite-card">
                                    <Link to={`/card/${fav.ID}`} state={{ userID, userType, isEvent: fav.Type === 'event' }} className="card-link">
                                        <img src={fav.Image} alt={fav.Title} className="favorite-img" />
                                        <div className="favorite-details">
                                            <h3>{fav.Title}</h3>
                                            <p>{fav.Type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
                                            <span>{fav.Location}</span>
                                            <span className="price">{fav.Price === "0" || fav.Price === "Gratuito" ? "Gratuito" : `R$${fav.Price}`}</span>
                                        </div>
                                    </Link>
                                    <button onClick={() => handleRemoveFavorite(fav.ID)} className="remove-btn">
                                        &#x2715; Remover
                                    </button>
                                </div>
                            ))
                        ) : (
                            <p><em>Você ainda não adicionou nenhum favorito.</em></p>
                        )}
                    </div>

                </div>
            </section>
        </>
    );
};

export default FavoritesPage;