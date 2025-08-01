// Função para mostrar modal de erro
function showErrorModal(message) {
  // Remove modal existente se houver
  const existingModal = document.getElementById('error-modal');
  if (existingModal) {
    existingModal.remove();
  }

  // Cria o modal
  const modal = document.createElement('div');
  modal.id = 'error-modal';
  modal.className = 'error-modal';
  
  modal.innerHTML = `
    <div class="error-modal-content">
      <h3>Erro de Validação</h3>
      <p>${message}</p>
      <div class="error-modal-buttons">
        <button class="error-btn-ok" onclick="closeErrorModal()">OK</button>
        <button class="error-btn-cancel" onclick="closeErrorModal()">Cancelar</button>
      </div>
    </div>
  `;
  
  document.body.appendChild(modal);
  
  // Adiciona evento para fechar ao clicar fora do modal
  modal.addEventListener('click', function(e) {
    if (e.target === modal) {
      closeErrorModal();
    }
  });
}

// Função para fechar o modal
function closeErrorModal() {
  const modal = document.getElementById('error-modal');
  if (modal) {
    modal.remove();
  }
}

// Função de validação do formulário
function validateForm() {
  const emailInput = document.querySelector('input[type="email"]');
  const passwordInput = document.querySelector('input[type="password"]');
  
  const email = emailInput.value.trim();
  const password = passwordInput.value.trim();
  
  // Validação do email
  if (!email) {
    showErrorModal('O campo de email não pode estar vazio. Por favor, digite seu email.');
    emailInput.focus();
    return false;
  }
  
  // Validação básica de formato de email
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!emailRegex.test(email)) {
    showErrorModal('Por favor, digite um email válido (exemplo: usuario@dominio.com).');
    emailInput.focus();
    return false;
  }
  
  // Validação da senha
  if (!password) {
    showErrorModal('O campo de senha não pode estar vazio. Por favor, digite sua senha.');
    passwordInput.focus();
    return false;
  }
  
  if (password.length < 6) {
    showErrorModal('A senha deve ter pelo menos 6 caracteres.');
    passwordInput.focus();
    return false;
  }
  
  return true;
}

// Adiciona event listeners quando a página carregar
document.addEventListener('DOMContentLoaded', function() {
  const loginButton = document.querySelector('.btn');
  
  if (loginButton) {
    loginButton.addEventListener('click', function(e) {
      e.preventDefault();
      
      if (validateForm()) {
        // Se a validação passou, pode redirecionar
        window.location.href = '../home/home.html';
      }
    });
  }
  
  // Adiciona validação em tempo real nos campos
  const emailInput = document.querySelector('input[type="email"]');
  const passwordInput = document.querySelector('input[type="password"]');
  
  if (emailInput) {
    emailInput.addEventListener('blur', function() {
      if (this.value.trim() && !(/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(this.value))) {
        this.style.borderColor = '#e74c3c';
      } else {
        this.style.borderColor = '#ccc';
      }
    });
  }
  
  if (passwordInput) {
    passwordInput.addEventListener('blur', function() {
      if (this.value.trim() && this.value.length < 6) {
        this.style.borderColor = '#e74c3c';
      } else {
        this.style.borderColor = '#ccc';
      }
    });
  }
});