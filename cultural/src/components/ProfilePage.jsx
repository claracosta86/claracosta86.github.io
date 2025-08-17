// src/components/ProfilePage.jsx
import { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './styles/profile.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';
import gobackIcon from '../assets/goback.png';

const ProfilePage = () => {
  const navigate = useNavigate();

  const [userType, setUserType] = useState('');
  const [userID, setUserID] = useState('');
  const [userName, setUserName] = useState('');
  const [userEmail, setUserEmail] = useState('');
  const [companyName, setCompanyName] = useState('');
  
  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await fetch("http://localhost:8080/user/get-information", {
          credentials: 'include'
        });
        if (response.ok) {
          const data = await response.json();
          setUserType(data.userType);
          setUserID(data.userID); 
        }
      } catch (error) {
        console.error("Erro ao buscar dados do usuário:", error);
        setUserType('common'); 
      }
    };
    fetchUserData();
  }, []);

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await fetch(`http://localhost:8080/users/${userID}/profile`, {
          credentials: 'include'
        });
        if (response.ok) {
          const data = await response.json();
          setUserEmail(data.email);
          setUserName(data.name);
          if (userType === "organizer") {
            setCompanyName(data.companyName);
          }
        }
      } catch (error) {
        console.error("Erro ao buscar dados do usuário:", error);
        setUserType('common'); 
      }
    };
    fetchUserData();
  }, [userID]);

  const handleLogout = () => {
    navigate('/');
  };

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
          <Link to="/home">
            <img src={gobackIcon} alt="Go Back Arrow" className="goback-img" />
          </Link>
          <h2>Seu Perfil</h2>
        </div>
        <label htmlFor="name">Nome</label>
        <input id="name" type="text" placeholder={userName} autoComplete="given-name" disabled />

        <label htmlFor="email">E-mail</label>
        <input id="email" type="email" placeholder={userEmail} autoComplete="email" disabled />

        {userType === 'organizer' && (
          <>
            <label htmlFor="companyName">Nome da Empresa</label>
            <input id="companyName" type="text" placeholder={companyName} autoComplete="organization" disabled />
          </>
        )}

        <div className="profile-actions">
          <Link to="/user/profile/edit" state={{ userType, userID }} className="profile-btn">Editar Perfil</Link>
          <Link to="/user/profile/change-password" state={{ userType, userID }} className="profile-btn">Alterar Senha</Link>
          <Link to="/user/favorites" state={{ userType, userID }} className="profile-btn">Meus Favoritos</Link>
        </div>
      </div>

      <div className="logout-container">
        <button onClick={handleLogout} className="logout-btn">Log Out</button>
      </div>
    </section>
  );
};

export default ProfilePage;