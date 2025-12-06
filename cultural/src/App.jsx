import './styles/App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import LoginPage from './pages/LoginPage'; 
import RegisterPage from './pages/RegisterPage';
import PasswordRecoveryPage from './pages/PasswordRecoveryPage';
import HomePage from './pages/HomePage';
import ProfilePage from './pages/ProfilePage';
import EditProfilePage from './pages/EditProfilePage';
import ChangePasswordPage from './pages/ChangePasswordPage';
import FavoritesPage from './pages/FavoritesPage';
import ManageCulturalPage from './pages/ManageCulturalPage';
import CardPage from './pages/CardPage';
import CreateCulturalPage from './pages/CreateCulturalPage';
import UserProvider from './contexts/UserContext';
import ProtectedRoute from './components/ProtectedRoute';
import OrganizerPage from './pages/OrganizerPage';
import EditCulturalPage from './pages/EditCulturalPage';
import CommentPage from './pages/CommentPage';
import SearchPage from './pages/SearchPage';

function App() {
  return (
    <UserProvider>
      <Router>
        <Routes>
          <Route path="/" element={<LandingPage />} />
          <Route path="/user/login" element={<LoginPage />} />
          <Route path="/user/register" element={<RegisterPage />} />
          <Route path="/user/password-recovery" element={<PasswordRecoveryPage />} />
          
          <Route element={<ProtectedRoute />}>
            <Route path="/home" element={<HomePage />} />
            <Route path="/user/profile" element={<ProfilePage />} />
            <Route path="/user/profile/edit" element={<EditProfilePage />} />
            <Route path="/user/profile/change-password" element={<ChangePasswordPage />} />
            <Route path="/user/favorites" element={<FavoritesPage />} />
            <Route path="/user/profile/manage-cultural" element={<ManageCulturalPage />} />
            <Route path="/card/:culturalType/:id" element={<CardPage />} />
            <Route path="/create-cultural" element={<CreateCulturalPage />} />
            <Route path="/organizer/:id" element={<OrganizerPage />} />
            <Route path="/edit-cultural/:culturalType/:id" element={<EditCulturalPage />} />
            <Route path="/comments/:culturalType/:id" element={<CommentPage />} />
            <Route path="/search" element={<SearchPage />} />
          </Route>
        </Routes>
      </Router>
    </UserProvider>
  );
}

export default App;
