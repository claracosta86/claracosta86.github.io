
document.addEventListener('DOMContentLoaded', function () {
  const registerForm = document.getElementById('registerUser');
  registerForm.addEventListener('submit', function(event) {
    event.preventDefault();

    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const passwordConfirm = document.getElementById('passwordConfirm').value;
    const documentNumber = document.getElementById('document').value;
    const type = document.getElementById('userType').value;
    const companyName = type === "organizer" ? document.getElementById('companyName').value : "";

    if (!name || !email || !password || !passwordConfirm || !documentNumber) {
      error = document.getElementById('error-empty-field')
      error.style.display = 'block';
      error.textContent = "Por favor, preencha todos os campos obrigatórios.";
      error.classList.add('error');
      return;
    }

    if (password.length < 8) {
      document.getElementById('error-empty-field').style.display = 'none';
      errorPassword = document.getElementById('error-password-mismatch')
      errorPassword.style.display = 'block';
      errorPassword.textContent = "A senha deve ter pelo menos 8 caracteres.";
      errorPassword.classList.add('error');
      return;
    }

    if (password !== passwordConfirm) {
      document.getElementById('error-empty-field').style.display = 'none';
      errorPassword = document.getElementById('error-password-mismatch')
      errorPassword.style.display = 'block';
      errorPassword.textContent = "As senhas não coincidem.";
      errorPassword.classList.add('error');
      return;
    }
    
    fetch("http://localhost:8080/users/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
          name, email, password, passwordConfirm, document: documentNumber, companyName, type
      }),
    })
    .then(res => {
          if (!res.ok) throw new Error("Erro ao cadastrar");
          return res.json();
        })
        .then(data => {
          console.log("Cadastro realizado com sucesso!", data);
          window.location.href = "/user/login";
        })
        .catch(err => {
          console.error("Erro:", err);
          alert("Erro ao cadastrar usuário.");
        });
  });
});
