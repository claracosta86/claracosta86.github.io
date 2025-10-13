import React, { createContext, useState, useContext } from 'react';

const UserContext = createContext(null);

const getInitialUser = () => {
  try {
    const storedUser = localStorage.getItem('user');
    return storedUser ? JSON.parse(storedUser) : null;
  } catch (error) {
    console.error("Falha ao ler o usuário do localStorage.", error);
    return null;
  }
};

export const UserProvider = ({ children }) => {
  const [user, _setUser] = useState(getInitialUser);

  const setUser = (userData) => {
    _setUser(userData);
    
    if (userData) {
      localStorage.setItem('user', JSON.stringify(userData));
    } else {
      localStorage.removeItem('user');
    }
  };

  return (
    <UserContext.Provider value={{ user, setUser }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => useContext(UserContext);
export default UserProvider;