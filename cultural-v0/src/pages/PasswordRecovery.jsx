import { useState } from 'react';
import { Link } from 'react-router-dom';

const PasswordRecovery = () => {
  const [email, setEmail] = useState('');
  const [isSubmitted, setIsSubmitted] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setIsLoading(true);

    try {
      // Simular chamada de API - em produção seria uma chamada real
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      // Mock de validação
      if (email && email.includes('@')) {
        setIsSubmitted(true);
      } else {
        setError('Por favor, insira um email válido');
      }
    } catch (err) {
      setError('Erro ao processar solicitação. Tente novamente.');
    } finally {
      setIsLoading(false);
    }
  };

  if (isSubmitted) {
    return (
      <div className="password-recovery-container">
        <div className="password-recovery-box">
          <div className="logo-container">
            <img src="/logo6.png" alt="Logo Cultural" className="logo-img" />
          </div>
          
          <div className="success-message">
            <h2>Email Enviado!</h2>
            <p>
              Se o email <strong>{email}</strong> estiver cadastrado em nossa base, 
              você receberá um link para redefinir sua senha.
            </p>
            <p>
              Verifique sua caixa de entrada e spam. O link expira em 1 hora.
            </p>
          </div>
          
          <div className="actions">
            <Link to="/login" className="back-to-login-btn">
              Voltar ao Login
            </Link>
            <Link to="/" className="back-home-btn">
              Voltar ao Início
            </Link>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="password-recovery-container">
      <div className="password-recovery-box">
        <div className="logo-container">
          <img src="/logo6.png" alt="Logo Cultural" className="logo-img" />
        </div>
        
        <h2>Recuperar Senha</h2>
        <p className="description">
          Digite seu email cadastrado e enviaremos um link para você redefinir sua senha.
        </p>
        
        {error && <div className="error-message">{error}</div>}
        
        <form onSubmit={handleSubmit} className="password-recovery-form">
          <div className="form-group">
            <label htmlFor="email">Email:</label>
            <input
              type="email"
              id="email"
              name="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              placeholder="Digite seu email"
            />
          </div>
          
          <button 
            type="submit" 
            className="recovery-btn" 
            disabled={isLoading}
          >
            {isLoading ? 'Enviando...' : 'Enviar Link de Recuperação'}
          </button>
        </form>
        
        <div className="links">
          <Link to="/login" className="link">
            Lembrou sua senha? Faça login
          </Link>
          <Link to="/register" className="link">
            Não tem uma conta? Cadastre-se
          </Link>
        </div>
        
        <Link to="/" className="back-link">
          ← Voltar
        </Link>
      </div>
    </div>
  );
};

export default PasswordRecovery;

