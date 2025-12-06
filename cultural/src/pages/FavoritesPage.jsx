// src/components/FavoritesPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import RemoveModal from '../components/RemoveModal/RemoveFromFavorites';
import NotificationModal from '../components/NotificationModal/NotificationModal';
import ConfirmModal from '../components/ConfirmModal/ConfirmComment';
import Header from '../components/Layout/Header';
import Footer from '../components/Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/favorites.css';

const FavoritesPage = () => {
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

  const [favorites, setFavorites] = useState([]);

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

  const [isRemoveModalOpen, setRemoveModalOpen] = useState(false);
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
                    TouristAttraction:
                      fav.type === 'tourist_attraction' ? detailData.touristAttraction : null,
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
    setRemoveModalOpen(true);
  };

  const closeRemoveModal = () => {
    setRemoveModalOpen(false);
    setFavoriteToRemove(null);
  };

  const closeConfirmModal = () => {
    setConfirmModalOpen(false);
  };

  const handleConfirmLogout = () => {
    navigate('/');
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

  return (
    <>
      <section className="screen" id="tela-home">
        <NotificationModal
          isOpen={isNotificationModalOpen}
          onClose={markNotificationsAsSeen}
          notifications={notifications}
          navigate={navigate}
          userID={userID}
          userType={userType}
        />
        <RemoveModal
          isOpen={isRemoveModalOpen}
          onClose={closeRemoveModal}
          onConfirm={handleConfirmRemove}
        />
        <ConfirmModal
          isOpen={isConfirmModalOpen}
          onClose={closeConfirmModal}
          onConfirm={handleConfirmLogout}
        />
        
        <Header
          onLogoutClick={() => setConfirmModalOpen(true)}
          onNotificationClick={fetchNotifications}
        />

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
                          <div className="details-box">
                            <span>
                              {fav.type === 'event' &&
                                fav.Event &&
                                (fav.Event.endDate === ''
                                  ? ` ${fav.Event.startDate}, de ${fav.Event.durationHours}`
                                  : ` ${fav.Event.startDate} - ${fav.Event.endDate}, de ${fav.Event.durationHours}`)}
                              {fav.type === 'tourist_attraction' &&
                                fav.TouristAttraction &&
                                ` ${fav.TouristAttraction.workingHours}`}
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
        <Footer userType={userType} />
      </section>
    </>
  );
};

export default FavoritesPage;
