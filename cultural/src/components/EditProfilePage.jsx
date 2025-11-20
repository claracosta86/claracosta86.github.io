// src/components/EditProfilePage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import './styles/profile.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import userIcon from '../assets/user-icon.png';
import homeIcon from '../assets/home-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import searchIcon from '../assets/search-icon.png';

const EditProfilePage = () => {
  const navigate = useNavigate();

  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [companyName, setCompanyName] = useState('');

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

  const [error, setError] = useState('');

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

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await fetch(`http://localhost:8080/users/${userID}/profile`, {
          credentials: 'include',
        });
        if (response.ok) {
          const data = await response.json();
          setEmail(data.email);
          setName(data.name);
          if (userType === 'organizer') {
            setCompanyName(data.companyName);
          }
        }
      } catch (error) {
        console.error('Erro ao buscar dados do usuário:', error);
        setUserType('common');
      }
    };
    fetchUserData();
  }, [userID]);

  const handleSubmit = async (event) => {
    event.preventDefault();

    setError('');

    try {
      const response = await fetch(`http://localhost:8080/users/${userID}/profile/edit`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, companyName }),
        credentials: 'include',
      });

      if (response.status === 400 || response.status === 404) {
        setError('Desculpe, tivemos um problema ao buscar seus dados. Por favor tente mais tarde.');
        return;
      }

      if (!response.ok) {
        throw new Error('Erro ao editar dados do usuário.');
      }

      if (response.ok) {
        console.log('Change OK:', userID);
        navigate(`/user/profile`);
      }
    } catch (err) {
      console.error('Erro:', err);
      setError('Ocorreu um erro. Tente novamente mais tarde.');
    }
  };

  const handleDelete = async () => {
    try {
      const response = await fetch(`http://localhost:8080/users/${userID}/profile/delete`, {
        method: 'DELETE',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      });

      if (response.status === 404) {
        setError('Desculpe, tivemos um problema ao buscar seus dados. Por favor tente mais tarde.');
        return;
      }

      if (!response.ok) {
        throw new Error('Erro ao editar dados do usuário.');
      }

      if (response.ok) {
        console.log('User deleted:', userID);
        navigate(`/`);
      }
    } catch (err) {
      console.error('Erro:', err);
      setError('Ocorreu um erro. Tente novamente mais tarde.');
    }
  };

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

  const closeConfirmModal = () => {
    setConfirmModalOpen(false);
  };

  const handleConfirmLogout = () => {
    navigate('/');
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
            <h2>Editar Perfil</h2>
          </div>
          <form onSubmit={handleSubmit}>
            <label htmlFor="new-name">Nome do Usuário</label>
            <input
              id="new-name"
              type="text"
              placeholder={name}
              onChange={(e) => setName(e.target.value)}
            />

            <label htmlFor="new-email">E-mail</label>
            <input
              id="new-email"
              type="email"
              placeholder={email}
              onChange={(e) => setEmail(e.target.value)}
            />

            {userType === 'organizer' && (
              <>
                <label htmlFor="new-company-name">Nome da Empresa</label>
                <input
                  id="new-company-name"
                  type="text"
                  placeholder={companyName}
                  onChange={(e) => setCompanyName(e.target.value)}
                />
              </>
            )}

            {error && <span className="error">{error}</span>}

            <div className="profile-actions">
              <div className="profile-actions-row">
                <button onClick={handleGoBackClick} className="profile-btn">
                  Cancelar
                </button>
                <button type="submit" className="profile-btn">
                  Confirmar
                </button>
              </div>
              <button onClick={handleDelete} className="delete-btn">
                Deletar Usuário
              </button>
            </div>
          </form>
        </div>
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

export default EditProfilePage;
