
registerUser.addEventListener('submit', function(event) {
  event.preventDefault();

  const name = document.getElementById('name').value;
  const email = document.getElementById('email').value;
  const password = document.getElementById('password').value;
  const passwordConfirm = document.getElementById('passwordConfirm').value;
  const document = document.getElementById('document').value;
  const companyName = document.getElementById('companyName').value;
  const role = document.getElementById('role').value;

  fetch("http://localhost:8080/users/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
        name, email, password, passwordConfirm, document, companyName, role
    }),
  })
  .then(res => {
        if (!res.ok) throw new Error("Erro ao cadastrar");
        return res.json();
      })
      .then(data => {
        console.log("Cadastro realizado com sucesso!", data);
        if (role === "organizer") {
          window.location.href = "../login/login-organizer.html";
        } else {
          window.location.href = "../login/login-common.html";
        }
      })
      .catch(err => {
        console.error("Erro:", err);
        alert("Erro ao cadastrar usuário.");
      });
});
