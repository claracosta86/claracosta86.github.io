import { useState } from 'react';

export const useNotifications = (userID) => {
  const [notifications, setNotifications] = useState([]);
  const [isNotificationModalOpen, setNotificationModalOpen] = useState(false);

  const fetchNotifications = async () => {
    try {
      const response = await fetch(`http://localhost:8080/notifications/${userID}`, {
        credentials: 'include',
      });
      if (response.ok) {
        const data = await response.json();
        setNotifications(data.culturals);
        console.log('Notificações recebidas:', data.culturals);
        setNotificationModalOpen(true);
      }
    } catch (error) {
      console.error('Erro ao buscar por novas notificações:', error);
    }
  };

  const markNotificationsAsSeen = async () => {
    try {
      const response = await fetch(`http://localhost:8080/notifications/${userID}/seen`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ notificationIDs: notifications.map((notif) => notif.ID) }),
        credentials: 'include',
      });
      if (response.ok) {
        console.log('Notificações marcadas como vistas com sucesso.');
      }
    } catch (error) {
      console.error('Erro ao atualizar notificações:', error);
    }
    setNotificationModalOpen(false);
  };

  return {
    notifications,
    isNotificationModalOpen,
    setNotificationModalOpen,
    fetchNotifications,
    markNotificationsAsSeen,
  };
};
