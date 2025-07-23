package api

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"

   chi  "github.com/go-chi/chi/v5"

    "poc2/back/model"
    "poc2/back/service"
    "poc2/back/lib/utils"

)

type UserHandler struct {
    userService service.UserService
    eventService service.EventService
    attractionService service.AttractionService
}

func NewUserHandler(userService service.UserService, eventService service.EventService, attractionService service.AttractionService) *UserHandler {
	return &UserHandler{
		userService: userService,
        eventService: eventService,
        attractionService: attractionService,
	}
}

// /users/register [POST]
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
    err = utils.WriteUserToCSV(user)
    if err != nil {
        http.Error(w, "Erro ao salvar usuário", http.StatusInternalServerError)
        return
    }

    log.Printf("Usuário recebido: %+v\n", user)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"sucesso"}`))
}

// /users/fetch [GET]
func (h *UserHandler) HandleGetUserData(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    users := []model.User{
        {
            ID: 1,
            Name: "Clara",
            Email: "clara004.costa@gmail.com",
        },
        {
            ID: 2,
            Name: "Costa",
            Email: "claraufmg4@gmail.com",
        },
    }
    // user, err := lib.ReadUsersFromCSV(user.ID)
    // if err != nil {
    //     http.Error(w, "Erro ao ler usuários", http.StatusInternalServerError)
    //     return
    // }

    userFavoritesList, _ := h.userService.FetchUserDataByID(users[0].ID)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(userFavoritesList)
}

// /users/update [PUT]
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

    err = h.userService.UpdateUserData(user)
    if err != nil {
        http.Error(w, "Erro ao atualizar usuário", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"usuário atualizado com sucesso"}`))
}

// /users/favorites/:id [GET]
func (h *UserHandler) HandleGetUserFavorites(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    userID := chi.URLParam(r, "id")
    favorites, err := h.userService.FetchUserFavorites(userID)
    if err != nil {
        http.Error(w, "Erro ao buscar favoritos", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(favorites)
}

// /users/delete/:id [DELETE]
func (h *UserHandler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    userID := r.Context().Value("userID").(string)
    err := h.userService.DeleteUserByID(userID)
    if err != nil {
        http.Error(w, "Erro ao deletar usuário", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"usuário deletado com sucesso"}`))
}

// /users/favorites/:id/event/add [POST]
func (h *UserHandler) HandleAddEventToFavorites(w http.ResponseWriter, r *http.Request) {
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

    err = h.userService.AddEventToFavorites(user)
    if err != nil {
        http.Error(w, "Erro ao adicionar evento aos favoritos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// /users/favorites/:id/attraction/add [POST]
func (h *UserHandler) HandleAddAttractionToFavorites(w http.ResponseWriter, r *http.Request) {
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

    err = h.userService.AddAttractionToFavorites(user)
    if err != nil {
        http.Error(w, "Erro ao adicionar atração aos favoritos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}