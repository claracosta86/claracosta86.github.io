document.addEventListener('DOMContentLoaded', function () {
    const userIconForm = document.getElementById('userIconForm');
    userIconForm.addEventListener('submit', function(event) {
        event.preventDefault();
        const userID = localStorage.getItem("userID");
        if (!userID) {
            window.location.href = "/login";
        }

        console.log("Usuário logado:", userID);

        fetch("http://localhost:8080/users/" + userID + "/profile/", {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        })
        .then(res => {
            if (!res.ok) throw new Error("Erro ao buscar dados do usuário.")
            return res.json();
        })
        .then(userData => {
            console.log("Login OK:", userData);
            window.location.href = "/user/profile";
        })
        .catch(err => {
            console.error(err);
            if (err.status === 404) {
                const errorLogin = document.getElementById('error-login')
                errorLogin.style.display = 'block';
                errorLogin.textContent = "Por favor, insira sua senha.";
                errorLogin.classList.add('error');
                return;
            }
        });
    });
});