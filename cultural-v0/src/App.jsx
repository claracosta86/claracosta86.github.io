import { UserProvider } from './contexts/UserContext';
import AppRouter from './components/AppRouter';
import './App.css';

function App() {
  return (
    <UserProvider>
      <AppRouter />
    </UserProvider>
  );
}

export default App;
