package api

import (
    "encoding/json"
    "log"
    "net/http"

    "poc2/back/model"

)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) HandleUserSave(w http.ResponseWriter, r *http.Request) {
	 if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    // Aqui você pode salvar o user em banco, arquivo, etc
    log.Printf("Usuário recebido: %+v\n", user)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"sucesso"}`))
}
