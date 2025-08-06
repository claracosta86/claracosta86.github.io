
document.addEventListener('DOMContentLoaded', function () {
  const registerForm = document.getElementById('registerUser');
  registerForm.addEventListener('submit', function(event) {
    event.preventDefault();

    const errorElements = document.querySelectorAll('.error');
    errorElements.forEach(err => {
      err.style.display = 'none';
      err.textContent = '';
    });

    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const passwordConfirm = document.getElementById('passwordConfirm').value;
    const documentNumber = document.getElementById('document').value;
    const type = document.getElementById('userType').value;
    const companyName = type === "organizer" ? document.getElementById('companyName').value : "";

    if (!name || !email || !password || !passwordConfirm || !documentNumber || (type === "organizer" && !companyName) ) {
      error = document.getElementById('error-empty-field')
      error.style.display = 'block';
      error.textContent = "Por favor, preencha todos os campos obrigatórios.";
      error.classList.add('error');
      return;
    }

    if (type === "organizer" && documentNumber.length !== 14) {
      document.getElementById('error-wrong-document-number').style.display = 'none';
      errorDocument = document.getElementById('error-wrong-document-number')
      errorDocument.style.display = 'block';
      errorDocument.textContent = "Por favor, insira um CNPJ válido.";
      errorDocument.classList.add('error');
      return;
    }

    if (type === "common" && documentNumber.length !== 11) {
      document.getElementById('error-wrong-document-number').style.display = 'none';
      errorDocument = document.getElementById('error-wrong-document-number')
      errorDocument.style.display = 'block';
      errorDocument.textContent = "Por favor, insira um CPF válido.";
      errorDocument.classList.add('error');
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
          name, email, password, document: documentNumber, companyName, type
      }),
    })
    .then(res => {
          if (res.status === 409) {
            errorEmail = document.getElementById('error-email-exists')
            errorEmail.style.display = 'block';
            errorEmail.textContent = "Este e-mail já está cadastrado.";
            errorEmail.classList.add('error');
          }
          if (!res.ok) throw new Error("Erro ao cadastrar usuário");
          return res.json();
        })
        .then(data => {
          console.log("Cadastro realizado com sucesso!", data);
          window.location.href = "/user/login";
        })
        .catch(err => {
          console.error("Erro:", err);
        });
  });
});
