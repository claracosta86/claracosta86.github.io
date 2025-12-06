// src/components/PasswordRecoveryPage.jsx
import { Link, useNavigate } from 'react-router-dom';
import { useState } from 'react';
import logo from '../assets/logo.png';
import './styles/login.css';

const PasswordRecoveryPage = () => {
  const navigate = useNavigate();

  const [error, setError] = useState('');
  const handleSendLinkClick = () => {
    const email = document.querySelector('input[id="email"]').value;
    if (email !== '') {
      navigate('/user/login');
    }
    else {
      setError('Por favor, insira seu e-mail.');
    }
  };

  return (
    <section className="screen" id="recovery-page">
      <header className="top-bar">
        <img src={logo} alt="Logo Cultural" className="logo-tiny" />
        <div className="right-section">
        </div>
      </header>
      <div className="login-box">
        <h2>Recuperação de Senha</h2>
        <label>Email</label>
        <input id="email" type="email" placeholder="email@exemplo.com" />
        {error && <span className="error">{error}</span>} 

        <br/><button className="login-btn" onClick={handleSendLinkClick}>Enviar link de recuperação</button> <br />

        <p className="small-letters">
          Lembrou sua senha?{' '}
          <Link to="/user/login" className="link">
            Volte ao login
          </Link>
        </p>
      </div>
    </section>
  );
};

export default PasswordRecoveryPage;
