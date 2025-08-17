package main

import (
    "log"
    "net/http"
	"database/sql"

    _ "github.com/go-sql-driver/mysql"

    "poc2/routes"
    "poc2/back/infrastructure/container"
)

func main() {
    log.Println("Server running at http://localhost:8080/")

	dsn := "claracosta86:sua_senha@tcp(127.0.0.1:3306)/POCII"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create dependency injection container
	container := container.NewContainer(db)

	// Setup routes with the new DDD structure
	router := routes.SetupRoutes(container)

	http.ListenAndServe(":8080", router)
}