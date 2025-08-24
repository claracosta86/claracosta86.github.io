import './styles/App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LandingPage from './components/LandingPage';
import LoginPage from './components/LoginPage'; 
import UserRegisterPage from './components/UserRegisterPage';
import PasswordRecoveryPage from './components/PasswordRecoveryPage';
import HomePage from './components/HomePage';
import ProfilePage from './components/ProfilePage';
import EditProfilePage from './components/EditProfilePage';
import ChangePasswordPage from './components/ChangePasswordPage';
import FavoritesPage from './components/FavoritesPage';
import ManageCulturalPage from './components/ManageCulturalPage';
import CardPage from './components/CardPage';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/user/login" element={<LoginPage />} />
        <Route path="/user/register" element={<UserRegisterPage />} />
        <Route path="/user/password-recovery" element={<PasswordRecoveryPage />} />
        <Route path="/home" element={<HomePage />} />
        <Route path="/user/profile" element={<ProfilePage />} />
        <Route path="/user/profile/edit" element={<EditProfilePage />} />
        <Route path="/user/profile/change-password" element={<ChangePasswordPage />} />
        <Route path="/user/favorites" element={<FavoritesPage />} />
        <Route path="/user/profile/manage-cultural" element={<ManageCulturalPage />} />
        <Route path="/card/:id" element={<CardPage />} />
      </Routes>
    </Router>
  );
}

export default App;
