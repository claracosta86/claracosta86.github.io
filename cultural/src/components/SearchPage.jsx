// src/components/SearchPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import NotificationModal from './NotificationModal/NotificationModal';
import ConfirmModal from './ConfirmModal/ConfirmComment';
import Header from './Layout/Header';
import Footer from './Layout/Footer';
import { useNotifications } from '../hooks/useNotifications';
import './styles/home.css';
import './styles/search.css';
import searchIcon from '../assets/search-icon.png';

const SearchPage = () => {
  const navigate = useNavigate();

  const [query, setQuery] = useState('');
  const [results, setResults] = useState([]);
  const [error, setError] = useState('');

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
    if (userID) {
      console.log('UserID recebido da página de login:', userID);
    }
    if (userType) {
      console.log('UserType recebido da página de login:', userType);
    }
  }, [user]);

  const handleSearch = async () => {
    if (!query.trim()) {
      setError('Por favor, digite um termo para buscar.');
      setResults([]);
      return;
    }
    setError('');
    try {
      const response = await fetch(`http://localhost:8080/culturais/search?q=${query}`);
      if (response.ok) {
        const data = await response.json();
        setResults(data || []);
        if (!data || data.length === 0) {
          setError('Nenhum resultado encontrado para sua busca.');
        }
      } else {
        setError('Erro ao buscar resultados. Tente novamente.');
      }
    } catch (err) {
      setError('Erro de conexão. Verifique sua rede e tente novamente.');
      console.error(err);
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

        <main className="search-container">
          <section className="background-container">
            <div className="category-box">
              <h2>Explore os eventos e atrações:</h2>
              <div className="search-box">
                <input
                  type="text"
                  value={query}
                  onChange={(e) => setQuery(e.target.value)}
                  placeholder="Buscar por eventos ou locais..."
                  className="search-input"
                />
                <button onClick={handleSearch} className="search-button">
                  <img src={searchIcon} alt="Buscar" />
                </button>
              </div>

              {error && <p className="error-message">{error}</p>}

              <div className="results-container">
                {results.map((item) => (
                  <Link
                    to={`/card/${item.id}`}
                    state={{ userID, userType, event: item.type === 'event' }}
                    key={item.id}
                    className="result-card"
                  >
                    <img
                      src={`http://localhost:8080/static/culturalthumbs/${item.image}`}
                      alt={item.title}
                      className="result-image"
                    />
                    <div className="result-details">
                      <h3>{item.title}</h3>
                      <p>{item.type === 'event' ? 'Evento' : 'Ponto Turístico'}</p>
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

export default SearchPage;
