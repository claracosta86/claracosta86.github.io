package lib

import (
	"encoding/csv"
	"fmt"
	"os"
    "log"

	"poc2/back/model"

)

func WriteUserToCSV(user model.User) error {
    file, err := os.OpenFile("users.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    
    log.Printf("Escrevendo usuário no CSV: %+v\n", user)
    writer := csv.NewWriter(file)
    defer writer.Flush()

    return writer.Write([]string{fmt.Sprintf("%d", user.ID), user.Name, user.Email, user.Password, user.Role})
}


func ReadUsersFromCSV() ([]model.User, error) {
    file, err := os.Open("users.csv")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var users []model.User
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    for _, record := range records {
        if len(record) >= 2 {
            users = append(users, model.User{
                Name:  record[0],
                Email: record[1],
            })
        }
    }
    return users, nil
}
