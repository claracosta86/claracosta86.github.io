// src/components/ManageCulturalPage.jsx
import { useUser } from '../contexts/UserContext';
import RemoveModal from '../components/RemoveModal/RemoveFromApp';
import NotificationModal from '../components/NotificationModal/NotificationModal';
import ConfirmModal from '../components/ConfirmModal/ConfirmComment';
import Header from '../components/Layout/Header';
import Footer from '../components/Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './styles/profile.css';
import './styles/manage.css';

const ManageCulturalPage = () => {
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

  useEffect(() => {
    if (userID) {
      console.log('UserID recebido:', userID);
    }
    if (userType) {
      console.log('UserType recebido:', userType);
    }
  }, [userID, userType]);

  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);
  const [isRemoveModalOpen, setRemoveModalOpen] = useState(false);
  const [culturalToRemove, setCulturalToRemove] = useState(null);

  const handleGoBackClick = () => {
    navigate(-1);
  };

  const [culturais, setCulturais] = useState([]);
  useEffect(() => {
    const fetchCulturais = async () => {
      if (!userID) return;
      try {
        const response = await fetch(`http://localhost:8080/users/${userID}/culturais`);
        if (response.ok) {
          const data = await response.json();
          const culturaisWithDetails = await Promise.all(
            data.map(async (cult) => {
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
          setCulturais(culturaisWithDetails || []);
        } else {
          console.error('Falha ao buscar culturais.');
          setCulturais([]);
        }
      } catch (error) {
        console.error('Erro ao buscar culturais:', error);
      }
    };
    fetchCulturais();
  }, [userID]);

  const handleEditClick = (id, type) => {
    navigate(`/edit-cultural/${type}/${id}`);
  };

  const openRemoveModal = (cultural) => {
    setCulturalToRemove(cultural);
    setRemoveModalOpen(true);
  };

  const closeRemoveModal = () => {
    setRemoveModalOpen(false);
    setCulturalToRemove(null);
  };

  const handleConfirmRemove = async () => {
    if (!culturalToRemove) return;

    try {
      const { id, type } = culturalToRemove;
      const response = await fetch(`http://localhost:8080/culturais/${type}/${id}`, {
        method: 'DELETE',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
      });

      if (response.ok) {
        setCulturais((prevCulturais) => prevCulturais.filter((cult) => cult.id !== id));
      } else {
        console.error('Não foi possível excluir o cultural. Tente novamente.');
      }
    } catch (error) {
      console.error('Erro de rede ao remover cultural:', error);
    } finally {
      closeRemoveModal();
    }
  };

  const closeConfirmModal = () => {
    setConfirmModalOpen(false);
  };

  const handleConfirmLogout = () => {
    navigate('/');
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
        <RemoveModal
          isOpen={isRemoveModalOpen}
          onClose={closeRemoveModal}
          onConfirm={handleConfirmRemove}
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
            <h2>Meus Culturais</h2>
          </div>
          <div className="favorites-list">
            {culturais.length > 0 ? (
              culturais.map(
                (cult) =>
                  cult.Title && (
                    <div key={cult.id} className="manage-card">
                      <Link
                        to={`/card/${cult.type}/${cult.id}`}
                        state={{ userID, userType, isEvent: cult.type === 'event' }}
                        className="card-link"
                      >
                        <img
                          src={`http://localhost:8080/static/culturalthumbs/${cult.Image}`}
                          alt={cult.Title}
                          className="manage-img"
                        />
                        <div className="manage-details">
                          <h3>{cult.Title}</h3>
                          <p>{cult.type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
                          <div className="details-content">
                            <span>{cult.Location}</span> <br />
                            <span className="price">
                              {cult.Price === 'R$0,00' || cult.Price === 'Gratuito'
                                ? 'Gratuito'
                                : `${cult.Price}`}
                            </span>
                          </div>
                        </div>
                      </Link>
                      <div className="actions-row">
                        <button
                          onClick={() => handleEditClick(cult.id, cult.type)}
                          className="edit-btn"
                        >
                          Editar
                        </button>
                        <button onClick={() => openRemoveModal(cult)} className="remove-btn">
                          Excluir
                        </button>
                      </div>
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

export default ManageCulturalPage;
