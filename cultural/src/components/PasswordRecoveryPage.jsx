import { Link } from 'react-router-dom';
import logo from '../assets/logo.png';
import './styles/login.css'; 

const PasswordRecoveryPage = () => {
  return (
    <section className="screen" id="tela-login">
      <a href="/">
        <img src={logo} alt="Logo Cultural" className="logo-img" />
      </a>
      <div className="login-box">
        <h2>Recuperação de Senha</h2>
        <label>Email</label>
        <input type="email" placeholder="email@exemplo.com" />

        <button className="btn">Enviar link de recuperação</button>

        <p className="small-letters">
          Lembrou sua senha? <Link to="/user/login" state={{ userType }} className="link">Login</Link>
        </p>
      </div>
    </section>
  );
};

export default PasswordRecoveryPage;