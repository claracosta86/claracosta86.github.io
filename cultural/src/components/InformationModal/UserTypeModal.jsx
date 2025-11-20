const InfoModal = ({ isOpen, onClose }) => {
  if (!isOpen) return null;

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-container" onClick={(e) => e.stopPropagation()}>
        <h2 className="modal-title">Tipos de Usuário</h2>

        <div className="modal-content">
          <div>
            <h3 className="type-title">👤 Usuário Comum</h3>
            <p>
              Como usuário, você pode descobrir, salvar e participar dos melhores eventos culturais
              da sua cidade.
            </p>
          </div>
          <div>
            <h3 className="type-title">🎤 Organizador</h3>
            <p>
              Como organizador, você tem as ferramentas para criar, divulgar e gerenciar seus
              próprios eventos.
            </p>
          </div>
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

export default InfoModal;