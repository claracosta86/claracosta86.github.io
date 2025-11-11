import './styles/App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LandingPage from './components/LandingPage';
import LoginPage from './components/LoginPage'; 
import RegisterPage from './components/RegisterPage';
import PasswordRecoveryPage from './components/PasswordRecoveryPage';
import HomePage from './components/HomePage';
import ProfilePage from './components/ProfilePage';
import EditProfilePage from './components/EditProfilePage';
import ChangePasswordPage from './components/ChangePasswordPage';
import FavoritesPage from './components/FavoritesPage';
import ManageCulturalPage from './components/ManageCulturalPage';
import CardPage from './components/CardPage';
import CreateCulturalPage from './components/CreateCulturalPage';
import UserProvider from './contexts/UserContext';
import ProtectedRoute from './components/ProtectedRoute';
import OrganizerPage from './components/OrganizerPage';
import EditCulturalPage from './components/EditCulturalPage';
import CommentaryPage from './components/CommentaryPage';
import SearchPage from './components/SearchPage';

function App() {
  return (
    <UserProvider>
      <Router>
        <Routes>
          <Route element={<ProtectedRoute />}>
            <Route path="/" element={<LandingPage />} />
            <Route path="/user/login" element={<LoginPage />} />
            <Route path="/user/register" element={<RegisterPage />} />
            <Route path="/user/password-recovery" element={<PasswordRecoveryPage />} />
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
            <Route path="/commentaries/:culturalType/:id" element={<CommentaryPage />} />
            <Route path="/search" element={<SearchPage />} />
          </Route>
        </Routes>
      </Router>
    </UserProvider>
  );
}

export default App;
