package back

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

type Image struct {
	AltText string `json:"alt_text"` // Descriptive text for the image
}

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	OpenDays    string    `json:"open_days"`
	Location    string    `json:"location"`
	Price      float64    `json:"price"` // Price in USD
	OpenTime        string    `json:"open_time"` // Time in HH:MM format
	IsAccessible    bool      `json:"is_accessible"` // Indica se um evento tem atenção à acessibilidade
	Contact         string    `json:"contact"` // Contact information for the event
	Image          Image     `json:"image"` // Image associated with the event
}

type TouristAttraction struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"` // ISO 8601 format
	Location    string `json:"location"`
	Price      float64 `json:"price"` // Price in R$
	DurationTime    string `json:"duration_time"` // Time in HH:MM format
	IsAccessible bool   `json:"is_accessible"` // Indica se um evento tem atenção à acessibilidade
	Contact string `json:"contact"` // Contact information for the tourist attraction
	Image Image `json:"image"` // Image associated with the tourist attraction
}

type UserFavorites struct {
	Events               []Event              `json:"events"`
	TouristAttractions   []TouristAttraction   `json:"tourist_attractions"`
}

