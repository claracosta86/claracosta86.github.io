package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"` // Omit password in JSON responses
	Role     string `json:"role"` // e.g., "admin", "user"
}

type UserFavorites struct {
	Events              EventCollection              `json:"events"`
	TouristAttractions   TouristAttractionCollection   `json:"tourist_attractions"`
}

type UserCollection []User 


func (u User) IsValid() bool {
	return u.ID > 0 && u.Name != "" && u.Email != "" && u.Password != "" && u.Role != ""
}

func (u *UserCollection) IsEmpty() bool {
	if u == nil || len(*u) == 0 {
		return true
	}
	return false
}
