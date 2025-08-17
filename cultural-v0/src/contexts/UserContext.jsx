import { createContext, useContext, useState, useEffect } from 'react';
import { authAPI, userAPI } from '../config/api';

const UserContext = createContext();

export const useUser = () => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error('useUser deve ser usado dentro de um UserProvider');
  }
  return context;
};

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [userType, setUserType] = useState(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Verificar se há dados do usuário no localStorage
    const savedUser = localStorage.getItem('user');
    const savedUserType = localStorage.getItem('userType');
    const authToken = localStorage.getItem('authToken');
    
    if (savedUser && savedUserType && authToken) {
      setUser(JSON.parse(savedUser));
      setUserType(savedUserType);
      setIsAuthenticated(true);
    }
    
    setIsLoading(false);
  }, []);

  const login = async (email, password) => {
    try {
      setIsLoading(true);
      
      // Chamada real para a API
      const response = await authAPI.login(email, password);
      
      if (response.success && response.data) {
        const userData = response.data.user;
        const token = response.data.token;
        const type = response.data.userType || 'common';
        
        // Salvar dados no estado
        setUser(userData);
        setUserType(type);
        setIsAuthenticated(true);
        
        // Salvar no localStorage
        localStorage.setItem('user', JSON.stringify(userData));
        localStorage.setItem('userType', type);
        localStorage.setItem('authToken', token);
        
        return { success: true };
      } else {
        throw new Error(response.message || 'Erro no login');
      }
    } catch (error) {
      console.error('Login failed:', error);
      return { 
        success: false, 
        message: error.message || 'Erro ao fazer login' 
      };
    } finally {
      setIsLoading(false);
    }
  };

  const register = async (userData) => {
    try {
      setIsLoading(true);
      
      // Chamada real para a API
      const response = await authAPI.register(userData);
      
      if (response.success && response.data) {
        const newUser = response.data.user;
        const token = response.data.token;
        const type = response.data.userType || 'common';
        
        // Salvar dados no estado
        setUser(newUser);
        setUserType(type);
        setIsAuthenticated(true);
        
        // Salvar no localStorage
        localStorage.setItem('user', JSON.stringify(newUser));
        localStorage.setItem('userType', type);
        localStorage.setItem('authToken', token);
        
        return { success: true };
      } else {
        throw new Error(response.message || 'Erro no registro');
      }
    } catch (error) {
      console.error('Registration failed:', error);
      return { 
        success: false, 
        message: error.message || 'Erro ao criar conta' 
      };
    } finally {
      setIsLoading(false);
    }
  };

  const logout = async () => {
    try {
      // Chamada real para a API (opcional)
      await authAPI.logout();
    } catch (error) {
      console.error('Logout API call failed:', error);
    } finally {
      // Limpar estado local independente da API
      setUser(null);
      setUserType(null);
      setIsAuthenticated(false);
      
      // Limpar localStorage
      localStorage.removeItem('user');
      localStorage.removeItem('userType');
      localStorage.removeItem('authToken');
    }
  };

  const updateUser = async (userData) => {
    try {
      setIsLoading(true);
      
      // Chamada real para a API
      const response = await userAPI.updateProfile(userData);
      
      if (response.success && response.data) {
        const updatedUser = response.data;
        
        // Atualizar estado
        setUser(updatedUser);
        
        // Atualizar localStorage
        localStorage.setItem('user', JSON.stringify(updatedUser));
        
        return { success: true };
      } else {
        throw new Error(response.message || 'Erro ao atualizar perfil');
      }
    } catch (error) {
      console.error('Profile update failed:', error);
      return { 
        success: false, 
        message: error.message || 'Erro ao atualizar perfil' 
      };
    } finally {
      setIsLoading(false);
    }
  };

  const value = {
    user,
    userType,
    isAuthenticated,
    isLoading,
    login,
    register,
    logout,
    updateUser,
  };

  return (
    <UserContext.Provider value={value}>
      {children}
    </UserContext.Provider>
  );
};

