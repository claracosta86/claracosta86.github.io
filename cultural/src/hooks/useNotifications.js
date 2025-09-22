import { useState } from 'react';

// O hook recebe o userID para saber de quem buscar as notificações
export const useNotifications = (userID) => {
  const [notifications, setNotifications] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const fetchNotifications = async () => {
    if (!userID) return; // Não faz nada se não houver userID
    
    setIsLoading(true);
    try {
      const response = await fetch(`http://localhost:8080/notifications/${userID}`, {
        credentials: 'include',
      });
      if (response.ok) {
        const data = await response.json();
        setNotifications(data.culturals || []);
      }
    } catch (error) {
      console.error("Erro ao buscar por novas notificações:", error);
      setNotifications([]); // Limpa em caso de erro
    } finally {
      setIsLoading(false);
    }
  };

  const openNotificationModal = () => {
    // Busca as notificações sempre que o modal for aberto
    fetchNotifications(); 
    setIsModalOpen(true);
  };

  const closeNotificationModal = () => {
    setIsModalOpen(false);
  };

  // Retornamos o estado e as funções que os componentes precisarão
  return {
    isModalOpen,
    notifications,
    isLoading,
    openNotificationModal,
    closeNotificationModal,
  };
};