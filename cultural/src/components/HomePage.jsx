// src/components/HomePage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import Header from './Layout/Header';
import Footer from './Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/home.css';

const HomePage = () => {
  const navigate = useNavigate();

  const [events, setEvents] = useState([]);
  const [attractions, setAttractions] = useState([]);
  const [isConfirmModalOpen, setConfirmModalOpen] = useState(false);

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
    const fetchCulturals = async () => {
      try {
        const response = await fetch('http://localhost:8080/culturais/home');
        const data = await response.json();
        setEvents(data.events || []);
        setAttractions(data.touristAttractions || []);
      } catch (error) {
        console.error('Erro ao buscar culturais:', error);
      }
    };
    fetchCulturals();
  }, []);

  const closeConfirmModal = () => {
    setConfirmModalOpen(false);
  };

  const handleConfirmLogout = () => {
    navigate('/');
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

      <section className="screen" id="home-page">
        <Header
          onLogoutClick={() => setConfirmModalOpen(true)}
          onNotificationClick={fetchNotifications}
        />

        <main className="home-container">
          <section className="background-container">
            <div className="category-box">
              <h2>Principais Eventos</h2>
              <div className="card-events">
                {events.map((event) => (
                  <Link to={`/card/event/${event.id}`} state={{ event: true }} key={event.id}>
                    <div className="card">
                      <p className="title">{event.title}</p>
                      <img
                        src={`http://localhost:8080/static/culturalthumbs/${event.image}`}
                        alt={event.title}
                      />
                    </div>
                  </Link>
                ))}
              </div>
              <h2>Principais Pontos Turísticos</h2>
              <div className="card-attractions">
                {attractions.map((attraction) => (
                  <Link
                    to={`/card/tourist_attraction/${attraction.id}`}
                    state={{ event: false }}
                    key={attraction.id}
                  >
                    <div className="card">
                      <p className="title">{attraction.title}</p>
                      <img
                        src={`http://localhost:8080/static/culturalthumbs/${attraction.image}`}
                        alt={attraction.title}
                      />
                    </div>
                  </Link>
                ))}
              </div>
            </div>
          </section>
        </main>
        <Footer userType={userType} />
      </section>
    </>
  );
};

export default HomePage;
