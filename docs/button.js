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
});

form.addEventListener('submit', function(event) {
  event.preventDefault();

  fetch("http://localhost:8080/list-user", {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  })
  .then(res => res.json())
});