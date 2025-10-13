import { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { useUser } from '../contexts/UserContext';
import '../styles/profile.css';

const Profile = () => {
  const { user, updateUser } = useUser();
  const navigate = useNavigate();
  const [isEditing, setIsEditing] = useState(false);
  const [editData, setEditData] = useState({
    name: user?.name || '',
    email: user?.email || ''
  });
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleEdit = () => {
    setIsEditing(true);
  };

  const handleSave = async () => {
    try {
      setIsLoading(true);
      setError('');
      
      const result = await updateUser(editData);
      
      if (result.success) {
        setIsEditing(false);
      } else {
        setError(result.message || 'Erro ao atualizar perfil');
      }
    } catch (err) {
      setError('Erro ao atualizar perfil. Tente novamente.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleCancel = () => {
    setEditData({
      name: user?.name || '',
      email: user?.email || ''
    });
    setIsEditing(false);
    setError('');
  };

  const handleChange = (e) => {
    setEditData({
      ...editData,
      [e.target.name]: e.target.value
    });
  };

  if (!user) {
    return <div>Carregando...</div>;
  }

  return (
    <div className="screen">
      <div className="top-bar-common">
        <div className="logo-container">
          <img src="/images/logo6.png" alt="Logo" className="logo-img" />
        </div>
        <div className="right-section">
          <img src="/images/notifications-icon.png" alt="Notificações" className="icon" />
          <img src="/images/user-icon.png" alt="Usuário" className="icon" />
        </div>
      </div>

      <div className="profile-box">
        <div className="header-title">
          <img src="/images/goback-icon.png" alt="Voltar" className="goback-img" onClick={() => navigate('/home')} />
          <h2>Meu Perfil</h2>
        </div>

        {error && <div className="error">{error}</div>}

        {!isEditing ? (
          <div>
            <label>Nome:</label>
            <input type="text" value={user.name} disabled />
            
            <label>Email:</label>
            <input type="email" value={user.email} disabled />
            
            <label>Tipo de Usuário:</label>
            <input type="text" value={user.userType === 'common' ? 'Usuário Comum' : 'Organizador'} disabled />
            
            <div className="profile-actions">
              <button className="btn" onClick={handleEdit}>Editar Perfil</button>
              <Link to="/favorites" className="btn">Meus Favoritos</Link>
              <Link to="/change-password" className="btn">Alterar Senha</Link>
            </div>
            
            <div className="logout-container">
              <Link to="/" className="logout-btn">Sair</Link>
            </div>
          </div>
        ) : (
          <div>
            <label>Nome:</label>
            <input
              type="text"
              name="name"
              value={editData.name}
              onChange={handleChange}
            />
            
            <label>Email:</label>
            <input
              type="email"
              name="email"
              value={editData.email}
              onChange={handleChange}
            />
            
            <div className="profile-actions">
              <button className="btn" onClick={handleSave} disabled={isLoading}>
                {isLoading ? 'Salvando...' : 'Salvar'}
              </button>
              <button className="btn" onClick={handleCancel}>Cancelar</button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default Profile;

