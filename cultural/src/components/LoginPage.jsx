// src/components/LoginPage.jsx
import { useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import './styles/login.css';
import logo from '../assets/logo.png';
import { Link } from 'react-router-dom';

const LoginPage = () => {
  const navigate = useNavigate();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [userType, setUserType] = useState('common');

  const [error, setError] = useState('');
  const [isTypeError, setIsTypeError] = useState(false);

  useEffect(() => {
    const fetchUserType = async () => {
      try {
        const response = await fetch("http://localhost:8080/users/get-type", {
             credentials: 'include'
        }); 
        if (response.ok) {
          const data = await response.json();
          setUserType(data.userType);
        }
      } catch (error) {
        console.error("Erro ao buscar o tipo de usuário:", error);
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
      const response = await fetch("http://localhost:8080/users/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
        credentials: 'include'
      });

      if (response.status === 404 || response.status === 401) {
        setError("Seu usuário ou senha estão incorretos.");
        return;
      }

      if (!response.ok) {
        throw new Error("Erro ao fazer login");
      }

      if (response.ok) {
        const data = await response.json();

        if (userType !== data.type) {
            setError("Seu usuário não pertence a esta categoria! Volte à página inicial.");
            setIsTypeError(true);
            return;
        }

        const formData = new URLSearchParams();
        formData.append('userType', data.type);
        formData.append('userID', data.userID);

        try {
            const response = await fetch("http://localhost:8080/users/set-information", {
                method: "POST",
                headers: {
                "Content-Type": "application/x-www-form-urlencoded"
                },
                body: formData,
                credentials: 'include'
            });
            if (response.ok) {
                navigate('/home');
            } else {
                console.error("Erro ao selecionar o tipo de usuário no backend.");
            }
            } catch (error) {
            console.error("Erro de rede ao comunicar com o backend:", error);
        }
      }

    } catch (err) {
      console.error("Erro:", err);
      setError("Ocorreu um erro. Tente novamente mais tarde.");
    }
  };

  return (
    <section className="screen" id="tela-login">
      <a href="/">
        <img src={logo} alt="Logo Cultural" className="logo-img" />
      </a>
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

          <label htmlFor="password">
            Senha <a href="/user/password-recovery" type="button" className="link-recovery">Esqueceu?</a>
          </label>
          <input
            id="password"
            type="password"
            placeholder="Digite sua senha"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />

          {error && <span className="error">{error}</span>}

          <p className="button-container">
            {isTypeError ? ( <Link to="/" type="button" className="btn"> Voltar </Link>) : 
                ( <button type="submit" className="btn"> Login </button> )
            }
          </p>
        </form>
        <p className="small-letters">
          Não tem uma conta? <Link to="/user/register" state={{ userType }} className="link">Cadastrar</Link>
        </p>
      </div>
    </section>
  );
};

export default LoginPage;