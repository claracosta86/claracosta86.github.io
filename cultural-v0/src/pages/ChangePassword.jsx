import { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { useUser } from '../contexts/UserContext';

const ChangePassword = () => {
  const { user } = useUser();
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    currentPassword: '',
    newPassword: '',
    confirmNewPassword: ''
  });
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  if (!user) {
    navigate('/login');
    return null;
  }

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    // Validações
    if (formData.newPassword !== formData.confirmNewPassword) {
      setError('As novas senhas não coincidem');
      return;
    }

    if (formData.newPassword.length < 6) {
      setError('A nova senha deve ter pelo menos 6 caracteres');
      return;
    }

    if (formData.currentPassword === formData.newPassword) {
      setError('A nova senha deve ser diferente da senha atual');
      return;
    }

    setIsLoading(true);

    try {
      // Simular chamada de API - em produção seria uma chamada real
      await new Promise(resolve => setTimeout(resolve, 1000));
      
      // Mock de validação de senha atual
      if (formData.currentPassword === '123456') {
        setSuccess('Senha alterada com sucesso!');
        setFormData({
          currentPassword: '',
          newPassword: '',
          confirmNewPassword: ''
        });
        
        // Redirecionar após 2 segundos
        setTimeout(() => {
          navigate('/profile');
        }, 2000);
      } else {
        setError('Senha atual incorreta');
      }
    } catch (err) {
      setError('Erro ao alterar senha. Tente novamente.');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="change-password-container">
      <div className="change-password-box">
        <div className="change-password-header">
          <Link to="/profile" className="back-btn">
            <img src="/images/goback.png" alt="Voltar" />
          </Link>
          <h2>Alterar Senha</h2>
        </div>
        
        {error && <div className="error-message">{error}</div>}
        {success && <div className="success-message">{success}</div>}
        
        <form onSubmit={handleSubmit} className="change-password-form">
          <div className="form-group">
            <label htmlFor="currentPassword">Senha Atual:</label>
            <input
              type="password"
              id="currentPassword"
              name="currentPassword"
              value={formData.currentPassword}
              onChange={handleChange}
              required
              placeholder="Digite sua senha atual"
            />
          </div>
          
          <div className="form-group">
            <label htmlFor="newPassword">Nova Senha:</label>
            <input
              type="password"
              id="newPassword"
              name="newPassword"
              value={formData.newPassword}
              onChange={handleChange}
              required
              placeholder="Digite sua nova senha"
            />
          </div>
          
          <div className="form-group">
            <label htmlFor="confirmNewPassword">Confirmar Nova Senha:</label>
            <input
              type="password"
              id="confirmNewPassword"
              name="confirmNewPassword"
              value={formData.confirmNewPassword}
              onChange={handleChange}
              required
              placeholder="Confirme sua nova senha"
            />
          </div>
          
          <button 
            type="submit" 
            className="change-password-btn" 
            disabled={isLoading}
          >
            {isLoading ? 'Alterando...' : 'Alterar Senha'}
          </button>
        </form>
        
        <div className="password-requirements">
          <h3>Requisitos da senha:</h3>
          <ul>
            <li>Mínimo de 6 caracteres</li>
            <li>Diferente da senha atual</li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default ChangePassword;

