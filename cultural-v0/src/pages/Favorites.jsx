import { useState, useEffect } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { useUser } from '../contexts/UserContext';
import Card from '../components/Card';

const Favorites = () => {
  const { user } = useUser();
  const navigate = useNavigate();
  const [favorites, setFavorites] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    if (!user) {
      navigate('/login');
      return;
    }

    // Simular carregamento de favoritos - em produção viria de uma API
    const loadFavorites = async () => {
      setIsLoading(true);
      try {
        await new Promise(resolve => setTimeout(resolve, 1000));
        
        // Mock de favoritos
        const mockFavorites = [
          {
            id: 1,
            title: 'Bienal do Livro',
            image: '/images/thumb-size/bienal-event.png',
            alt: 'Bienal do livro',
            type: 'event'
          },
          {
            id: 2,
            title: 'Praça Liberdade',
            image: '/images/thumb-size/liberdade-attraction.png',
            alt: 'Praça da Liberdade',
            type: 'attraction'
          },
          {
            id: 3,
            title: 'My Chemical Romance Ao Vivo',
            image: '/images/thumb-size/mcr-event.png',
            alt: 'MCR Ao Vivo',
            type: 'event'
          }
        ];
        
        setFavorites(mockFavorites);
      } catch (error) {
        console.error('Erro ao carregar favoritos:', error);
      } finally {
        setIsLoading(false);
      }
    };

    loadFavorites();
  }, [user, navigate]);

  const removeFavorite = (id) => {
    setFavorites(prev => prev.filter(fav => fav.id !== id));
  };

  if (!user) {
    return null;
  }

  return (
    <div className="favorites-container">
      <div className="favorites-header">
        <Link to="/profile" className="back-btn">
          <img src="/images/goback-icon.png" alt="Voltar" />
        </Link>
        <h1>Meus Favoritos</h1>
      </div>

      <div className="favorites-content">
        {isLoading ? (
          <div className="loading">
            <p>Carregando favoritos...</p>
          </div>
        ) : favorites.length === 0 ? (
          <div className="empty-state">
            <p>Você ainda não tem favoritos.</p>
            <Link to="/home" className="browse-btn">
              Explorar eventos e atrações
            </Link>
          </div>
        ) : (
          <div className="favorites-grid">
            {favorites.map(favorite => (
              <div key={favorite.id} className="favorite-item">
                <Card
                  id={favorite.id}
                  title={favorite.title}
                  image={favorite.image}
                  alt={favorite.alt}
                  type={favorite.type}
                />
                <button
                  onClick={() => removeFavorite(favorite.id)}
                  className="remove-favorite-btn"
                  title="Remover dos favoritos"
                >
                  ❌
                </button>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default Favorites;

