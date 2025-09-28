package cultural

type Cultural struct{
	Events EventCollection `json:"events"`
	TouristAttractions TouristAttractionCollection `json:"touristAttractions"`
}

type TouristAttraction struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OpenDays    string `json:"open_days"`
	OpenTime    string `json:"open_time"` // Time in HH:MM format
	Location    string `json:"location"`
	Price      string `json:"price"` // Price in R$
	IsAccessible bool   `json:"is_accessible"` 
	OrganizerID int  `json:"organizerID"`
	OrganizerEmail string   `json:"organizerEmail"`
	Image      string `json:"image"`
}

type TouristAttractionCollection []TouristAttraction

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   string    `json:"start_date"` // YYYY-MM-DD HH:MM
	EndDate     string    `json:"end_date"`   // YYYY-MM-DD HH:MM
	DurationTime string    `json:"duration_time"` // HH:MM
	Location     string    `json:"location"`
	Price       string    `json:"price"` // Price in R$
	IsAccessible bool      `json:"is_accessible"` // Indica se um evento tem atenção à acessibilidade
	OrganizerID int  `json:"organizerID"`
	OrganizerEmail string   `json:"organizerEmail"`
	Image       string     `json:"image"` // Image associated with the event
}

type EventCollection []Event
