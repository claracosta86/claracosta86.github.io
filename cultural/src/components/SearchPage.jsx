// src/components/HomePage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './styles/home.css';
import './styles/search.css';
import logo from '../assets/logo.png';
import homeIcon from '../assets/home-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';
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

const SearchPage = () => {
  const navigate = useNavigate();

  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);
  const [error, setError] = useState('');

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

  const handleSearch = async () => {
    if (!query.trim()) {
      setError('Por favor, digite um termo para buscar.');
      setResults([]);
      return;
    }
    setError('');
    try {
      const response = await fetch(`http://localhost:8080/culturais/search?q=${query}`);
      if (response.ok) {
        const data = await response.json();
        setResults(data || []);
        if (!data || data.length === 0) {
          setError('Nenhum resultado encontrado para sua busca.');
        }
      } else {
        setError('Erro ao buscar resultados. Tente novamente.');
      }
    } catch (err) {
      setError('Erro de conexão. Verifique sua rede e tente novamente.');
      console.error(err);
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

        <main className="search-container">
          <section className="background-container">
            <div className="category-box">
              <h2>Explore os eventos e atrações:</h2>
              <div className="search-box">
                <input
                  type="text"
                  value={query}
                  onChange={(e) => setQuery(e.target.value)}
                  placeholder="Buscar por eventos ou locais..."
                  className="search-input"
                />
                <button onClick={handleSearch} className="search-button">
                  <img src={searchIcon} alt="Buscar" />
                </button>
              </div>

              {error && <p className="error-message">{error}</p>}

              <div className="results-container">
                {results.map((item) => (
                  <Link
                    to={`/card/${item.id}`}
                    state={{ userID, userType, event: item.type === 'event' }}
                    key={item.id}
                    className="result-card"
                  >
                    <img
                      src={`/thumb-size/${item.image}`}
                      alt={item.title}
                      className="result-image"
                    />
                    <div className="result-details">
                      <h3>{item.title}</h3>
                      <p>{item.type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
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

export default SearchPage;
