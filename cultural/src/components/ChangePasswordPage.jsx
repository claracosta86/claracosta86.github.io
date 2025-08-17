// src/components/ChangePasswordPage.jsx
import { useState } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import './styles/profile.css';
import logo from '../assets/logo.png';
import notificationsIcon from '../assets/notifications-icon.png';
import userIcon from '../assets/user-icon.png';
import gobackIcon from '../assets/goback.png';

const ChangePasswordPage = () => {
  const navigate = useNavigate();

  const location = useLocation();
  const userType = location.state?.userType || 'common';
  const userID = location.state?.userID || '';

  const [newPassword, setNewPassword] = useState('');
  const [newPasswordConfirmation, setNewPasswordConfirmation] = useState('');
  const [currentPassword, setCurrentPassword] = useState('');

  const [error, setError] = useState('');
  
  const handleSubmit = async (event) => {
    event.preventDefault();

    setError('');

    if (newPassword.length < 8) {
      setError("A senha deve ter pelo menos 8 caracteres.");
      return;
    }

    if (newPassword !== newPasswordConfirmation) {
      setError('A nova senha não coincide com a confirmação.');
      return;
    }

    try {
        const response = await fetch(`http://localhost:8080/users/${userID}/profile/change-password`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ currentPassword, newPassword }),
        credentials: 'include'
      });

      if (response.status === 401) {
        setError("A senha atual está incorreta.");
        return;
      }

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
          <Link to="/user/profile">
            <img src={gobackIcon} alt="Go Back Arrow" className="goback-img" />
          </Link>
          <h2>Alterar Senha</h2>
        </div>
        <form onSubmit={handleSubmit}>
          <label htmlFor="new-password">Nova Senha</label>
          <input id="new-password" type="password" placeholder="Nova senha" value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)} />

          <label htmlFor="new-password-confirmation">Confirme a Nova Senha</label>
          <input id="new-password-confirmation" type="password" placeholder="Repita a nova senha" value={newPasswordConfirmation}
            onChange={(e) => setNewPasswordConfirmation(e.target.value)} />

          <label htmlFor="current-password">Senha Atual</label>
          <input id="current-password" type="password" placeholder="Digite sua senha atual" value={currentPassword}
            onChange={(e) => setCurrentPassword(e.target.value)} />

          {error && <span className="error">{error}</span>}

          <div className="profile-actions">
            <button type="submit" className="profile-btn">Alterar Senha</button>
          </div>
        </form>
      </div>
    </section>
  );
};

export default ChangePasswordPage;