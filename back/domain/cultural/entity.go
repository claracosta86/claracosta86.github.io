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
	Price      float64 `json:"price"` // Price in R$
	IsAccessible bool   `json:"is_accessible"` 
	Organizer   Organizer `json:"organizer"` // Contact information for the attraction
	Image      string `json:"image"`
}

type TouristAttractionCollection []TouristAttraction

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   string    `json:"start_date"` // YYYY-MM-DD HH:MM
	FinishDate  string    `json:"finish_date"` // YYYY-MM-DD HH:MM
	DurationTime string    `json:"duration_time"` // HH:MM
	Location     string    `json:"location"`
	Price       float64    `json:"price"` // Price in R$
	IsAccessible bool      `json:"is_accessible"` // Indica se um evento tem atenção à acessibilidade
	Organizer    Organizer  `json:"organizer"` // Contact information for the event
	Image       string     `json:"image"` // Image associated with the event
}

type EventCollection []Event