// src/components/ManageCulturalPage.jsx
import { useUser } from '../contexts/UserContext';
import RemoveModal from './RemoveModal/RemoveFromApp'; 
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './styles/profile.css';
import './styles/manage.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import userIcon from '../assets/user-icon.png';
import homeIcon from '../assets/home-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import searchIcon from '../assets/search-icon.png';

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

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

  const [isRemoveModalOpen, setRemoveModalOpen] = useState(false);
  const [culturalToRemove, setCulturalToRemove] = useState(null);

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
                    Event: cult.type === 'event' ? detailData.event : null,
                    TouristAttraction: cult.type === 'tourist_attraction' ? detailData.touristAttraction : null,
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

  const handleEditClick = (id, type) => {
    navigate(`/edit-cultural/${type}/${id}`);
  };

  const openRemoveModal = (cultural) => {
    setCulturalToRemove(cultural);
    setRemoveModalOpen(true);
  };

  const closeRemoveModal = () => {
    setRemoveModalOpen(false);
    setCulturalToRemove(null);
  };

  const handleConfirmRemove = async () => {
    if (!culturalToRemove) return;

    try {
      const { id, type } = culturalToRemove;
      const response = await fetch(`http://localhost:8080/culturais/${type}/${id}`, {
        method: 'DELETE',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      });

      if (response.ok) {
        setCulturais((prevCulturais) => prevCulturais.filter((cult) => cult.id !== id));
      } else {
        console.error('Não foi possível excluir o cultural. Tente novamente.');
      }
    } catch (error) {
      console.error('Erro de rede ao remover cultural:', error);
    } finally {
      closeRemoveModal();
    }
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
        <div className="profile-box">
          <div className="header-title">
            <h2>Meus Culturais</h2>
          </div>
          <div className="favorites-list">
            {culturais.length > 0 ? (
              culturais.map(
                (cult) =>
                  cult.Title &&
                   (
                    <div key={cult.id} className="manage-card">
                      <Link
                        to={`/card/${cult.type}/${cult.id}`}
                        state={{ userID, userType, isEvent: cult.type === 'event' }}
                        className="card-link"
                      >
                        <img
                          src={`http://localhost:8080/static/culturalthumbs/${cult.Image}`}
                          alt={cult.Title}
                          className="manage-img"
                        />
                        <div className="manage-details">
                          <h3>{cult.Title}</h3> 
                          <p>{cult.type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
                          <div className="details-content">
                            <span>{cult.Location}</span> <br />
                            <span>
                                {cult.type === 'event' && cult.Event && (
                                  cult.Event.endDate === "" 
                                    ? ` ${cult.Event.startDate}, de ${cult.Event.durationHours}`
                                    : ` ${cult.Event.startDate} - ${cult.Event.endDate}, de ${cult.Event.durationHours}`
                                )}
                                {cult.type !== 'event' &&
                                  cult.TouristAttraction &&
                                  ` ${cult.TouristAttraction.workingHours}`}
                            </span>
                            <span className="price">
                              {cult.Price === 'R$0,00' || cult.Price === 'Gratuito'
                                ? 'Gratuito'
                                : `${cult.Price}`}
                            </span>

                          </div>
                        </div>
                      </Link>
                      <div className="actions-row">
                        <button onClick={() => handleEditClick(cult.id, cult.type)} className="edit-btn">Editar</button>
                      <button onClick={() => openRemoveModal(cult)} className="remove-btn">Excluir</button>
                      </div>
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
