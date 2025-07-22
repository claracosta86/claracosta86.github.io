package api

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "os"
    "encoding/csv"

    "poc2/back/model"
    "poc2/back/lib"
    "poc2/back/service"

)

type UserHandler struct {
    userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// /save-user [POST]
func (h *UserHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
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
    fmt.Printf("Usuário recebido: %+v\n", user)

    err = h.userService.RegisterUser(user)
    err = lib.WriteUserToCSV(user)
    if err != nil {
        http.Error(w, "Erro ao salvar usuário", http.StatusInternalServerError)
        return
    }

    log.Printf("Usuário recebido: %+v\n", user)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"sucesso"}`))
}

// /fetch-users [GET]
func (h *UserHandler) HandleGetUserData(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    user, err := lib.ReadUsersFromCSV(user.ID)
    if err != nil {
        http.Error(w, "Erro ao ler usuários", http.StatusInternalServerError)
        return
    }

    userFavoritesList := service.FetchUserFavorites(user.ID)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(userFavoritesList)
}

// /update-user [PUT]
func (h *UserHandler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    err = h.userService.UpdateUser(user)
    if err != nil {
        http.Error(w, "Erro ao atualizar usuário", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"usuário atualizado com sucesso"}`))
}