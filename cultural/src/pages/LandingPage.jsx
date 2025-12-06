// src/components/LandingPage.jsx
import { useUser } from '../contexts/UserContext';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import InfoModal from '../components/InformationModal/UserTypeModal';
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
      <main className="landing-screen">
        <InfoModal 
          isOpen={isInfoModalOpen} 
          onClose={() => setInfoModalOpen(false)} 
        />
        <div className="landing-container">
          <img src={logo} alt="Logo Cultural" className="landing-logo" />
          <div className="landing-box">
            <div className="welcome-container">
              <h2>Bem-vind@!</h2>
              <button
                onClick={() => setInfoModalOpen(true)}
                id="userType-information-button"
                aria-label="Ver informações sobre os tipos de usuário"
                onMouseEnter={() => setIconHovered(true)}
                onMouseLeave={() => setIconHovered(false)}
              >
                <img
                  src={isIconHovered ? hooverIcon : infoIcon}
                  alt=""
                  className="userType-information"
                />
              </button>
            </div>
            <form>
              <div className="landing-actions">
                <button
                  type="button"
                  onClick={() => handleUserTypeSelection('common')}
                  className="landing-button"
                  aria-label="Entrar como Usuário Comum"
                >
                  Sou Usuário
                </button>
                <button
                  type="button"
                  onClick={() => handleUserTypeSelection('organizer')}
                  className="landing-button"
                  aria-label="Entrar como Organizador de Eventos"
                >
                  Sou Organizador
                </button>
              </div>
              <p className="actions-separator" aria-hidden="true"> ou </p>
              <button 
                type="button" 
                onClick={() => navigate('/user/register')} 
                className="landing-button"
                aria-label="Criar uma nova conta"
              >
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
