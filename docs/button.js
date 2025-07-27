const saveUser = document.getElementById('saveUser');
const listUsers = document.getElementById('listUsers');
const resultado = document.getElementById('resultado');

saveUser.addEventListener('submit', function(event) {
  event.preventDefault();

  const nome = document.getElementById('nome').value;
  const email = document.getElementById('email').value;
  const password = document.getElementById('password').value;
  const role = document.getElementById('role').value;

  fetch("http://localhost:8080/users/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ nome, email, password, role }),
  })
  .then(res => res.json())
});
