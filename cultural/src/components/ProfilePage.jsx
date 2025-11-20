// src/components/ProfilePage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import Header from './Layout/Header';
import Footer from './Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/profile.css';

const ProfilePage = () => {
  const navigate = useNavigate();

  const [userName, setUserName] = useState('');
  const [userEmail, setUserEmail] = useState('');
  const [companyName, setCompanyName] = useState('');

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

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

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
      }
    };
    fetchUserData();
  }, [userID, userType]);

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
        onClose={markNotificationsAsSeen}
        notifications={notifications}
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
        <Header
          onLogoutClick={() => setConfirmModalOpen(true)}
          onNotificationClick={fetchNotifications}
        />

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
              <Link to="/user/profile/change-password" className="profile-btn">
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
        <Footer userType={userType} />
      </section>
    </>
  );
};

export default ProfilePage;
