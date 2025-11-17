// src/components/ProfilePage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import './styles/profile.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import userIcon from '../assets/user-icon.png';
import homeIcon from '../assets/home-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import searchIcon from '../assets/search-icon.png';

const ProfilePage = () => {
  const navigate = useNavigate();

  const [userName, setUserName] = useState('');
  const [userEmail, setUserEmail] = useState('');
  const [companyName, setCompanyName] = useState('');

  const { user } = useUser();
  const userID = user.userID;
  const userType = user.type;

  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);

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
          setUserEmail(data.email);
          setUserName(data.name);
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
            <h2>Seu Perfil</h2>
          </div>
          <label htmlFor="name">Nome</label>
          <input id="name" type="text" placeholder={userName} autoComplete="given-name" disabled />

          <label htmlFor="email">E-mail</label>
          <input id="email" type="email" placeholder={userEmail} autoComplete="email" disabled />

          {userType === 'organizer' && (
            <>
              <label htmlFor="companyName">Nome da Empresa</label>
              <input
                id="companyName"
                type="text"
                placeholder={companyName}
                autoComplete="organization"
                disabled
              />
            </>
          )}

          <div className="profile-actions">
            <div className="profile-actions-row">
              <Link to="/user/profile/edit" state={{ userType, userID }} className="profile-btn">
                Editar Perfil
              </Link>
              <Link
                to="/user/profile/change-password"
                state={{ userType, userID }}
                className="profile-btn"
              >
                Alterar Senha
              </Link>
            </div>
            {userType === 'organizer' && (
              <Link
                to="/user/profile/manage-cultural"
                state={{ userType, userID }}
                className="profile-btn"
              >
                Meus Culturais
              </Link>
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

export default ProfilePage;
