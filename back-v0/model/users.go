package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Email    string `json:"email"`
	Document      string `json:"document"` // enum: [CPF, CNPJ]
	CompanyName string `json:"companyName"`
	Password string `json:"password"`
	Type     string `json:"type"` //  enum: ["organizer", "common"]
}

type UserCollection []User

func (u *UserCollection) IsEmpty() bool {
	if u == nil || len(*u) == 0 {
		return true
	}
	return false
}

type UserFavorites struct {
	Events              EventCollection              `json:"events"`
	TouristAttractions   TouristAttractionCollection   `json:"touristAttractions"`
}

type UserFavorite struct {
	ID		int    `json:"id"`
	Type    string `json:"type"` // enum: ["event", "tourist_attraction"]
}

type Organizer struct {
	ID      int    `json:"id"`
	Contact string `json:"contact"` // Contact information for the organizer
}

type PasswordUpdate struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}