import { useNavigate } from 'react-router-dom';

const UserTypeSelection = () => {
  const navigate = useNavigate();

  const handleUserTypeSelect = (userType) => {
    // Em uma implementação real, você poderia salvar a preferência
    // e redirecionar para login/registro
    if (userType === 'common') {
      navigate('/login');
    } else {
      navigate('/register');
    }
  };

  return (
    <main className="phone">
      <div className="screen">
        <img src="/images/logo6.png" className="logo-img" alt="Cultural logo" />
        <div className="box">
          <h2>Bem-vindo(a)!</h2>
          <div className="user-type-buttons">
            <button 
              onClick={() => handleUserTypeSelect('common')} 
              className="btn"
            >
              Sou Usuário
            </button>
            <button 
              onClick={() => handleUserTypeSelect('organizer')} 
              className="btn"
            >
              Sou Organizador
            </button>
          </div>
        </div>
        <img src="/images/index-bottom.png" className="footer-img" alt="Grupo de pessoas" />
      </div>
    </main>
  );
};

export default UserTypeSelection;

