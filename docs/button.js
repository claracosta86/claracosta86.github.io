const form = document.getElementById('userForm');
const resultado = document.getElementById('resultado');

form.addEventListener('submit', function(event) {
  event.preventDefault();

  const nome = document.getElementById('nome').value;
  const email = document.getElementById('email').value;

  fetch("http://localhost:8080/save-user", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ nome, email }),
  })
  .then(res => res.json())
  .then(data => {
    resultado.textContent = "Resposta do servidor:\n" + JSON.stringify(data, null, 2);
  })
  .catch(err => {
    resultado.textContent = "Erro: " + err.message;
  });
});
