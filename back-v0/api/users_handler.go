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

const userIDKey = "userID"

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
// /users/profile/:id/edit [PATCH]
func (h *UserHandler) HandleEditUserProfile(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPatch {
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
// /users/profile/:id/change-password [PATCH]
func (h *UserHandler) HandleChangeUserPassword(w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodPatch {
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
// /users/favorites/:id/:type/:favoriteID/add [POST]
func (h *UserHandler) HandleAddToFavorites(w http.ResponseWriter, r *http.Request) {
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

    var favorite model.UserFavorite
    err = json.NewDecoder(r.Body).Decode(&favorite)
    if err != nil {
        http.Error(w, "Dados inválidos", http.StatusBadRequest)
        return
    }

    err = h.userService.AddToFavorites(ctx, userID, favorite)
    if err != nil {
        http.Error(w, "Erro ao adicionar evento aos favoritos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// @Remove an event from user favorites
// /users/favorites/:id/:type/:favoriteID/delete [POST]
func (h *UserHandler) HandleDeleteFromFavorites(w http.ResponseWriter, r *http.Request) {
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

    strFavoriteID := chi.URLParam(r, "favoriteID")
    favoriteID, err := strconv.Atoi(strFavoriteID)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    err = h.userService.DeleteFromFavorites(ctx,userID, favoriteID)
    if err != nil {
        http.Error(w, "Erro ao remover evento dos favoritos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}