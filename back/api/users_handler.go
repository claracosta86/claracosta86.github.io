package api

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "strconv"   

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

// @Register a new user handler
// @Accept json
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

// @Fetch user profile data
// /users/fetch/:id [GET]
func (h *UserHandler) HandleGetUserData(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    user, err := h.userService.GetUserDataByID(userID)
    if err != nil {
        http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// @Update user profile data
// @Accept json
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

// @Get users favorites events and attractions
// /users/favorites/:id [GET]
func (h *UserHandler) HandleGetUserFavorites(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    favorites, err := h.userService.GetUserFavoritesByID(userID)
    if err != nil {
        http.Error(w, "Erro ao buscar favoritos", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(favorites)
}

// @Deletes an user from the system
// /users/delete/:id [DELETE]
func (h *UserHandler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    err = h.userService.DeleteUserByID(userID)
    if err != nil {
        http.Error(w, "Erro ao deletar usuário", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"usuário deletado com sucesso"}`))
}

// @Add an event to user favorites
// @Accept json
// /users/favorites/:id/event/add [POST]
func (h *UserHandler) HandleAddEventToFavorites(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "UserID inválido", http.StatusBadRequest)
        return
    }

    var event model.Event
    err = json.NewDecoder(r.Body).Decode(&event)
    if err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    err = h.userService.AddEventToFavorites(userID, event)
    if err != nil {
        http.Error(w, "Erro ao adicionar evento aos favoritos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// @Remove an event from user favorites
// /users/favorites/:id/event/:eventID/delete [POST]
func (h *UserHandler) HandleDeleteEventFromFavorites(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    strEventID := chi.URLParam(r, "eventID")
    eventID, err := strconv.Atoi(strEventID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    err = h.userService.DeleteEventFromFavorites(userID, eventID)
    if err != nil {
        http.Error(w, "Erro ao remover evento dos favoritos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// @Add an attraction from user favorites
// @Accept json
// /users/favorites/:id/attraction/add [POST]
func (h *UserHandler) HandleAddAttractionToFavorites(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    var attraction model.TouristAttraction
    err = json.NewDecoder(r.Body).Decode(&attraction)
    if err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    err = h.userService.AddAttractionToFavorites(userID, attraction)
    if err != nil {
        http.Error(w, "Erro ao adicionar atração aos favoritos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// @Remove an attraction from user favorites
// /users/favorites/:id/attraction/:attractionID/delete [POST]
func (h *UserHandler) HandleDeleteAttractionFromFavorites(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    strAttractionID := chi.URLParam(r, "attractionID")
    attractionID, err := strconv.Atoi(strAttractionID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    err = h.userService.DeleteAttractionFromFavorites(userID, attractionID)
    if err != nil {
        http.Error(w, "Erro ao remover atração dos favoritos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}