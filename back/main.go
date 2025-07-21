// package main	

// import "fmt"

// func main() {

// 	user := handle_user()
// 	event := add_new_event()
// 	attraction := add_new_tourist_attraction()



// }

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

	"poc2/back/api"

)

type MeuEvento struct {
    Nome string `json:"nome"`
    Idade int   `json:"idade"`
}

func tratarEvento(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var evento MeuEvento
    err := json.NewDecoder(r.Body).Decode(&evento)
    if err != nil {
        http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
        return
    }

    fmt.Printf("Evento recebido: %+v\n", evento)

    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Evento tratado com sucesso")
}

// func main() {
//     http.HandleFunc("/claracosta86.github.io/", tratarEvento)
//     fmt.Println("Servidor rodando em http://localhost:8080")
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }


func enableCors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "https://seu-usuario.github.io")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == http.MethodOptions {
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/save-user", api.NewUserHandler().HandleUserSave)

    log.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", enableCors(mux)))
}