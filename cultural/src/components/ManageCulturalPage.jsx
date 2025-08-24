// src/components/ManageCulturalPage.jsx
import { useEffect } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import './styles/profile.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';
import gobackIcon from '../assets/goback.png';

const ManageCulturalPage = () => {
  const navigate = useNavigate();

  const location = useLocation();
  const userID = location.state?.userID || '';
  const userType = location.state?.userType || 'common';
  
  useEffect(() => {
    if (userID) {
      console.log("UserID recebido:", userID);
    }
    if (userType) {
      console.log("UserType recebido:", userType);
    }
  }, [userID, userType]);
  

  const headerClass = userType === 'organizer' ? 'top-bar-organizer' : 'top-bar-common';

  return (
    <section className="screen" id="tela-profile">
      <header className={headerClass}>
        <div className="logo-container">
          <Link to="/home">
            <img src={logo} alt="Logo Cultural" className="logo-tiny" />
          </Link>
        </div>
        <div className="right-section">
          <div className="icons">
            {userType === 'organizer' && (
              <a href="#" className="add-btn">Adicionar Cultural</a>
            )}
            <img src={notificationsIcon} alt="Notificações" className="icon" />
            <Link to="/user/profile">
              <img src={userIcon} alt="Usuário" className="icon" />
            </Link>
          </div>
        </div>
      </header>
      <div className="profile-box">
        <div className="header-title">
          <Link to="/user/profile" state={{ userType, userID }}>
            <img src={gobackIcon} alt="Go Back Arrow" className="goback-img" />
          </Link>
          <h2>Meus Culturais</h2>
        </div>
        {/* Aqui você pode adicionar a lógica para listar os favoritos */}
      </div>
    </section>
  );
};

export default ManageCulturalPage;