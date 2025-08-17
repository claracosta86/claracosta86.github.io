import { useState, useEffect } from 'react';

export const useFavorites = () => {
  const [favorites, setFavorites] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  // Carregar favoritos do localStorage
  useEffect(() => {
    const savedFavorites = localStorage.getItem('favorites');
    if (savedFavorites) {
      try {
        setFavorites(JSON.parse(savedFavorites));
      } catch (error) {
        console.error('Erro ao carregar favoritos:', error);
        setFavorites([]);
      }
    }
  }, []);

  // Salvar favoritos no localStorage sempre que mudar
  useEffect(() => {
    localStorage.setItem('favorites', JSON.stringify(favorites));
  }, [favorites]);

  const addFavorite = (item) => {
    setFavorites(prev => {
      // Verificar se já existe
      const exists = prev.find(fav => fav.id === item.id && fav.type === item.type);
      if (exists) {
        return prev; // Já existe, não adicionar novamente
      }
      return [...prev, item];
    });
  };

  const removeFavorite = (id, type) => {
    setFavorites(prev => prev.filter(fav => !(fav.id === id && fav.type === type)));
  };

  const toggleFavorite = (item) => {
    const exists = favorites.find(fav => fav.id === item.id && fav.type === item.type);
    if (exists) {
      removeFavorite(item.id, item.type);
    } else {
      addFavorite(item);
    }
  };

  const isFavorite = (id, type) => {
    return favorites.some(fav => fav.id === id && fav.type === type);
  };

  const clearFavorites = () => {
    setFavorites([]);
  };

  return {
    favorites,
    isLoading,
    addFavorite,
    removeFavorite,
    toggleFavorite,
    isFavorite,
    clearFavorites
  };
};

