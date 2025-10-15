// src/components/CardPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import './styles/card.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import homeIcon from '../assets/home-icon.png';
import addIcon from '../assets/add-icon.png';
import searchIcon from '../assets/search-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import unfavoriteIcon from '../assets/unfavorite-icon.png';
import userIcon from '../assets/user-icon.png';
import locationIcon from '../assets/location-icon.png';
import clockIcon from '../assets/clock-icon.png';
import priceIcon from '../assets/price-icon.png';
import accessibleIcon from '../assets/accessibility-icon.png';
import mailIcon from '../assets/mail-icon.png';
import logoutIcon from '../assets/logout-icon.png';

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
            <button onClick={handleLinkClick(notif.id, notif.culturalType)}>{notif.title}</button>
          </p>
        );
      case 'canceled':
        return (
          <p>
            O cultural{' '}
            <button onClick={handleLinkClick(notif.id, notif.culturalType)}>{notif.title}</button> foi
            cancelado.
          </p>
        );
      case 'closed':
        return (
          <p>
            O cultural{' '}
            <button onClick={handleLinkClick(notif.id, notif.culturalType)}>{notif.title}</button> foi
            encerrado.
          </p>
        );
      case 'commented':
        return (
          <p>
            Veja os novos comentários de{' '}
            <button onClick={handleLinkClick(notif.id, notif.culturalType)}>{notif.title}</button>.
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

const CardPage = () => {
  const navigate = useNavigate();

  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);

  const { id, culturalType } = useParams();

  const [culturalData, setCulturalData] = useState(null);

  const { user } = useUser();
  const userID = user.userID;
  const userType = user.type;

  useEffect(() => {
    const fetchCulturalDataAndFavorites = async () => {
          console.log(`%cuseEffect ACIONADO em ${new Date().toLocaleTimeString()}`, 'color: orange');

      try {
        console.log(`Fetching cultural details for culturalType: ${culturalType}, id: ${id}`);
        const culturalResponse = await fetch(`http://localhost:8080/culturais/${culturalType}/${id}`);
        const cultural = await culturalResponse.json();

        console.log(`Fetching favorites for userID: ${userID}`);
        const favoritesResponse = await fetch(`http://localhost:8080/users/${userID}/favorites/`);
        const favoritesData = await favoritesResponse.json();

        const isFav = favoritesData.some((fav) => fav.id == cultural.id);

        if (isFav) {
          try {
            const updateLastSeen = await fetch(
              `http://localhost:8080/users/${userID}/favorites/last-seen`,
              {
                method: 'PATCH',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ culturalID: cultural.id, culturalType: culturalType }),
                credentials: 'include',
              }
            );
            if (!updateLastSeen.ok) {
              console.error('Falha ao atualizar último visto.');
            }
          } catch (error) {
            console.error('Erro ao atualizar último visto:', error);
          }
        }

        console.log('Cultural details fetched:', cultural);
        console.log('User favorites fetched:', favoritesData);
        console.log(`Is favorite: ${isFav}`);
        setCulturalData({
          ...cultural,
          isFavorite: isFav,
        });
      } catch (error) {
        console.error('Erro ao buscar detalhes ou favoritos:', error);
      }
    };
    fetchCulturalDataAndFavorites();
  }, [id, culturalType, userID]);

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

  const handleFavoriteIconClick = async () => {
    if (!culturalData) return;
    try {
      const response = await fetch(`http://localhost:8080/users/${userID}/favorites`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          isFavorite: !culturalData.isFavorite,
          culturalType: culturalType,
          culturalID: culturalData.id,
        }),
        credentials: 'include',
      });
      if (response.ok) {
        setCulturalData((prevData) => ({
          ...prevData,
          isFavorite: !prevData.isFavorite,
        }));
      }
    } catch (error) {
      console.error('Erro ao atualizar favorito:', error);
    }
  };

  const handleNotificationCloseClick = async () => {
    if (notification.length === 0) {
      setNotificationModalOpen(false);
      return;
    }
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

  if (!culturalData) {
    return <div>Carregando...</div>; 
  }

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
          <section>
            <div className="details-container">
              <div className="header-details">
                <h2 className="cultural-title">{culturalData.title}</h2>
              </div>
              <img
                src={`/thumb-size/${culturalData.image}`}
                alt={culturalData.title}
                className="event-image"
              />

              <div className="info-box">
                <p>
                  <img src={locationIcon} alt="Localização" className="info-icon" />
                  <strong>Endereço:</strong> {culturalData.location}
                </p>
                <p>
                  <img src={clockIcon} alt="Horário" className="info-icon" />
                  {culturalType === 'event' ? (
                    <strong>Data e Horário:</strong>
                  ) : (
                    <strong>Horário de Funcionamento:</strong>
                  )}
                  {culturalType === 'event' &&
                    culturalData.event &&
                    culturalData.event.end_date === '' &&
                    ` ${culturalData.event.start_date}, de ${culturalData.event.duration_time}`}
                  {culturalType === 'event' &&
                    culturalData.event &&
                    ` ${culturalData.event.start_date} - ${culturalData.event.end_date}, de ${culturalData.event.duration_time}`}
                  {culturalType !== 'event' &&
                    culturalData.tourist_attraction &&
                    ` ${culturalData.tourist_attraction.open_days}, ${culturalData.tourist_attraction.open_time}`}
                </p>
                <p>
                  <img src={priceIcon} alt="Preço" className="info-icon" />
                  <strong>Preço: </strong>
                  {culturalData.price === 'R$0,00' || culturalData.price === 'Gratuito'
                    ? 'Gratuito'
                    : `${culturalData.price}`}
                </p>
                <p>
                  <img src={accessibleIcon} alt="Acessível" className="info-icon" />
                  <strong>Acessível:</strong> {culturalData.accessible ? 'Sim' : 'Não'}
                </p>
                <p>
                  <img src={mailIcon} alt="Contato" className="info-icon" />
                  <strong>Contato:</strong> {culturalData.organizer.email}
                </p>
                <p className="description">
                  <strong>Descrição:</strong> {culturalData.description}
                </p>
              </div>
            </div>

            <div id="organizer-container">
              <button className="see-organizer-btn">
                <Link to={`/organizer/${culturalData.organizer.id}`}>Conhecer Organizador</Link>
              </button>
            </div>

            <div className="comments-section">
              <h3>Comentários</h3>
              <div className="comment-box">
                {/* A lista de comentários será renderizada aqui. 
                    A caixa ficará vazia se não houver comentários.
                  */}
              </div>
              <button className="add-comment-btn">Adicionar Comentário</button>
            </div>

            <div className="down-container">
              <div className="down-container-row">
                <button onClick={handleGoBackClick} className="down-btn">
                  Voltar
                </button>
                <button
                  onClick={handleFavoriteIconClick}
                  src={culturalData.isFavorite ? favoriteIcon : unfavoriteIcon}
                  alt="Favoritar"
                  className="down-btn"
                >
                  {culturalData.isFavorite ? '	Desfavoritar ♡' : 'Favoritar ❤'}
                </button>
              </div>
            </div>
          </section>
        </main>
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

export default CardPage;
