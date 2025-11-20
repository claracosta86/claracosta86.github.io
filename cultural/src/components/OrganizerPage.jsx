// src/components/OrganizerPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import Header from './Layout/Header';
import Footer from './Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/organizer.css';

const OrganizerPage = () => {
  const navigate = useNavigate();

  const { user } = useUser();
  const userID = user.userID;
  const userType = user.type;

  const { id: organizerID } = useParams();

  const {
    notifications,
    isNotificationModalOpen,
    setNotificationModalOpen,
    fetchNotifications,
    markNotificationsAsSeen,
  } = useNotifications(userID);

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

  const [organizerData, setOrganizerData] = useState({
    name: '',
    email: '',
    organizerSince: '',
    id: organizerID,
  });
  const [culturais, setCulturais] = useState([]);

  useEffect(() => {
    const fetchOrganizerData = async () => {
      if (!userID) return;

      try {
        const response = await fetch(
          `http://localhost:8080/users/${organizerID}/culturais/organizer`
        );
        if (!response.ok) {
          console.error('Falha ao buscar dados do organizador.');
          return;
        }

        const data = await response.json();
        setOrganizerData({
          name: data.name,
          email: data.email,
          organizerSince: data.organizerSince,
          id: data.id,
        });

        if (data.culturalItems && data.culturalItems.length > 0) {
          const culturaisWithDetails = await Promise.all(
            data.culturalItems.map(async (cult) => {
              if (cult.type && cult.id) {
                const detailResponse = await fetch(
                  `http://localhost:8080/culturais/${cult.type}/${cult.id}`
                );
                if (detailResponse.ok) {
                  const detailData = await detailResponse.json();
                  return {
                    ...cult,
                    Title: detailData.title,
                    Image: detailData.image,
                    Location: detailData.location,
                    Price: detailData.price,
                    Event: cult.type === 'event' ? detailData.event : null,
                    TouristAttraction:
                      cult.type === 'tourist_attraction' ? detailData.touristAttraction : null,
                  };
                }
              }
              return cult;
            })
          );
          setCulturais(culturaisWithDetails);
        } else {
          setCulturais([]);
        }
      } catch (error) {
        console.error('Erro na requisição:', error);
      }
    };

    fetchOrganizerData();
  }, [userID, organizerID]);

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
            <h2>Conheça o Organizador</h2>
          </div>

          <div key={organizerData.id} className="organizer-info">
            <p>
              <b>Nome:</b> {organizerData.name}
            </p>
            <p>
              <b>Contato:</b> {organizerData.email}
            </p>
            <p>
              <b>Tempo na plataforma:</b> {organizerData.organizerSince}
            </p>
          </div>
          <p className="separator-unique"></p>
          <div className="favorites-list">
            <h3>Culturais do Organizador</h3>
            {culturais.length > 0 ? (
              culturais.map(
                (cult) =>
                  cult.Title &&
                  cult.Image &&
                  cult.Location && (
                    <div key={cult.id} className="favorite-card">
                      <Link
                        to={`/card/${cult.type}/${cult.id}`}
                        state={{ userID, userType, isEvent: cult.type === 'event' }}
                        className="card-link"
                      >
                        <img
                          src={`/thumb-size/${cult.Image}`}
                          alt={cult.Title}
                          className="favorite-img"
                        />
                        <div className="favorite-details">
                          <h3>{cult.Title}</h3>
                          <p>{cult.type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
                          <div className="details-box">
                            <span id="Working Hours">
                              {cult.type === 'event' &&
                                cult.Event &&
                                (cult.Event.endDate === ''
                                  ? ` ${cult.Event.startDate}, de ${cult.Event.durationHours}`
                                  : ` ${cult.Event.startDate} - ${cult.Event.endDate}, de ${cult.Event.durationHours}`)}
                              {cult.type !== 'event' &&
                                cult.TouristAttraction &&
                                ` ${cult.TouristAttraction.workingHours}`}
                            </span>
                            <span className="price">
                              {cult.Price === 'R$0,00' || cult.Price === 'Gratuito'
                                ? 'Gratuito'
                                : `${cult.Price}`}
                            </span>
                          </div>
                        </div>
                      </Link>
                    </div>
                  )
              )
            ) : (
              <p>
                <em>Você ainda não adicionou nenhum cultural.</em>
              </p>
            )}
          </div>
          <div className="profile-actions">
            <div className="profile-actions-row">
              <button onClick={handleGoBackClick} className="profile-btn">
                Voltar
              </button>
            </div>
          </div>
        </div>
        <Footer userType={userType} />
      </section>
    </>
  );
};

export default OrganizerPage;
