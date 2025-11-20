// src/components/LandingPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import InfoModal from './InformationModal/UserTypeModal';
import logo from '../assets/logo.png';
import infoIcon from '../assets/blueinfo-icon.png';
import hooverIcon from '../assets/redinfo-icon.png';
import './styles/landing.css';

const LandingPage = () => {
  const navigate = useNavigate();

  const { setUser } = useUser();
  
  const [isInfoModalOpen, setInfoModalOpen] = useState(false);
  const [isIconHovered, setIconHovered] = useState(false);

  const handleUserTypeSelection = async (userType) => {
    setUser({ type: userType });
    localStorage.setItem('user', JSON.stringify({ type: userType }));
    navigate('/user/login');
  };

  return (
    <>
      <InfoModal 
        isOpen={isInfoModalOpen} 
        onClose={() => setInfoModalOpen(false)} 
      />
      <main className="screen">
        <div className="landing-container">
          <img src={logo} alt="Logo Cultural" className="logo-img" />
          <div className="box">
            <div className="welcome-container">
              <h2>Bem-vind@!</h2>
              <button
                onClick={() => setInfoModalOpen(true)}
                className="info-btn"
                aria-label="Ver informação sobre tipos de usuário"
                onMouseEnter={() => setIconHovered(true)}
                onMouseLeave={() => setIconHovered(false)}
              >
                <img
                  src={isIconHovered ? hooverIcon : infoIcon}
                  alt="Informação tipos de usuário"
                  className="info-img"
                />
              </button>
            </div>
            <form>
              <div className="actions">
                <button
                  type="button"
                  onClick={() => handleUserTypeSelection('common')}
                  className="btn"
                >
                  Sou Usuário
                </button>
                <button
                  type="button"
                  onClick={() => handleUserTypeSelection('organizer')}
                  className="btn"
                >
                  Sou Organizador
                </button>
              </div>
              <p className="separator"> ou </p>
              <button type="button" onClick={() => navigate('/user/register')} className="btn">
                Criar Conta
              </button>
            </form>
          </div>
          <b>@Cultural :) 2025</b>
        </div>
      </main>
    </>
  );
};

export default LandingPage;
