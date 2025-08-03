document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.getElementById('loginForm');
    loginForm.addEventListener('submit', function(event) {
        event.preventDefault();

        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;

        if (!email && !password) {
            errorEmail = document.getElementById('error-empty-email')
            errorPassword = document.getElementById('error-empty-password')
            
            errorEmail.style.display = 'block';
                        errorPassword.classList.add('error');

            errorPassword.style.display = 'block';
            errorPassword.textContent = "Por favor, preencha todos os campos.";
            return;
        } else if (!email) {
            errorEmail = document.getElementById('error-empty-email')
            errorEmail.style.display = 'block';
            errorEmail.textContent = "Por favor, insira seu e-mail.";
            errorEmail.classList.add('error');
            return;
        } else if (!password) {
            errorPassword = document.getElementById('error-empty-password')
            errorPassword.style.display = 'block';
            errorPassword.textContent = "Por favor, insira sua senha.";
            errorPassword.classList.add('error');
            return;
        }

        fetch("http://localhost:8080/users/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ email, password }),
        })
        .then(res => {
            if (!res.ok) throw new Error("Erro ao fazer login")
                return res.json();
        })
        .then(data => {
            const userID = data.userID; 
            
            fetch("http://localhost:8080/users/fetch/" + userID, {
                method: "GET",
                headers: { "Content-Type": "application/json" },
            })
            .then(res2 => {
                if (!res2.ok) throw new Error("Erro ao buscar dados do usuário.")
                return res2.json();
            })
            .then(userData => {
                console.log("Login OK:", userData);
                window.location.href = "/home";
            })
            .catch(err => {
                console.error(err);
                alert("Erro ao buscar dados do usuário.");
            });
        })
        .catch(err => {   
            console.error("Erro:", err);
            if (err.status === 404) {
                showErrorModal("Usuário não encontrado. Verifique o e-mail e tente novamente.");
                }
            alert("Erro ao fazer login.");
        });
    });
});

function showErrorModal(message) {
  const modal = document.getElementById('errorModal');
  const messageContainer = document.getElementById('errorMessage');
  messageContainer.textContent = message;
  modal.classList.remove('hidden');
}

function closeErrorModal() {
  const modal = document.getElementById('errorModal');
  modal.classList.add('hidden');
}