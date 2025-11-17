import React from 'react';  

const RemoveModal = ({ isOpen, onClose, onConfirm }) => {
  if (!isOpen) return null;

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={(e) => e.stopPropagation()}>
        <h2 className="modal-title">Confirmar Remoção</h2>
        <div className="modal-content">
          <p>
            Você realmente deseja excluir o evento/ponto turístico da plataforma?
          </p>
        </div>
        <div className="modal-actions">
          <button onClick={onClose} className="modal-close-btn">
            Cancelar
          </button>
          <button onClick={onConfirm} className="modal-confirm-btn">
            Remover
          </button>
        </div>
      </div>
    </div>
  );
};

export default RemoveModal;