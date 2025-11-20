import React from 'react';  

const ConfirmModal = ({ isOpen, onClose, onConfirm }) => {
  if (!isOpen) return null;

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={(e) => e.stopPropagation()}>
        <h2 className="modal-title">Confirmar Logout</h2>
        <div className="modal-content">
          <p>
            Você realmente deseja sair do app?
          </p>
        </div>
        <div className="modal-actions">
          <button onClick={onClose} className="modal-close-btn">
            Cancelar
          </button>
          <button onClick={onConfirm} className="modal-confirm-btn">
            Sair
          </button>
        </div>
      </div>
    </div>
  );
};

export default ConfirmModal;