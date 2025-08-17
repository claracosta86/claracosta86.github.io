import React, { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';
import './styles/register.css';
import logo from '../assets/logo.png';

const UserRegisterPage = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [passwordConfirm, setPasswordConfirm] = useState('');
  const [document, setDocument] = useState('');
  const [companyName, setCompanyName] = useState('');

  const location = useLocation();
  console.log("Location state:", location.state);
  const userType = location.state?.userType || 'common';

  const [error, setError] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();

    setError('');

    if (!name || !email || !password || !passwordConfirm || !document || (userType === 'organizer' && !companyName)) {
      setError("Por favor, preencha todos os campos obrigatórios.");
      return;
    }

    if (userType === 'organizer' && document.length !== 14) {
      setError("Por favor, insira um CNPJ válido.");
      return;
    }

    if (userType === 'common' && document.length !== 11) {
      setError("Por favor, insira um CPF válido.");
      return;
    }

    if (password.length < 8) {
      setError("A senha deve ter pelo menos 8 caracteres.");
      return;
    }

    if (password !== passwordConfirm) {
      setError("A senha não coincide com a confirmação.");
      return;
    }
    
    try {
      const response = await fetch("http://localhost:8080/users/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          name, 
          email, 
          password, 
          document,
          companyName: userType === 'organizer' ? companyName : '',
          type: userType
        }),
      });

      if (response.status === 409) {
        setError("Este e-mail já está cadastrado.");
        return;
      }
      
      if (!response.ok) {
        throw new Error("Erro ao cadastrar usuário");
      }

      console.log("Cadastro realizado com sucesso!");
      window.location.href = "/user/login";

    } catch (err) {
      console.error("Erro:", err);
      setError("Ocorreu um erro ao cadastrar. Tente novamente mais tarde.");
    }
  };

  return (
    <section className="screen" id="tela-login">
      <a href="/">
        <img src={logo} alt="Logo Cultural" className="logo-img" />
      </a>
      <div className="login-box">
        <h2>Crie sua conta</h2>
        <form id="registerUser" onSubmit={handleSubmit}>
          <label htmlFor="name" className="required">Nome</label>
          <input id="name" type="text" placeholder="Nome completo" autoComplete="given-name" value={name} onChange={(e) => setName(e.target.value)} />

          <label htmlFor="email" className="required">Email</label>
          <input id="email" type="email" placeholder="email@exemplo.com" autoComplete="email" value={email} onChange={(e) => setEmail(e.target.value)} />

          {userType === 'organizer' ? (
            <React.Fragment>
              <label htmlFor="document" className="required">CNPJ</label>
              <input id="document" type="text" placeholder="CNPJ (apenas números)" value={document} onChange={(e) => setDocument(e.target.value)} />
              <label htmlFor="companyName" className="required">Nome da Empresa</label>
              <input id="companyName" type="text" placeholder="Nome da empresa" value={companyName} onChange={(e) => setCompanyName(e.target.value)} />
            </React.Fragment>
          ) : (
            <React.Fragment>
              <label htmlFor="document" className="required">CPF</label>
              <input id="document" type="text" placeholder="CPF (apenas números)" value={document} onChange={(e) => setDocument(e.target.value)} />
            </React.Fragment>
          )}

          <label htmlFor="password" className="required">Senha</label>
          <input id="password" type="password" placeholder="Digite sua senha" value={password} onChange={(e) => setPassword(e.target.value)} />

          <label htmlFor="passwordConfirm" className="required">Confirme a Senha</label>
          <input id="passwordConfirm" type="password" placeholder="Repita sua senha" value={passwordConfirm} onChange={(e) => setPasswordConfirm(e.target.value)} />

          {error && <span id="error-message" className="error">{error}</span>}
          
          <input type="hidden" id="userType" value={userType} />

          <p className="button-container">
            <button type="submit" className="btn">Cadastrar</button>
          </p>
        </form>

        <p className="medium-letters">Já tem uma conta? <Link to="/user/login" className="link">Login</Link></p>
      </div>
    </section>
  );
};

export default UserRegisterPage;