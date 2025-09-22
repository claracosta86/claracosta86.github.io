import React from 'react';
import './NotificationModal.css'; // Crie ou mova o CSS correspondente

const NotificationModal = ({ isOpen, onClose, notifications, onNotificationClick }) => {
  if (!isOpen) return null;

  // A função de navegação agora é recebida via props
  const handleLinkClick = (culturalID) => {
    onNotificationClick(culturalID);
    onClose(); // Fecha o modal após o clique
  };

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={e => e.stopPropagation()}>
        <h2 className="modal-title">Notificações</h2>
        <div className="modal-content">
          {notifications.length === 0 ? (
            <p>Você não tem novas notificações.</p>
          ) : (
            notifications.map((notif) => (
              <div key={notif.ID} className="notification-item">
                <p>
                  Veja as atualizações de{' '}
                  <button onClick={() => handleLinkClick(notif.ID)}>
                    {notif.Title}
                  </button>
                </p>
              </div>
            ))
          )}
        </div>
        <div className="modal-actions">
          <button onClick={onClose} className="modal-close-btn">
            Entendi
          </button>
        </div>
      </div>
    </div>
  );
};

export default NotificationModal;