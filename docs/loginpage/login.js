document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.getElementById('loginForm');
    loginForm.addEventListener('submit', function(event) {
        event.preventDefault();

        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;

        if (!email && !password) {
            error = document.getElementById('error-empty-field')
            error.style.display = 'block';
            error.classList.add('error');
            error.textContent = "Por favor, insira seu e-mail e senha.";
            return;
        } else if (!email) {
            error = document.getElementById('error-empty-field')
            error.style.display = 'block';
            error.textContent = "Por favor, insira seu e-mail.";
            error.classList.add('error');
            return;
        } else if (!password) {
            error = document.getElementById('error-empty-field')
            error.style.display = 'block';
            error.textContent = "Por favor, insira sua senha.";
            error.classList.add('error');
            return;
        }

        fetch("http://localhost:8080/users/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ email, password }),
        })
            .then(res => {
                console.log("Response:", res.status);
                if (!res.ok) throw new Error("Erro ao fazer login");
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
                    if (err.status === 404) {
                        errorLogin = document.getElementById('error-login')
                        errorLogin.style.display = 'block';
                        errorLogin.textContent = "Por favor, insira sua senha.";
                        errorLogin.classList.add('error');
                        return;
                    }
                });
            })
        .catch(err => {   
            console.error("Erro:", err);
            if (err.status === 404) {
                errorLogin = document.getElementById('error-login')
                errorLogin.style.display = 'block';
                errorLogin.textContent = "Seu usuário ou senha estão incorretos.";
                errorLogin.classList.add('error');
                return;
            }
        });
    });
});

function showErrorModal(message) {
  const modal = document.getElementById('error');
  const messageContainer = document.getElementById('errorMessage');
  messageContainer.textContent = message;
  modal.classList.remove('hidden');
}

function closeErrorModal() {
  const modal = document.getElementById('error');
  modal.classList.add('hidden');
}