import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { useUser } from '../contexts/UserContext';
import UserTypeSelection from './UserTypeSelection';
import Home from '../pages/Home';
import Login from '../pages/Login';
import Register from '../pages/Register';
import Profile from '../pages/Profile';
import Favorites from '../pages/Favorites';
import ChangePassword from '../pages/ChangePassword';
import PasswordRecovery from '../pages/PasswordRecovery';

// Componente para rotas protegidas
const ProtectedRoute = ({ children }) => {
  const { isAuthenticated } = useUser();
  
  if (!isAuthenticated) {
    return <Navigate to="/login" replace />;
  }
  
  return children;
};

const AppRouter = () => {
  const { isAuthenticated } = useUser();

  return (
    <Router>
      <Routes>
        {/* Rota pública - seleção de tipo de usuário */}
        <Route path="/" element={<UserTypeSelection />} />
        
        {/* Rotas de autenticação */}
        <Route path="/login" element={
          isAuthenticated ? <Navigate to="/home" replace /> : <Login />
        } />
        <Route path="/register" element={
          isAuthenticated ? <Navigate to="/home" replace /> : <Register />
        } />
        <Route path="/password-recovery" element={<PasswordRecovery />} />
        
        {/* Rotas protegidas */}
        <Route path="/home" element={
          <ProtectedRoute>
            <Home />
          </ProtectedRoute>
        } />
        
        <Route path="/profile" element={
          <ProtectedRoute>
            <Profile />
          </ProtectedRoute>
        } />
        
        <Route path="/favorites" element={
          <ProtectedRoute>
            <Favorites />
          </ProtectedRoute>
        } />
        
        <Route path="/change-password" element={
          <ProtectedRoute>
            <ChangePassword />
          </ProtectedRoute>
        } />
        
        {/* Rota para páginas de eventos e atrações (futuras implementações) */}
        <Route path="/event/:id" element={
          <ProtectedRoute>
            <div>Página do Evento (em desenvolvimento)</div>
          </ProtectedRoute>
        } />
        
        <Route path="/attraction/:id" element={
          <ProtectedRoute>
            <div>Página da Atração (em desenvolvimento)</div>
          </ProtectedRoute>
        } />
        
        {/* Rota 404 */}
        <Route path="*" element={
          <div className="not-found">
            <h1>Página não encontrada</h1>
            <p>A página que você está procurando não existe.</p>
            <a href="/">Voltar ao início</a>
          </div>
        } />
      </Routes>
    </Router>
  );
};

export default AppRouter;

