
document.addEventListener('DOMContentLoaded', function () {
  const registerForm = document.getElementById('registerUser');
  registerForm.addEventListener('submit', function(event) {
    event.preventDefault();

  const name = document.getElementById('name').value;
  const email = document.getElementById('email').value;
  const password = document.getElementById('password').value;
  const passwordConfirm = document.getElementById('passwordConfirm').value;
  const documentNumber = document.getElementById('document').value;
  const role = document.getElementById('role').value;

  const companyNameElement = document.getElementById('companyName');
  const companyName = companyNameElement ? companyNameElement.value : "";

  fetch("http://localhost:8080/users/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
        name, email, password, passwordConfirm, document: documentNumber, companyName, role
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
})});
