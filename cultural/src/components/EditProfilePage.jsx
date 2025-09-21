// src/components/EditProfilePage.jsx
import { useState, useEffect } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import './styles/profile.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';
import gobackIcon from '../assets/goback.png';

const EditProfilePage = () => {
    const navigate = useNavigate();

    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [companyName, setCompanyName] = useState('');
  
    const [error, setError] = useState('');
  
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

    useEffect(() => {
      const fetchUserData = async () => {
        try {
          const response = await fetch(`http://localhost:8080/users/${userID}/profile`, {
            credentials: 'include'
          });
          if (response.ok) {
            const data = await response.json();
            setEmail(data.email);
            setName(data.name);
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

  const handleSubmit = async (event) => {
    event.preventDefault();

    setError('');

    try {
        const response = await fetch(`http://localhost:8080/users/${userID}/profile/edit`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, email, companyName }),
        credentials: 'include'
      });

      if (response.status === 400 || response.status === 404) {
        setError("Desculpe, tivemos um problema ao buscar seus dados. Por favor tente mais tarde.");
        return;
      }

      if (!response.ok) {
        throw new Error("Erro ao editar dados do usuário.");
      }

      if (response.ok) {
        console.log("Change OK:", userID);
        navigate(`/user/profile`);
      }

    } catch (err) {
      console.error("Erro:", err);
      setError("Ocorreu um erro. Tente novamente mais tarde.");
    }
  };

  const handleDelete = async () => {
    try {
        const response = await fetch(`http://localhost:8080/users/${userID}/profile/delete`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        credentials: 'include'
      });

       if (response.status === 404) {
        setError("Desculpe, tivemos um problema ao buscar seus dados. Por favor tente mais tarde.");
        return;
      }

      if (!response.ok) {
        throw new Error("Erro ao editar dados do usuário.");
      }

      if (response.ok) {
        console.log("User deleted:", userID);
        navigate(`/`);
      } 

    } catch (err) {
      console.error("Erro:", err);
      setError("Ocorreu um erro. Tente novamente mais tarde.");
    }
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
          <Link to="/user/profile" state={{ userType, userID }}>
            <img src={gobackIcon} alt="Go Back Arrow" className="goback-img" />
          </Link>
          <h2>Editar Perfil</h2>
        </div>
        <form onSubmit={handleSubmit}>
          <label htmlFor="new-name">Nome do Usuário</label>
          <input id="new-name" type="text" placeholder={name}
              onChange={(e) => setName(e.target.value)} />

          <label htmlFor="new-email">E-mail</label>
          <input id="new-email" type="email" placeholder={email}
              onChange={(e) => setEmail(e.target.value)} />

          {userType === 'organizer' && (
            <>
              <label htmlFor="new-company-name">Nome da Empresa</label>
              <input id="new-company-name" type="text" placeholder={companyName}
                  onChange={(e) => setCompanyName(e.target.value)} />
            </>
          )}

          {error && <span className="error">{error}</span>}
          
          <div className="profile-actions">
            <button type="submit" className="profile-btn">Alterar Perfil</button>
          </div>
        </form>
      </div>
      <div className="delete-container">
        <button onClick={handleDelete} className="delete-btn">Deletar Usuário</button>
      </div>
    </section>
  );
};

export default EditProfilePage;