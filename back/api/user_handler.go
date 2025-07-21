package api

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"

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
    fmt.Printf("Usuário recebido: %+v\n", user)

    err = writeUserToCSV(user)
    if err != nil {
        http.Error(w, "Erro ao salvar usuário", http.StatusInternalServerError)
        return
    }

    // Aqui você pode salvar o user em banco, arquivo, etc
    log.Printf("Usuário recebido: %+v\n", user)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"sucesso"}`))
}

func (h *UserHandler) HandleListUsers(w http.ResponseWriter, r *http.Request) {
    users, err := readUsersFromCSV()
    if err != nil {
        http.Error(w, "Erro ao ler usuários", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func writeUserToCSV(user User) error {
    file, err := os.OpenFile("users.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    return writer.Write([]string{user.Name, user.Email})
}

func readUsersFromCSV() ([]User, error) {
    file, err := os.Open("users.csv")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var users []User
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    for _, record := range records {
        if len(record) >= 2 {
            users = append(users, User{
                Name:  record[0],
                Email: record[1],
            })
        }
    }
    return users, nil
}
