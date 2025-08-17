import { useUser } from '../contexts/UserContext';
import { useNavigate } from 'react-router-dom';

const Header = () => {
  const { userType, user, logout } = useUser();
  const navigate = useNavigate();

  const handleUserIconClick = () => {
    navigate('/profile');
  };

  const handleAddCultural = () => {
    // Implementar lógica para adicionar evento/atração
    console.log('Adicionar Cultural');
  };

  const handleLogout = () => {
    logout();
    navigate('/');
  };

  return (
    <header className={`top-bar ${userType === 'organizer' ? 'top-bar-organizer' : 'top-bar-common'}`}>
      <div className="logo-container">
        <img src="/logo6.png" alt="Logo Cultural" className="logo-img" />
      </div>

      <div className="right-section">
        <div className="icons">
          {userType === 'organizer' && (
            <button onClick={handleAddCultural} className="add-btn">
              Adicionar Cultural
            </button>
          )}
          
          <img 
            src="/images/notifications-icon.png" 
            alt="Notificações" 
            className="icon" 
          />
          
          <button onClick={handleUserIconClick} className="icon user-icon-btn">
            <img 
              src="/images/user-icon.png" 
              alt="Usuário" 
              className="icon" 
            />
          </button>

          {user && (
            <button onClick={handleLogout} className="logout-btn">
              Sair
            </button>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;

