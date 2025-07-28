package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Email    string `json:"email"`
	Document      string `json:"document"` // enum: [CPF, CNPJ]
	CompanyName string `json:"company_name"`
	Password string `json:"password,omitempty"` // Omit password in JSON responses
	PasswordConfirm string `json:"password_confirm,omitempty"` // Omit password confirm in JSON responses
	Role     string `json:"role"` //  enum: ["organizer", "common"]
}

type UserCollection []User


func (u User) IsValid() bool {
	 if u.Name == "" || u.Email == "" || u.Password == "" ||
        u.Role == "" || u.Document == "" || (u.PasswordConfirm != u.Password) {
        return false
    }

    if u.Role == "organizer" && u.CompanyName == "" {
        return false
    }

    return true
}

func (u *UserCollection) IsEmpty() bool {
	if u == nil || len(*u) == 0 {
		return true
	}
	return false
}

type UserFavorites struct {
	Events              EventCollection              `json:"events"`
	TouristAttractions   TouristAttractionCollection   `json:"tourist_attractions"`
}