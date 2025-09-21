// src/components/CardPage.jsx
import { useState, useEffect } from 'react';
import { Link, useNavigate, useLocation, useParams } from 'react-router-dom';
import './styles/card.css'; 
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import unfavoriteIcon from '../assets/unfavorite-icon.png';
import gobackIcon from '../assets/goback.png';
import userIcon from '../assets/user-icon.png';
import locationIcon from '../assets/location-icon.png';
import clockIcon from '../assets/clock-icon.png';
import priceIcon from '../assets/price-icon.png';
import accessibleIcon from '../assets/accessibility-icon.png';
import mailIcon from '../assets/mail-icon.png';

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
  console.log("Rendering CardPage");
  const navigate = useNavigate();

  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);
  
  const { id } = useParams();
  const [culturalData, setCulturalData] = useState(null);

  const location = useLocation();
  const userID = location.state?.userID || '';
  const userType = location.state?.userType || '';
  let type = "event" ;
  if (location.state?.event == false) {
    type = "attraction";
  }

  useEffect(() => {
   const fetchCulturalDetails = async () => {
            try {
              console.log(`Fetching cultural details for type: ${type}, id: ${id}`);
                const response = await fetch(`http://localhost:8080/cultural/${type}/${id}`);
                const data = await response.json();
                setCulturalData(data);
            } catch (error) {
                console.error("Erro ao buscar detalhes do evento:", error);
            }
        };
        fetchCulturalDetails();
    }, [id, type]); 

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

  const handleFavoriteIconClick = async () => {
    if (!culturalData) return;
    try {
      const response = await fetch(`http://localhost:8080/users/${userID}/favorites`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ isFavorite: !culturalData.isFavorite, culturalType: type, culturalID: culturalData.ID }),
        credentials: 'include'
      });
      if (response.ok) {
        setCulturalData(prevData => ({
          ...prevData,
          isFavorite: !prevData.isFavorite
        }));
      }
    } catch (error) { 
      console.error("Erro ao atualizar favorito:", error);
    }
  };

  if (!culturalData) {
    return <div>Carregando...</div>;
  }
  
  const topBarClass = userType === 'organizer' ? 'top-bar-organizer' : 'top-bar-common';


    return (
      <>
      <NotificationModal isOpen={isNotificationModalOpen} onClose={() => setNotificationModalOpen(false)} notifications={notification} />
      <section className="screen" id="tela-home">
        <header className={topBarClass}>
          <div className="logo-container">
            <Link to="/home">
              <img src={logo} alt="Logo Cultural" className="logo-tiny" />
            </Link>
          </div>
          <div className="right-section">
            <div className="icons">
              {userType === 'organizer' && (
                <a href="" className="add-btn">Adicionar Cultural</a>
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
          <section>
            <div className="details-container">
                  <div className="header-details">
                      <button onClick={() => navigate(-1)} className="back-button">
                          <img src={gobackIcon} alt="Voltar" className="goback-img" />
                      </button>
                      <h2 className="cultural-title">{culturalData.title}</h2>
                      <img onClick={handleFavoriteIconClick} src={culturalData.isFavorite ? favoriteIcon : unfavoriteIcon} alt="Favoritar" className="favorite-icon" />
                  </div>
                  <img src={`/thumb-size/${culturalData.image}`} alt={culturalData.title} className="event-image" />

                  <div className="info-box">
                      <p><img src={locationIcon} alt="Localização" className="info-icon" /><strong>Endereço:</strong> {culturalData.location}</p>
                      <p><img src={clockIcon} alt="Horário" className="info-icon" /><strong>Horário de Funcionamento:</strong> {culturalData.duration}</p>
                      <p><img src={priceIcon} alt="Preço" className="info-icon" /><strong>Preço: R$</strong> {culturalData.price}</p>
                      <p><img src={accessibleIcon} alt="Acessível" className="info-icon" /><strong>Acessível:</strong> {culturalData.accessible ? 'Sim' : 'Não'}</p>
                      <p><img src={mailIcon} alt="Contato" className="info-icon" /><strong>Contato:</strong> {culturalData.organizer.email}</p>
                      <p className="description"><strong>Descrição:</strong> {culturalData.description}</p>
                  </div>
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
            </section>
        </main>
      </section>
      </>
    );
  };

  export default CardPage;