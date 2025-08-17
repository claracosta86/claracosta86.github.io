import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useUser } from '../contexts/UserContext';
import Header from '../components/Header';
import SearchBar from '../components/SearchBar';
import Card from '../components/Card';
import '../styles/home.css';

const Home = () => {
  const { user, userType } = useUser();
  const navigate = useNavigate();
  const [searchTerm, setSearchTerm] = useState('');

  // Mock data - em produção viria da API
  const [events] = useState([
    { id: 1, title: 'Bienal do Livro', image: '/images/thumb-size/bienal-event.png', type: 'event' },
    { id: 2, title: 'DCC Week', image: '/images/thumb-size/dccweek-event.png', type: 'event' },
    { id: 3, title: 'IWNB', image: '/images/thumb-size/iwnb-event.png', type: 'event' },
    { id: 4, title: 'MCR', image: '/images/thumb-size/mcr-event.png', type: 'event' },
    { id: 5, title: 'PSETE', image: '/images/thumb-size/psete-event.png', type: 'event' },
    { id: 6, title: 'CRU', image: '/images/thumb-size/cru-event.png', type: 'event' }
  ]);

  const [attractions] = useState([
    { id: 1, title: 'Igrejinha da Pampulha', image: '/images/thumb-size/igrejinha-attraction.png', type: 'attraction' },
    { id: 2, title: 'Liberdade', image: '/images/thumb-size/liberdade-attraction.png', type: 'attraction' },
    { id: 3, title: 'Mangabeiras', image: '/images/thumb-size/mangabeiras-attraction.png', type: 'attraction' },
    { id: 4, title: 'Mercado Central', image: '/images/thumb-size/mercado-attraction.png', type: 'attraction' }
  ]);

  const handleSearch = (term) => {
    setSearchTerm(term);
  };

  const filteredEvents = events.filter(event => 
    event.title.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const filteredAttractions = attractions.filter(attraction => 
    attraction.title.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="screen">
      <Header />
      
      <div className="home-container">
        <SearchBar onSearch={handleSearch} />
        
        <div className="background-container">
          <div className="category-box">
            <h2 className="title">Principais Eventos</h2>
            <div className="card-events">
              {filteredEvents.map(event => (
                <Card key={event.id} {...event} />
              ))}
            </div>
          </div>
          
          <div className="category-box">
            <h2 className="title">Principais Pontos Turísticos</h2>
            <div className="card-attractions">
              {filteredAttractions.map(attraction => (
                <Card key={attraction.id} {...attraction} />
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;

