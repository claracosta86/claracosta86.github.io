// src/components/RegisterPage.jsx
import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './styles/register.css';
import logo from '../assets/logo.png';
import logoutIcon from '../assets/logout-icon.png';

const RegisterPage = () => {
  const navigate = useNavigate();
  const [userType, setUserType] = useState('');

  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [passwordConfirm, setPasswordConfirm] = useState('');
  const [document, setDocument] = useState('');
  const [companyName, setCompanyName] = useState('');

  const [error, setError] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();

    setError('');

    if (
      !name ||
      !email ||
      !password ||
      !passwordConfirm ||
      !document ||
      (userType === 'organizer' && !companyName)
    ) {
      setError('Por favor, preencha todos os campos obrigatórios.');
      return;
    }

    if (userType === 'organizer' && document.length !== 14) {
      setError('Por favor, insira um CNPJ válido.');
      return;
    }

    if (userType === 'common' && document.length !== 11) {
      setError('Por favor, insira um CPF válido.');
      return;
    }

    if (password.length < 8) {
      setError('A senha deve ter pelo menos 8 caracteres.');
      return;
    }

    if (password !== passwordConfirm) {
      setError('A senha não coincide com a confirmação.');
      return;
    }

    try {
      const response = await fetch('http://localhost:8080/users/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          name,
          email,
          password,
          document,
          companyName: userType === 'organizer' ? companyName : '',
          type: userType,
        }),
      });

      if (response.status === 409) {
        setError('Este e-mail já está cadastrado.');
        return;
      }

      if (!response.ok) {
        throw new Error('Erro ao cadastrar usuário');
      }

      console.log('Cadastro realizado com sucesso!');
      window.location.href = '/user/login';
    } catch (err) {
      console.error('Erro:', err);
      setError('Ocorreu um erro ao cadastrar. Tente novamente mais tarde.');
    }
  };

  const handleTypeChange = (event) => {
    setUserType(event.target.value);
    setError('');
    setName('');
    setEmail('');
    setPassword('');
    setPasswordConfirm('');
    setDocument('');
    setCompanyName('');
  };

  return (
    <section className="screen" id="tela-login">
      <header className="top-bar">
        <img src={logo} alt="Logo Cultural" className="logo-tiny" />
        <div className="right-section">
          <Link to="/">
            <img src={logoutIcon} alt="Log-out" className="icon" />
          </Link>
        </div>
      </header>
      <div className="register-box">
        <h2>Crie sua conta</h2>
        <form id="registerUser" onSubmit={handleSubmit}>
          <fieldset className="selection-box">
            <legend>Selecione seu tipo de usuário:</legend>
            <div className="radio-option">
              <input
                type="radio"
                id="organizer"
                name="userType"
                value="organizer"
                checked={userType === 'organizer'}
                onChange={handleTypeChange}
              />
              <label htmlFor="organizer">Organizador</label>
            </div>
            <div className="radio-option">
              <input
                type="radio"
                id="common"
                name="userType"
                value="common"
                checked={userType === 'common'}
                onChange={handleTypeChange}
              />
              <label htmlFor="common">Comum</label>
            </div>
          </fieldset>

          {userType && (
            <>
              <label htmlFor="name" className="required">
                Nome
              </label>
              <input
                id="name"
                type="text"
                placeholder="Nome completo"
                autoComplete="given-name"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />

              <label htmlFor="email" className="required">
                Email
              </label>
              <input
                id="email"
                type="email"
                placeholder="email@exemplo.com"
                autoComplete="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />

              {userType === 'organizer' ? (
                <React.Fragment>
                  <label htmlFor="document" className="required">
                    CNPJ
                  </label>
                  <input
                    id="document"
                    type="text"
                    placeholder="CNPJ (apenas números)"
                    value={document}
                    onChange={(e) => setDocument(e.target.value)}
                  />
                  <label htmlFor="companyName" className="required">
                    Nome da Empresa
                  </label>
                  <input
                    id="companyName"
                    type="text"
                    placeholder="Nome da empresa"
                    value={companyName}
                    onChange={(e) => setCompanyName(e.target.value)}
                  />
                </React.Fragment>
              ) : (
                <React.Fragment>
                  <label htmlFor="document" className="required">
                    CPF
                  </label>
                  <input
                    id="document"
                    type="text"
                    placeholder="CPF (apenas números)"
                    value={document}
                    onChange={(e) => setDocument(e.target.value)}
                  />
                </React.Fragment>
              )}

              <label htmlFor="password" className="required">
                Senha
              </label>
              <input
                id="password"
                type="password"
                placeholder="Digite sua senha"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />

              <label htmlFor="passwordConfirm" className="required">
                Confirme a Senha
              </label>
              <input
                id="passwordConfirm"
                type="password"
                placeholder="Repita sua senha"
                value={passwordConfirm}
                onChange={(e) => setPasswordConfirm(e.target.value)}
              />

              {error && (
                <span id="error-message" className="error">
                  {error}
                </span>
              )}

              <input type="hidden" id="userType" value={userType} />

              <div className="register-actions">
                <button type="button" onClick={() => navigate('/')} className="btn">
                  Cancelar
                </button>
                <button type="submit" className="btn">
                  Cadastrar
                </button>
              </div>
            </>
          )}
        </form>
      </div>
    </section>
  );
};

export default RegisterPage;
