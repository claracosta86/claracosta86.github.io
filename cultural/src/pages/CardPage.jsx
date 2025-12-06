// src/components/CardPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import NotificationModal from '../components/NotificationModal/NotificationModal';
import ConfirmModal from '../components/ConfirmModal/ConfirmComment';
import Header from '../components/Layout/Header';
import Footer from '../components/Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/card.css';
import favoriteIcon from '../assets/favorite-icon.png';
import unfavoriteIcon from '../assets/unfavorite-icon.png';
import locationIcon from '../assets/location-icon.png';
import clockIcon from '../assets/clock-icon.png';
import priceIcon from '../assets/price-icon.png';
import accessibleIcon from '../assets/accessibility-icon.png';
import mailIcon from '../assets/mail-icon.png';

const CardPage = () => {
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

  const [culturalData, setCulturalData] = useState(null);
  const [comments, setComments] = useState([]);

  useEffect(() => {
    const fetchCulturalDataAndFavorites = async () => {
      try {
        console.log(`Fetching cultural details for culturalType: ${culturalType}, id: ${id}`);
        const culturalResponse = await fetch(
          `http://localhost:8080/culturais/${culturalType}/${id}`
        );
        const cultural = await culturalResponse.json();

        console.log(`Fetching favorites for userID: ${userID}`);
        const favoritesResponse = await fetch(`http://localhost:8080/users/${userID}/favorites/`);
        const favoritesData = await favoritesResponse.json();

        const isFav = favoritesData.some((fav) => fav.id == cultural.id);

        if (isFav) {
          try {
            const updateLastSeen = await fetch(
              `http://localhost:8080/users/${userID}/favorites/last-seen`,
              {
                method: 'PATCH',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ culturalID: cultural.id, culturalType: culturalType }),
                credentials: 'include',
              }
            );
            if (!updateLastSeen.ok) {
              console.error('Falha ao atualizar último visto.');
            }
          } catch (error) {
            console.error('Erro ao atualizar último visto:', error);
          }
        }

        console.log('Cultural details fetched:', cultural);
        console.log('User favorites fetched:', favoritesData);
        console.log(`Is favorite: ${isFav}`);
        setCulturalData({
          ...cultural,
          isFavorite: isFav,
        });
      } catch (error) {
        console.error('Erro ao buscar detalhes ou favoritos:', error);
      }
    };
    const fetchCulturalComment = async () => {
      try {
        const response = await fetch(`http://localhost:8080/comments/${culturalType}/${id}`);
        const data = await response.json();
        console.log('Comentários culturais recebidos:', data);
        setComments(data.comments);
      } catch (error) {
        console.error('Erro ao buscar comentários culturais:', error);
      }
    };
    fetchCulturalDataAndFavorites();
    fetchCulturalComment();
  }, [id, culturalType, userID]);

  const handleFavoriteIconClick = async () => {
    if (!culturalData) return;
    try {
      const response = await fetch(`http://localhost:8080/users/${userID}/favorites`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          isFavorite: !culturalData.isFavorite,
          culturalType: culturalType,
          culturalID: culturalData.id,
        }),
        credentials: 'include',
      });
      if (response.ok) {
        setCulturalData((prevData) => ({
          ...prevData,
          isFavorite: !prevData.isFavorite,
        }));
      }
    } catch (error) {
      console.error('Erro ao atualizar favorito:', error);
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

  const getWorkingHoursLines = (wh) => {
    if (!wh) return [];

    const dias = [
      'Domingo',
      'Segunda-feira',
      'Terça-feira',
      'Quarta-feira',
      'Quinta-feira',
      'Sexta-feira',
      'Sábado',
    ];

    const regexDias = new RegExp(`(${dias.join('|')})`, 'g');
    let s = wh.trim();

    s = s.replace(regexDias, (match) => `\n${match}`);

    s = s.trim();

    return s
      .split('\n')
      .map((linha) => linha.trim())
      .filter(Boolean);
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
            <div className="details-container">
              <div className="header-details">
                <h2 className="cultural-title">{culturalData.title}</h2>
              </div>
              <img
                src={`http://localhost:8080/static/culturalthumbs/${culturalData.image}`}
                alt={culturalData.title}
                className="event-image"
              />

              <div className="info-box">
                <p>
                  <img src={locationIcon} alt="Localização" className="info-icon" />
                  <strong>Endereço:</strong> {culturalData.location}
                </p>
                <p>
                  <img src={clockIcon} alt="Horário" className="info-icon" />
                  {culturalType === 'event' ? (
                    <strong>Data e Horário:</strong>
                  ) : (
                    <strong>Horário de Funcionamento:</strong>
                  )}
                  <br />
                  {culturalType === 'event' &&
                    culturalData.event &&
                    (culturalData.event.endDate === ''
                      ? ` ${culturalData.event.startDate}, de ${culturalData.event.durationHours}`
                      : ` ${culturalData.event.startDate} - ${culturalData.event.endDate}, de ${culturalData.event.durationHours}`)}
                  {culturalType !== 'event' && culturalData.touristAttraction && (
                    <span>
                      {getWorkingHoursLines(culturalData.touristAttraction.workingHours).map(
                        (linha, idx) => (
                          <div key={idx}>{linha}</div>
                        )
                      )}
                    </span>
                  )}
                </p>
                <p>
                  <img src={priceIcon} alt="Preço" className="info-icon" />
                  <strong>Preço: </strong>
                  {culturalData.price === 'R$0,00' || culturalData.price === 'Gratuito'
                    ? 'Gratuito'
                    : `${culturalData.price}`}
                </p>
                <p>
                  <img src={accessibleIcon} alt="Acessível" className="info-icon" />
                  <strong>Acessível:</strong> {culturalData.isAccessible ? 'Sim' : 'Não'}
                </p>
                <p>
                  <img src={mailIcon} alt="Contato" className="info-icon" />
                  <strong>Contato:</strong> {culturalData.organizer.email}
                </p>
                <p className="description">
                  <strong>Descrição:</strong> {culturalData.description}
                </p>
              </div>
            </div>

            <div id="organizer-container">
              <Link 
                to={`/organizer/${culturalData.organizer.id}`} 
                className="see-organizer-btn"
              >
                Conhecer Organizador
              </Link>
            </div>

            <div className="comments-section">
              <h3>Comentários</h3>
              <div className="comment-box">
                {comments != null ? (
                  comments.map((comment) => (
                    <div key={comment.id} className="comment-item">
                      <p>
                        <strong>{comment.userName}:</strong> {comment.comment}
                      </p>
                    </div>
                  ))
                ) : (
                  <p>Não há comentários ainda. Seja o primeiro a comentar!</p>
                )}
              </div>
              <Link
                to={`/comments/${culturalType}/${culturalData.id}`}
                className="add-comment-btn"
              >
                Adicionar Comentário
              </Link>
            </div>
          </section>
          <div className="down-actions-container">
            <div className="down-actions-row">
              <button onClick={handleGoBackClick} className="down-btn">
                Voltar
              </button>
              <button
                onClick={handleFavoriteIconClick}
                src={culturalData.isFavorite ? favoriteIcon : unfavoriteIcon}
                alt="Favoritar"
                className="down-btn"
              >
                {culturalData.isFavorite ? '	Desfavoritar ♡' : 'Favoritar ❤'}
              </button>
            </div>
          </div>
        </main>
        <Footer userType={userType} />
      </section>
    </>
  );
};

export default CardPage;
