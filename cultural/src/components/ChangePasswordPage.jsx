// src/components/ChangePasswordPage.jsx
import { useUser } from './UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import './styles/profile.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import logoutIcon from '../assets/logout-icon.png';
import userIcon from '../assets/user-icon.png';
import homeIcon from '../assets/home-icon.png';
import addIcon from '../assets/add-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import searchIcon from '../assets/search-icon.png';

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

const ChangePasswordPage = () => {
  const navigate = useNavigate();

  const [newPassword, setNewPassword] = useState('');
  const [newPasswordConfirmation, setNewPasswordConfirmation] = useState('');
  const [currentPassword, setCurrentPassword] = useState('');

  const [error, setError] = useState('');

  const user = useUser();
  const userID = user.userID;
  const userType = user.type;
  
  useEffect(() => {
    if (userID) {
      console.log("UserID recebido:", userID);
    }
    if (userType) {
      console.log("UserType recebido:", userType);
    }
  }, [userID, userType]);
  
  const handleSubmit = async (event) => {
    event.preventDefault();

    setError('');

    if (newPassword.length < 8) {
      setError("A senha deve ter pelo menos 8 caracteres.");
      return;
    }

    if (newPassword !== newPasswordConfirmation) {
      setError('A nova senha não coincide com a confirmação.');
      return;
    }

    try {
        const response = await fetch(`http://localhost:8080/users/${userID}/profile/change-password`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ currentPassword, newPassword }),
        credentials: 'include'
      });

      if (response.status === 401) {
        setError("A senha atual está incorreta.");
        return;
      }

      if (response.status === 400 || response.status === 404) {
        setError("Desculpe, tivemos um problema ao buscar seus dados. Por favor tente mais tarde.");
        return;
      }

      if (!response.ok) {
        throw new Error("Erro ao editar dados do usuário.");
      }

      if (response.ok) {
        console.log("Change OK:", userID);
        navigate(`/user/profile`);
      }

    } catch (err) {
      console.error("Erro:", err);
      setError("Ocorreu um erro. Tente novamente mais tarde.");
    }
  };

  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);
  
    const handleNotificationIconClick = async () => {
      const fetchNewNotifications = async () => {
          try {
            const response = await fetch(`http://localhost:8080/notifications/${userID}`, {
                 credentials: 'include'
            });
            if (response.ok) {
              const data = await response.json();
              setNotification(data.culturals);
              console.log("Notificações recebidas:", data.culturals);
            }
          } catch (error) {
            console.error("Erro ao buscar por novas notificações:", error);
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

  const handleGoBackClick = () => {
    navigate(-1);
  };
    

  return (
    <>
    <NotificationModal isOpen={isNotificationModalOpen} onClose={() => handleNotificationCloseClick()} notifications={notification} navigate={navigate} userID={userID} userType={userType}/> 
    <section className="screen" id="tela-profile">
      <header className="top-bar">
              <img src={logo} alt="Logo Cultural" className="logo-tiny" />
              <div className="right-section">
                <Link to="/">
                  <img src={logoutIcon} alt="Log-out" className="icon" />
                </Link>
                <div onClick={handleNotificationIconClick} className="icon-button-container">
                  <img src={notificationsIcon} id="notifications-icon" alt="Notificações" className="icon" />
                </div>
              </div>
            </header>
      <div className="profile-box">
        <div className="header-title">
          <h2>Alterar Senha</h2>
        </div>
        <form onSubmit={handleSubmit}>
          <label htmlFor="new-password">Nova Senha</label>
          <input id="new-password" type="password" placeholder="Nova senha" value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)} />

          <label htmlFor="new-password-confirmation">Confirme a Nova Senha</label>
          <input id="new-password-confirmation" type="password" placeholder="Repita a nova senha" value={newPasswordConfirmation}
            onChange={(e) => setNewPasswordConfirmation(e.target.value)} />

          <label htmlFor="current-password">Senha Atual</label>
          <input id="current-password" type="password" placeholder="Digite sua senha atual" value={currentPassword}
            onChange={(e) => setCurrentPassword(e.target.value)} />

          {error && <span className="error">{error}</span>}

          <div className="profile-actions">
            <div className="profile-actions-row">
              <button onClick={handleGoBackClick(userID, userType)} className="profile-btn">Cancelar</button>
              <button type="submit" className="profile-btn">Alterar Senha</button>
            </div>
          </div>
        </form>
      </div>
      <footer className="footer">
        <Link to={`/home`}  state={{ userID, userType }}><img src={homeIcon} alt="Logo Cultural" /></Link>
        <Link to={`/search`} state={{ userID, userType }}><img src={searchIcon} alt="Buscar"/></Link>
        {userType === 'organizer' && (
          <Link to={`/create-cultural`} state={{ userID, userType }}><img src={addIcon} alt="Adicionar" className='mostImportantButton' /></Link>
        )}
        <Link to={`/user/favorites`} state={{ userID, userType }}><img src={favoriteIcon} alt="Favoritos" /></Link>
        <Link to={`/user/profile`} state={{ userID, userType }}><img src={userIcon} alt="Usuário" /></Link>
      </footer>
    </section>
    </>
  );
};

export default ChangePasswordPage;