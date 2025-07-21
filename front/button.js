const botao = document.querySelector("#meuBotao");
botao.addEventListener("click", () => {
  fetch("https://seu-backend-go.com/api/evento", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      nome: "Clara",
      idade: 23
    }),
  })
  .then(res => res.text())
  .then(console.log)
  .catch(console.error);
});
