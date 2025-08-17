// Tipos de usuário
export const USER_TYPES = {
  COMMON: 'common',
  ORGANIZER: 'organizer'
};

// Tipos de conteúdo
export const CONTENT_TYPES = {
  EVENT: 'event',
  ATTRACTION: 'attraction'
};

// Categorias de eventos
export const EVENT_CATEGORIES = [
  'Música',
  'Teatro',
  'Dança',
  'Literatura',
  'Arte',
  'Cinema',
  'Esporte',
  'Educação',
  'Gastronomia',
  'Outros'
];

// Categorias de atrações
export const ATTRACTION_CATEGORIES = [
  'Histórico',
  'Cultural',
  'Religioso',
  'Natural',
  'Arquitetônico',
  'Gastronômico',
  'Comercial',
  'Recreativo',
  'Outros'
];

// Filtros de data
export const DATE_FILTERS = [
  { value: 'all', label: 'Todas as datas' },
  { value: 'today', label: 'Hoje' },
  { value: 'tomorrow', label: 'Amanhã' },
  { value: 'this-week', label: 'Esta semana' },
  { value: 'this-month', label: 'Este mês' }
];

// Estados de carregamento
export const LOADING_STATES = {
  IDLE: 'idle',
  LOADING: 'loading',
  SUCCESS: 'success',
  ERROR: 'error'
};

// Mensagens de erro padrão
export const ERROR_MESSAGES = {
  NETWORK_ERROR: 'Erro de conexão. Verifique sua internet.',
  UNAUTHORIZED: 'Você não está autorizado a acessar este recurso.',
  NOT_FOUND: 'Recurso não encontrado.',
  VALIDATION_ERROR: 'Dados inválidos. Verifique as informações.',
  SERVER_ERROR: 'Erro interno do servidor. Tente novamente mais tarde.',
  UNKNOWN_ERROR: 'Ocorreu um erro inesperado. Tente novamente.'
};

// Mensagens de sucesso padrão
export const SUCCESS_MESSAGES = {
  LOGIN_SUCCESS: 'Login realizado com sucesso!',
  REGISTER_SUCCESS: 'Conta criada com sucesso!',
  PROFILE_UPDATED: 'Perfil atualizado com sucesso!',
  PASSWORD_CHANGED: 'Senha alterada com sucesso!',
  FAVORITE_ADDED: 'Adicionado aos favoritos!',
  FAVORITE_REMOVED: 'Removido dos favoritos!'
};

// Configurações de paginação
export const PAGINATION_CONFIG = {
  DEFAULT_PAGE_SIZE: 10,
  MAX_PAGE_SIZE: 50,
  PAGE_SIZE_OPTIONS: [5, 10, 20, 50]
};

// Configurações de cache
export const CACHE_CONFIG = {
  USER_DATA_TTL: 24 * 60 * 60 * 1000, // 24 horas
  FAVORITES_TTL: 7 * 24 * 60 * 60 * 1000, // 7 dias
  SEARCH_RESULTS_TTL: 5 * 60 * 1000 // 5 minutos
};

