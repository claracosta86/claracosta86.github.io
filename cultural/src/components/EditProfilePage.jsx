// src/components/EditProfilePage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import Header from './Layout/Header';
import Footer from './Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/profile.css';

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

  const {
    notifications,
    isNotificationModalOpen,
    setNotificationModalOpen,
    fetchNotifications,
    markNotificationsAsSeen,
  } = useNotifications(userID);

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
      }
    };
    fetchUserData();
  }, [userID, userType]);

  const handleSubmit = async (event) => {
    setError('');
    event.preventDefault();

    try {
      const response = await fetch(`http://localhost:8080/users/${userID}/profile/edit`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, companyName }),
        credentials: 'include',
      });

      if (response.status === 400 || response.status === 404) {
        setError(
          'Desculpe, tivemos um problema ao buscar seus dados. Por favor tente mais tarde.'
        );
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
        setError(
          'Desculpe, tivemos um problema ao buscar seus dados. Por favor tente mais tarde.'
        );
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
            <h2>Editar Perfil</h2>
          </div>
          <form onSubmit={handleSubmit}>
            <label htmlFor="new-name">Nome do Usuário</label>
            <input
              id="new-name"
              type="text"
              value={name}
              placeholder="Novo nome (se tiver alteração)"
              onChange={(e) => setName(e.target.value)}
            />
            <label htmlFor="new-email">E-mail</label>
            <input
              id="new-email"
              type="email"
              value={email}
              placeholder="Novo e-mail (se tiver alteração)"
              onChange={(e) => setEmail(e.target.value)}
            />

            {userType === 'organizer' && (
              <>
                <label htmlFor="new-company-name">Nome da Empresa</label>
                <input
                  id="new-company-name"
                  type="text"
                  value={companyName}
                  placeholder="Novo nome da empresa (se tiver alteração)"
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
        <Footer userType={userType} />
      </section>
    </>
  );
};

export default EditProfilePage;
