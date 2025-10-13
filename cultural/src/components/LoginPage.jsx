// src/components/LoginPage.jsx
import { useUser } from '../contexts/UserContext';
import { useNavigate, Link } from 'react-router-dom';
import { useState, useEffect } from 'react';
import './styles/login.css';
import logo from '../assets/logo.png';
import visiblePassword from '../assets/visiblepassword-icon.png';
import invisiblePassword from '../assets/invisiblepassword-icon.png';

const LoginPage = () => {
  const navigate = useNavigate();

  const { setUser } = useUser();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [userType, setUserType] = useState('common');

  const [error, setError] = useState('');
  const [isTypeError, setIsTypeError] = useState(false);

  const [showPassword, setShowPassword] = useState(false);
  const togglePasswordVisibility = () => {
    setShowPassword(!showPassword);
  };

  useEffect(() => {
    const fetchUserType = async () => {
      try {
        const response = await fetch('http://localhost:8080/users/get-type', {
          credentials: 'include',
        });
        if (response.ok) {
          const data = await response.json();
          setUserType(data.userType);
        }
      } catch (error) {
        console.error('Erro ao buscar o tipo de usuário:', error);
      }
    };
    fetchUserType();
  }, []);

  const handleSubmit = async (event) => {
    event.preventDefault();

    // Reset error messages on new submission
    setError('');

    // Client-side validation
    if (!email && !password) {
      setError('Por favor, insira seu e-mail e senha.');
      return;
    }
    if (!email) {
      setError('Por favor, insira seu e-mail.');
      return;
    }
    if (!password) {
      setError('Por favor, insira sua senha.');
      return;
    }

    try {
      const response = await fetch('http://localhost:8080/users/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
        credentials: 'include',
      });

      if (response.status === 404 || response.status === 401) {
        setError('Seu usuário ou senha estão incorretos.');
        return;
      }

      if (!response.ok) {
        throw new Error('Erro ao fazer login');
      }

      if (response.ok) {
        const data = await response.json();

        if (userType !== data.type) {
          setError('Seu usuário não pertence a esta categoria! Volte à página inicial.');
          setIsTypeError(true);
          return;
        }

        const formData = new URLSearchParams();
        formData.append('userType', data.type);
        formData.append('userID', data.userID);

        try {
          const response = await fetch('http://localhost:8080/users/set-information', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: formData,
            credentials: 'include',
          });
          if (response.ok) {
            setUser(data);
            localStorage.setItem('user', JSON.stringify(data));
            navigate('/home');
          } else {
            console.error('Erro ao selecionar o tipo de usuário no backend.');
          }
        } catch (error) {
          console.error('Erro de rede ao comunicar com o backend:', error);
        }
      }
    } catch (err) {
      console.error('Erro:', err);
      setError('Ocorreu um erro. Tente novamente mais tarde.');
    }
  };

  return (
    <section className="screen" id="tela-login">
      <img src={logo} alt="Logo Cultural" className="logo-img" />
      <div className="login-box">
        <h2>Entre na sua conta</h2>
        <form id="loginForm" onSubmit={handleSubmit}>
          <label htmlFor="email">Email</label>
          <input
            id="email"
            type="email"
            placeholder="email@exemplo.com"
            autoComplete="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />

          <div className="password-container">
            <label htmlFor="password">Senha</label>
            <input
              id="password"
              type={showPassword ? 'text' : 'password'}
              placeholder="Digite sua senha"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <button type="button" className="password-toggle" onClick={togglePasswordVisibility}>
              <img
                className="toggle-btn"
                src={showPassword ? visiblePassword : invisiblePassword}
                alt="Toggle password visibility"
              />
            </button>
          </div>
          {error && <span className="error">{error}</span>}
          <div className="login-actions">
            <button type="submit" className="login-btn">
              Login
            </button>
            <div className="login-actions-row">
              <button onClick={() => navigate('/')} className="login-btn">
                Voltar
              </button>
              <button onClick={() => navigate('/user/password-recovery')} className="login-btn">
                Recuperar Senha
              </button>
            </div>
          </div>
        </form>
      </div>
    </section>
  );
};

export default LoginPage;
