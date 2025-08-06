package api

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "strconv"   
    stderrors "errors"
    "strings"

   chi  "github.com/go-chi/chi/v5"

    "poc2/back/lib/errors"
    "poc2/back/model"
    "poc2/back/service"

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
// [405] Invalid HTTP method
// [400] Invalid data
// [409] Conflict - User already exists
// [500] Internal Server Error
// [201] User registered successfully
// /users/register [POST]
func (h *UserHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }   
    
    ctx := r.Context()

    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, errors.ErrInvalidCredentials.Error(), http.StatusBadRequest)
        return
    }
    fmt.Printf("Usuário recebido: %+v\n", user)

    err = h.userService.RegisterUser(ctx, user)
    if err != nil && strings.Contains(err.Error(), errors.ErrUserAlreadyExists.Error()) {
        http.Error(w, errors.ErrUserAlreadyExists.Error(), http.StatusConflict)
        return
    } else if err != nil {
        http.Error(w, "Erro ao salvar usuário", http.StatusInternalServerError)
        return
    }
    
    log.Printf("Usuário recebido: %+v\n", user)

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"status":"sucesso"}`))
}

// @Verify user login
// @Accept json
// [400] Invalid data
// [401] Invalid credentials
// [403] User not authorized
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] User logged in successfully
// /users/login [POST]
func (h *UserHandler) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    ctx := r.Context()

    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, errors.ErrInvalidCredentials.Error(), http.StatusBadRequest)
        return
    }

    userID, err := h.userService.GetUserIDByEmail(ctx, user.Email)
    if stderrors.Is(err, errors.ErrUserNotFound){
        http.Error(w, errors.ErrUserNotFound.Error(), http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
        return
    }

    verifyPassword, err := h.userService.VerifyUserPassword(ctx, userID, user.Password)
    if err != nil {
        http.Error(w, "Erro ao verificar senha", http.StatusInternalServerError)
        return
    }

    if !verifyPassword {
        http.Error(w, "Senha incorreta", http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(map[string]any{
		"userID": userID,
	})
}


// @Fetch user profile data
// /users/fetch/:id [GET]
// func (h *UserHandler) HandleGetUserData(w http.ResponseWriter, r *http.Request) {
//     if r.Method != http.MethodGet {
//         http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
//         return
//     }

    // strUserID := chi.URLParam(r, "id")
    // userID, err := strconv.Atoi(strUserID)
    // if err != nil {
    //     http.Error(w, "ID inválido", http.StatusBadRequest)
    //     return
    // }

    // user := model.User{
    //     ID: 86, // This should be replaced with actual user ID fetching logic
    //     Name: "John Doe",
    //     Email: "john.doe@example.com",
    // }


    // user, err := h.userService.GetUserDataByID(userID)
    // if err != nil {
    //     http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
    //     return
    // }

//     w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(user)
// }

// @Get user profile page
// [400] Invalid data
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [200] User data recovered successfully
// /users/profile/:id/ [GET]
func (h *UserHandler) HandleGetUserProfile(w http.ResponseWriter, r *http.Request) {    
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

    user, err := h.userService.GetUserDataByID(r.Context(), userID)
    if err != nil {
        if stderrors.Is(err, errors.ErrUserNotFound) {
            http.Error(w, "Usuário não encontrado", http.StatusNotFound)
            return
        }
        http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _ = json.NewEncoder(w).Encode(map[string]any{
        "userName": user.Name,
        "userEmail": user.Email,
        "userID": user.ID,
        "userType": user.Type,
    })
}

// @Change user name or email
// @Accept json
// [400] Invalid data
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [204] User profile edited in successfully
// /users/profile/:id/edit [PUT]
func (h *UserHandler) HandleEditUserProfile(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    ctx := r.Context()

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    var userUpdates model.User
    err = json.NewDecoder(r.Body).Decode(&userUpdates)
    if err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    userUpdates.ID = userID

    err = h.userService.UpdateUserProfile(ctx, userUpdates)
    if err != nil {
        http.Error(w, "Erro ao atualizar perfil", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// @Change user password
// @Accept json
// [400] Invalid data
// [401] Incorrect password
// [404] User not found
// [405] Invalid HTTP method
// [500] Internal Server Error
// [204] User password edited in successfully
// /users/profile/:id/change-password [PUT]
func (h *UserHandler) HandleChangeUserPassword(w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodPut {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    ctx := r.Context()

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    var passwordUpdate model.PasswordUpdate
    err = json.NewDecoder(r.Body).Decode(&passwordUpdate)
    if err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    err = h.userService.UpdateUserPassword(ctx, userID, passwordUpdate)
    if err != nil {
        if stderrors.Is(err, errors.ErrIncorrectPassword) {
            http.Error(w, "Senha inválida", http.StatusUnauthorized)
            return
        }
        http.Error(w, "Erro ao atualizar senha", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}


// @Get users favorites events and attractions
// /users/favorites/:id [GET]
func (h *UserHandler) HandleGetUserFavorites(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    ctx := r.Context()

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    favorites, err := h.userService.GetUserFavoritesByID(ctx,userID)
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

    ctx := r.Context()

    strUserID := chi.URLParam(r, "id")
    userID, err := strconv.Atoi(strUserID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    err = h.userService.DeleteUserByID(ctx, userID)
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

    ctx := r.Context()

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

    err = h.userService.AddEventToFavorites(ctx, userID, event)
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

    ctx := r.Context()

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

    err = h.userService.DeleteEventFromFavorites(ctx,userID, eventID)
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

    ctx := r.Context()

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

    err = h.userService.AddAttractionToFavorites(ctx,userID, attraction)
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

    ctx := r.Context()

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

    err = h.userService.DeleteAttractionFromFavorites(ctx,userID, attractionID)
    if err != nil {
        http.Error(w, "Erro ao remover atração dos favoritos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}