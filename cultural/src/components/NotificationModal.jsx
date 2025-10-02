import React from 'react';
import './NotificationModal.css'; // Crie ou mova o CSS correspondente

const NotificationModal = ({ isOpen, onClose, notifications }) => {
  if (!isOpen) return null;

  const handleLinkClick = (culturalID, culturalType) =>  async () => {
    let isEvent = culturalType === 'event' ? true : false
    navigate(`/card/${culturalID}`, { state: {userID, userType, "event": isEvent } });
  };

  const renderNotificationContent = (notif) => {
  switch (notif.notificationType) {
    case "updated":
      return <p>Veja as atualizações de <button onClick={() => handleLinkClick(notif.id, notif.type)}>{notif.title}</button></p>;
    case "canceled":
      return <p>O cultural <button onClick={() => handleLinkClick(notif.id, notif.type)}>{notif.title}</button> foi cancelado.</p>;
    case "closed":
      return <p>O cultural <button onClick={() => handleLinkClick(notif.id, notif.type)}>{notif.title}</button> foi encerrado.</p>;
    case "commented":
      return <p>Veja os novos comentários de <button onClick={() => handleLinkClick(notif.id, notif.type)}>{notif.title}</button>.</p>;
    default:
      return null;
  }
};

  
  return (
     <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={e => e.stopPropagation()}>
        <h2 className="modal-title">
          Notificações
        </h2>

        <div className="modal-content">
          {notifications.length === 0 ? (
            <p>Você não tem novas notificações.</p>
          ) : (
            notifications.map((notif, index) => (
              <div key={notif.ID} className="notification-item">
                {renderNotificationContent(notif)}
              </div>
            ))
          )}
        </div>

        <div className="modal-actions">
          <button
            onClick={onClose}
            className="modal-close-btn"
          >
            Fechar
          </button>
        </div>
      </div>
    </div>
  );
};


export default NotificationModal;