fetch("http://localhost:8080/save-user", {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({ nome: "Clara", email: "clara@email.com" }),
})
.then(res => res.json())
.then(data => console.log(data))
.catch(console.error);
