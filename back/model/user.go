package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"` // Omit password in JSON responses
	Role     string `json:"role"` // e.g., "admin", "user"
}

type UserList struct {
	Users []User `json:"users"`
}

type UserFavorites struct {
	Events              EventCollection              `json:"events"`
	TouristAttractions   TouristAttractionCollection   `json:"tourist_attractions"`
}

