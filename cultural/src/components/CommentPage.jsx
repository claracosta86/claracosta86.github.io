// src/components/CardPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import './styles/card.css';
import './styles/favorites.css';
import './styles/comment.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import homeIcon from '../assets/home-icon.png';
import addIcon from '../assets/add-icon.png';
import searchIcon from '../assets/search-icon.png';
import favoriteIcon from '../assets/favorite-icon.png';
import userIcon from '../assets/user-icon.png';
import logoutIcon from '../assets/logout-icon.png';

const CommentPage = () => {
  const navigate = useNavigate();

  const [notification, setNotification] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

  const { id, culturalType } = useParams();

  const [error, setError] = useState('');

  const [culturalData, setCulturalData] = useState(null);

  const [comment, setComment] = useState('');

  const { user } = useUser();
  const userID = user.userID;
  const userType = user.type;

  useEffect(() => {
    const fetchCulturalDataAndFavorites = async () => {
      console.log(`%cuseEffect ACIONADO em ${new Date().toLocaleTimeString()}`, 'color: orange');

      try {
        console.log(`Fetching cultural details for culturalType: ${culturalType}, id: ${id}`);
        const culturalResponse = await fetch(`http://localhost:8080/culturais/${culturalType}/${id}`);
        const cultural = await culturalResponse.json();
        console.log('Cultural data received:', cultural);
        setCulturalData({...cultural});
      } catch (error) {
        console.error('Erro ao buscar dados culturais:', error);
      }
    };
    fetchCulturalDataAndFavorites();
  }, [id, culturalType, userID]);

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
    if (notification.length === 0) {
      setNotificationModalOpen(false);
      return;
    }
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

  const handleAddCommentClick = async () => {
    if (comment.trim() === '') {
      setError('Seu comentário não pode estar vazio.');
      return;
    }

    try {
      const response = await fetch(`http://localhost:8080/comments/`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          cultural_type: culturalType,
          culturalID: culturalData.id,
          user_id: userID,
          comment: comment,
        }),
        credentials: 'include',
      });
      if (response.ok) {
        console.log('Comentário adicionado com sucesso.');
        navigate(-1); 
      }
    } catch (error) {
      console.error('Erro ao adicionar comentário:', error);
    }
  };

  if (!culturalData) {
    return <div>Carregando...</div>;
  }

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

        <main className="home-container">
          <section className="main-content">
            <div key={culturalData.id} className="favorite-card">
                <img
                  src={`/thumb-size/${culturalData.image}`}
                  alt={culturalData.title}
                  className="favorite-img"
                />
              <div className="favorite-details">
                <h3>{culturalData.title}</h3>
                <p>{culturalData.type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
                <div className="details-box">
                  <span>{culturalData.location}</span><br />

                  <span>
                      {culturalType === 'event' && culturalData.event && (
                        culturalData.event.end_date === "" 
                          ? ` ${culturalData.event.start_date}, de ${culturalData.event.working_hours}`
                          : ` ${culturalData.event.start_date} - ${culturalData.event.end_date}, de ${culturalData.event.working_hours}`
                      )}
                      {culturalType !== 'event' &&
                        culturalData.touristAttraction &&
                        ` ${culturalData.touristAttraction.open_days}, ${culturalData.touristAttraction.open_time}`}
                  </span>

                  <span className="price">
                      {culturalData.price === 'R$0,00' || culturalData.price === 'Gratuito'
                        ? 'Gratuito'
                        : `${culturalData.price}`}
                  </span>
                </div>
              </div>
            </div> 

            <div className="comments-section">
              <div className="navigation-header">
                <h2>Adicionar Comentário</h2>
              </div>
              <textarea
                className="comment-text-box"
                name="comment"
                placeholder="Escreva seu comentário aqui..." 
                value={comment}
                onChange={(e) => setComment(e.target.value)}
              />

              {error && <span className="error">{error}</span>}
            </div>
          </section>
          <div className="down-actions-container">
            <div className="down-actions-row">
              <button onClick={handleGoBackClick} className="down-btn">
                Voltar
              </button>
              <button
                onClick={handleAddCommentClick}
                className="down-btn"
              >
                Adicionar Comentário
              </button>
            </div>
          </div>
        </main>
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

export default CommentPage;
