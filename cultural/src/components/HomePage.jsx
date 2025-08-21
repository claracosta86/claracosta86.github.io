// src/components/HomePage.jsx
import { useState, useEffect } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import './styles/home.css'; // Importa o CSS da página home
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';
import searchIcon from '../assets/search-icon.png';
import bienalEvent from '../assets/thumb-size/bienal-event.png';
import mcrEvent from '../assets/thumb-size/mcr-event.png';
import dccWeekEvent from '../assets/thumb-size/dccweek-event.png';
import iwnbEvent from '../assets/thumb-size/iwnb-event.png';
import cruEvent from '../assets/thumb-size/cru-event.png';
import liberdadeAttraction from '../assets/thumb-size/liberdade-attraction.png';
import igrejinhaAttraction from '../assets/thumb-size/igrejinha-attraction.png';
import pseteAttraction from '../assets/thumb-size/psete-attraction.png';
import mercadoAttraction from '../assets/thumb-size/mercado-attraction.png';
import mangabeirasAttraction from '../assets/thumb-size/mangabeiras-attraction.png';


const NotificationModal = ({ isOpen, onClose, notifications }) => {
  if (!isOpen) return null;

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
                <p> Veja as atualizações de {notif.Name}</p>
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


const HomePage = () => {
  const navigate = useNavigate();

  const [userType, setUserType] = useState('common');
  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);
  
  const location = useLocation();
  const userID = location.state?.userID || '';

  useEffect(() => {
      const fetchUserType = async () => {
        try {
          const response = await fetch("http://localhost:8080/users/get-type", {
               credentials: 'include'
          });
          if (response.ok) {
            const data = await response.json();
            setUserType(data.userType);
          }
        } catch (error) {
          console.error("Erro ao buscar o tipo de usuário:", error);
        }
      };
      fetchUserType();
    }, []);

  useEffect(() => {
    if (userID) {
      console.log("UserID recebido da página de login:", userID);
    }
  }, [userID]);

  const handleUserIconClick = async () => {
    navigate('/user/profile');
  };

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

      <section className="search-bar">
        <img src={searchIcon} alt="Buscar" className="search-icon" />
        <input id="search-input" type="text" placeholder="Buscar..." />
      </section>

      <main className="home-container">
        <section className="background-container">
          <div className="category-box">
            <h2>Principais Eventos</h2>
            <div className="card-events">
              <div className="card">
                <p className="title">Bienal do Livro</p>
                <img src={bienalEvent} alt="Bienal do livro" />
              </div>
              <div className="card">
                <p className="title">My Chemical Romance Ao Vivo</p>
                <img src={mcrEvent} alt="MCR Ao Vivo" />
              </div>
              <div className="card">
                <p className="title">DCC Week</p>
                <img src={dccWeekEvent} alt="DCC Week" />
              </div>
              <div className="card">
                <p className="title">I Wanna Be Tour</p>
                <img src={iwnbEvent} alt="I Wanna Be Tour" />
              </div>
              <div className="card">
                <p className="title">Jogo do Cruzeiro</p>
                <img src={cruEvent} alt="Jogo do Cruzeiro" />
              </div>
            </div>
          </div>
          <div className="category-box">
            <h2>Principais Pontos Turísticos</h2>
            <div className="card-attractions">
              <div className="card">
                <p className="title">Praça Liberdade</p>
                <img src={liberdadeAttraction} alt="Praça da Liberdade" />
              </div>
              <div className="card">
                <p className="title">Igreja da Pampulha</p>
                <img src={igrejinhaAttraction} alt="Igreja da Pampulha" />
              </div>
              <div className="card">
                <p className="title">Pirulito da Praça Sete</p>
                <img src={pseteAttraction} alt="Pirulito da Praça Sete" />
              </div>
              <div className="card">
                <p className="title">Mercado Central</p>
                <img src={mercadoAttraction} alt="Mercado Central" />
              </div>
              <div className="card">
                <p className="title">Parque das Mangabeiras</p>
                <img src={mangabeirasAttraction} alt="Parque das Mangabeiras" />
              </div>
            </div>
          </div>
        </section>
      </main>
    </section>
    </>
  );
};

export default HomePage;