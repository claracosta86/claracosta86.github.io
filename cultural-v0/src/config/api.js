// Configuração da API
const API_BASE_URL = 'http://localhost:8080'; // Ajuste para a porta do seu backend Go

// Endpoints da API
export const API_ENDPOINTS = {
  // Autenticação
  LOGIN: `${API_BASE_URL}/api/users/login`,
  REGISTER: `${API_BASE_URL}/api/users/register`,
  LOGOUT: `${API_BASE_URL}/api/users/logout`,
  
  // Usuários
  USER_PROFILE: `${API_BASE_URL}/api/users/profile`,
  UPDATE_PROFILE: `${API_BASE_URL}/api/users/profile`,
  CHANGE_PASSWORD: `${API_BASE_URL}/api/users/change-password`,
  
  // Eventos
  EVENTS: `${API_BASE_URL}/api/events`,
  EVENT_DETAILS: (id) => `${API_BASE_URL}/api/events/${id}`,
  
  // Atrações
  ATTRACTIONS: `${API_BASE_URL}/api/attractions`,
  ATTRACTION_DETAILS: (id) => `${API_BASE_URL}/api/attractions/${id}`,
  
  // Favoritos
  FAVORITES: `${API_BASE_URL}/api/users/favorites`,
  ADD_FAVORITE: `${API_BASE_URL}/api/users/favorites`,
  REMOVE_FAVORITE: (id) => `${API_BASE_URL}/api/users/favorites/${id}`,
};

// Função para fazer requisições HTTP
export const apiRequest = async (url, options = {}) => {
  const defaultOptions = {
    headers: {
      'Content-Type': 'application/json',
    },
    ...options,
  };

  // Adicionar token de autenticação se existir
  const token = localStorage.getItem('authToken');
  if (token) {
    defaultOptions.headers.Authorization = `Bearer ${token}`;
  }

  try {
    const response = await fetch(url, defaultOptions);
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
    
    return await response.json();
  } catch (error) {
    console.error('API request failed:', error);
    throw error;
  }
};

// Funções específicas da API
export const authAPI = {
  login: async (email, password) => {
    return apiRequest(API_ENDPOINTS.LOGIN, {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    });
  },

  register: async (userData) => {
    return apiRequest(API_ENDPOINTS.REGISTER, {
      method: 'POST',
      body: JSON.stringify(userData),
    });
  },

  logout: async () => {
    return apiRequest(API_ENDPOINTS.LOGOUT, {
      method: 'POST',
    });
  },
};

export const userAPI = {
  getProfile: async () => {
    return apiRequest(API_ENDPOINTS.USER_PROFILE);
  },

  updateProfile: async (userData) => {
    return apiRequest(API_ENDPOINTS.UPDATE_PROFILE, {
      method: 'PUT',
      body: JSON.stringify(userData),
    });
  },

  changePassword: async (passwordData) => {
    return apiRequest(API_ENDPOINTS.CHANGE_PASSWORD, {
      method: 'POST',
      body: JSON.stringify(passwordData),
    });
  },
};

export const eventsAPI = {
  getAll: async () => {
    return apiRequest(API_ENDPOINTS.EVENTS);
  },

  getById: async (id) => {
    return apiRequest(API_ENDPOINTS.EVENT_DETAILS(id));
  },
};

export const attractionsAPI = {
  getAll: async () => {
    return apiRequest(API_ENDPOINTS.ATTRACTIONS);
  },

  getById: async (id) => {
    return apiRequest(API_ENDPOINTS.ATTRACTION_DETAILS(id));
  },
};

export const favoritesAPI = {
  getAll: async () => {
    return apiRequest(API_ENDPOINTS.FAVORITES);
  },

  add: async (itemData) => {
    return apiRequest(API_ENDPOINTS.ADD_FAVORITE, {
      method: 'POST',
      body: JSON.stringify(itemData),
    });
  },

  remove: async (id) => {
    return apiRequest(API_ENDPOINTS.REMOVE_FAVORITE(id), {
      method: 'DELETE',
    });
  },
};
