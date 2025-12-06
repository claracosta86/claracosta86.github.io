// src/components/ChangePasswordPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import NotificationModal from '../components/NotificationModal/NotificationModal';
import ConfirmModal from '../components/ConfirmModal/ConfirmComment';
import Header from '../components/Layout/Header';
import Footer from '../components/Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/profile.css';

const ChangePasswordPage = () => {
  const navigate = useNavigate();

  const [newPassword, setNewPassword] = useState('');
  const [newPasswordConfirmation, setNewPasswordConfirmation] = useState('');
  const [currentPassword, setCurrentPassword] = useState('');

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

  const handleSubmit = async (event) => {
    setError('');
    event.preventDefault();

    if (newPassword.length < 8) {
      setError('A senha deve ter pelo menos 8 caracteres.');
      return;
    }

    if (newPassword !== newPasswordConfirmation) {
      setError('A nova senha não coincide com a confirmação.');
      return;
    }

    try {
      const response = await fetch(
        `http://localhost:8080/users/${userID}/profile/change-password`,
        {
          method: 'PATCH',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ currentPassword, newPassword }),
          credentials: 'include',
        }
      );

      if (response.status === 401) {
        setError('A senha atual está incorreta.');
        return;
      }

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

  const closeConfirmModal = () => {
    setConfirmModalOpen(false);
  };

  const handleConfirmLogout = () => {
    navigate('/');
  };

  const handleGoBackClick = () => {
    setError('');
    navigate(-1);
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
        <ConfirmModal
          isOpen={isConfirmModalOpen}
          onClose={closeConfirmModal}
          onConfirm={handleConfirmLogout}
        />

        <Header
          onLogoutClick={() => setConfirmModalOpen(true)}
          onNotificationClick={fetchNotifications}
        />
        <div className="profile-box">
          <div className="header-title">
            <h2>Alterar Senha</h2>
          </div>
          <form onSubmit={handleSubmit}>
            <label htmlFor="new-password">Nova Senha</label>
            <input
              id="new-password"
              type="password"
              placeholder="Nova senha"
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
            />

            <label htmlFor="new-password-confirmation">Confirme a Nova Senha</label>
            <input
              id="new-password-confirmation"
              type="password"
              placeholder="Repita a nova senha"
              value={newPasswordConfirmation}
              onChange={(e) => setNewPasswordConfirmation(e.target.value)}
            />

            <label htmlFor="current-password">Senha Atual</label>
            <input
              id="current-password"
              type="password"
              placeholder="Digite sua senha atual"
              value={currentPassword}
              onChange={(e) => setCurrentPassword(e.target.value)}
            />

            {error && <span className="error">{error}</span>}

            <div className="profile-actions">
              <div className="profile-actions-row">
                <button onClick={handleGoBackClick} className="profile-btn">
                  Cancelar
                </button>
                <button onClick={handleSubmit} className="profile-btn">
                  Confirmar
                </button>
              </div>
            </div>
          </form>
        </div>
        <Footer userType={userType} />
      </section>
    </>
  );
};

export default ChangePasswordPage;
