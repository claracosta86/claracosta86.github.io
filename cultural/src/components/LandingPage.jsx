import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import indexBottom from '../assets/index-bottom.png';
import logo from '../assets/logo.png';
import infoIcon from '../assets/blueinfo-icon.png';
import hooverIcon from '../assets/redinfo-icon.png';
import './styles/landing.css';


const InfoModal = ({ isOpen, onClose }) => {
  if (!isOpen) return null;

  return (
     <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={e => e.stopPropagation()}>
        <h2 className="modal-title">
          Tipos de Usuário
        </h2>

        <div className="modal-content">
          <div>
            <h3 className="type-title">👤 Usuário Comum</h3>
            <p>
              Como usuário, você pode descobrir, salvar e participar dos melhores eventos culturais da sua cidade.
            </p>
          </div>
          <div>
            <h3 className="type-title">🎤 Organizador</h3>
            <p>
              Como organizador, você tem as ferramentas para criar, divulgar e gerenciar seus próprios eventos.
            </p>
          </div>
        </div>

        <div className="modal-actions">
          <button
            onClick={onClose}
            className="modal-close-btn"
          >
            Entendi
          </button>
        </div>
      </div>
    </div>
  );
};

const LandingPage = () => {
  const navigate = useNavigate();

  const [isInfoModalOpen, setInfoModalOpen] = useState(false);
  const [isIconHovered, setIconHovered] = useState(false);

  const handleUserTypeSelection = async (userType) => {
    try {
      const response = await fetch("http://localhost:8080/users/select-type", {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        },
        body: `userType=${userType}`,
        credentials: 'include'
      });
      if (response.ok) {
        navigate('/user/login');
      } else {
        console.error("Erro ao selecionar o tipo de usuário no backend.");
      }
    } catch (error) {
      console.error("Erro de rede ao comunicar com o backend:", error);
    }
  };

  return (
    <>
      <InfoModal isOpen={isInfoModalOpen} onClose={() => setInfoModalOpen(false)} />
      <main className="phone">
        <div className="screen">
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
            </form>
          </div>
          <img src={indexBottom} alt="Grupo de pessoas" className="footer-img" />
        </div>
      </main>
    </>
  );
};

export default LandingPage;