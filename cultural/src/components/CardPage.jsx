// src/components/CardPage.jsx
import { useState, useEffect } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import './styles/home.css'; 
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';

const NotificationModal = ({ isOpen, onClose, notifications }) => {
  if (!isOpen) return null;

  
  const handleLinkClick = (culturalID) =>  async () => {
    navigate(`/card/${culturalID}`);
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
            notifications.map((notif, index) => (
              <div key={index} className="notification-item">
                <p> Veja as atualizações de <button onClick={() => handleLinkClick(notif.ID)}>{notif.Title}</button></p>
              </div>
            ))
          )}
        </div>

        <div className="modal-actions">
          <button
            onClick={onClose}
            className="modal-close-btn"
          >
            Entendi
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
  
  const location = useLocation();
  const userID = location.state?.userID || '';
  const userType = location.state?.userType || '';

  useEffect(() => {
    if (userID) {
      console.log("UserID recebido da página de login:", userID);
    }
    if (userType) {
      console.log("UserType recebido da página de login:", userType);
    }
  }, [userID, userType]);

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

  const handleUserIconClick = async () => {
    navigate('/user/profile',  { state: {userID: userID, userType: userType} });
  };
  
  const topBarClass = userType === 'organizer' ? 'top-bar-organizer' : 'top-bar-common';


    return (
      <>
      <NotificationModal isOpen={isNotificationModalOpen} onClose={() => setNotificationModalOpen(false)} notifications={notification} />
      <section className="screen" id="tela-home">
        <header className={topBarClass}>
          <div className="logo-container">
            <Link to="/">
              <img src={logo} alt="Logo Cultural" className="logo-tiny" />
            </Link>
          </div>
          <div className="right-section">
            <div className="icons">
              {userType === 'organizer' && (
                <a href="#" className="add-btn">Adicionar Cultural</a>
              )}
              <div onClick={handleNotificationIconClick} className="icon-button-container">
                <img src={notificationsIcon} id="notifications-icon" alt="Notificações" className="icon" />
              </div>
              <div onClick={handleUserIconClick} className="icon-button-container">
                  <img src={userIcon} id="user-icon" alt="Usuário" className="icon" />
              </div>
            </div>
          </div>
        </header>
  
        <main className="home-container">
          <section className="background-container">
         
          </section>
        </main>
      </section>
      </>
    );
  };

  export default CardPage;