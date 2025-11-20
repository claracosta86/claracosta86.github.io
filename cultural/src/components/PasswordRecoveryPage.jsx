// src/components/PasswordRecoveryPage.jsx
import { Link } from 'react-router-dom';
import logo from '../assets/logo.png';
import './styles/login.css';

const PasswordRecoveryPage = () => {
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
        <input type="email" placeholder="email@exemplo.com" />

        <button className="btn">Enviar link de recuperação</button> <br />

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
