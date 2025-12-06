// src/components/CardPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import NotificationModal from '../components/NotificationModal/NotificationModal';
import ConfirmModal from '../components/ConfirmModal/ConfirmComment';
import Header from '../components/Layout/Header';
import Footer from '../components/Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/card.css';
import './styles/favorites.css';
import './styles/comment.css';

const CommentPage = () => {
  const navigate = useNavigate();

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

  const { id, culturalType } = useParams();

  const [error, setError] = useState('');

  const [culturalData, setCulturalData] = useState(null);

  const [comment, setComment] = useState('');

  useEffect(() => {
    const fetchCulturalDataAndFavorites = async () => {
      console.log(`%cuseEffect ACIONADO em ${new Date().toLocaleTimeString()}`, 'color: orange');

      try {
        console.log(`Fetching cultural details for culturalType: ${culturalType}, id: ${id}`);
        const culturalResponse = await fetch(
          `http://localhost:8080/culturais/${culturalType}/${id}`
        );
        const cultural = await culturalResponse.json();
        console.log('Cultural data received:', cultural);
        setCulturalData({ ...cultural });
      } catch (error) {
        console.error('Erro ao buscar dados culturais:', error);
      }
    };
    fetchCulturalDataAndFavorites();
  }, [id, culturalType, userID]);

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

        <main className="home-container">
          <section className="main-content">
            <div key={culturalData.id} className="favorite-card">
              <img
                src={`http://localhost:8080/static/culturalthumbs/${culturalData.image}`}
                alt={culturalData.title}
                className="favorite-img"
              />
              <div className="favorite-details">
                <h3>{culturalData.title}</h3>
                <p>{culturalData.type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
                <div className="details-box">
                  <span>{culturalData.location}</span>
                  <br />

                  <span>
                    {culturalType === 'event' &&
                      culturalData.event &&
                      (culturalData.event.end_date === ''
                        ? ` ${culturalData.event.start_date}, de ${culturalData.event.working_hours}`
                        : ` ${culturalData.event.start_date} - ${culturalData.event.end_date}, de ${culturalData.event.working_hours}`)}
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
              <button onClick={handleAddCommentClick} className="down-btn">
                Adicionar
              </button>
            </div>
          </div>
        </main>
        <Footer userType={userType} />
      </section>
    </>
  );
};

export default CommentPage;
