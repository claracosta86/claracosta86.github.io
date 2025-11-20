package model

type CreateCulturalRequest struct {
	ID                int                    `json:"id"`
	Title             string                 `json:"title"`
	Type              string                 `json:"type"`
	Description       string                 `json:"description"`
	Event             EventDateInformation   `json:"event"`
	TouristAttraction TouristAttractionHours `json:"tourist_attraction"`
	Location          string                 `json:"location"`
	Price             string                 `json:"price"`         // Price in R$
	IsAccessible      bool                   `json:"is_accessible"` // Indica se um evento tem atenção à acessibilidade
	OrganizerID       int                    `json:"organizer_id"`  // Contact information for the event
	Image             string                 `json:"image"`         // Image associated with the event
}

type EventDateInformation struct {
	StartDate    string `json:"start_date"`    // YYYY-MM-DD HH:MM
	EndDate      string `json:"end_date"`      // YYYY-MM-DD HH:MM
	WorkingHours string `json:"working_hours"` // HH:MM
}

type TouristAttractionHours struct {
	OpenDays string `json:"open_days"`
	OpenTime string `json:"open_time"` // Time in HH:MM format
}

type CulturalResponse struct {
	ID                int                    `json:"id"`
	Title             string                 `json:"title"`
	Description       string                 `json:"description"`
	Event             EventDateInformation   `json:"event"`
	TouristAttraction TouristAttractionHours `json:"tourist_attraction"`
	Location          string                 `json:"location"`
	Price             string                 `json:"price"`         // Price in R$
	IsAccessible      bool                   `json:"is_accessible"` // Indica se um evento tem atenção à acessibilidade
	Organizer         Organizer              `json:"organizer"`     // Contact information for the event
	Image             string                 `json:"image"`         // Image associated with the event
}

type UpdateCulturalRequest struct {
	ID                int                    `json:"id"`
	Title             string                 `json:"title"`
	Type              string                 `json:"type"`
	Description       string                 `json:"description"`
	Event             EventDateInformation   `json:"event"`
	TouristAttraction TouristAttractionHours `json:"tourist_attraction"`
	Location          string                 `json:"location"`
	Price             string                 `json:"price"`         // Price in R$
	IsAccessible      bool                   `json:"is_accessible"` // Indica se um evento tem atenção à acessibilidade
	OrganizerID       int                    `json:"organizer_id"`  // Contact information for the event
	Image             string                 `json:"image"`         // Image associated with the event
}

type CreateCulturalResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type AllCulturaisResponse struct {
	Events             []Event             `json:"events"`
	TouristAttractions []TouristAttraction `json:"tourist_attractions"`
}

type Image struct {
	URL string `json:"url"`
}

type Event struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	StartDate    string `json:"start_date"` // YYYY-MM-DD HH:MM
	EndDate      string `json:"end_date"`   // YYYY-MM-DD HH:MM
	WorkingHours string `json:"working_hours"`
	Location     string `json:"location"`
	Price        string `json:"price"`         // Price in R$
	IsAccessible bool   `json:"is_accessible"` // Indica se um evento tem atenção à acessibilidade
	OrganizerID  int    `json:"organizer_id"`  // Contact information for the event
	Image        string `json:"image"`         //
}

type TouristAttraction struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	Location     string `json:"location"`
	Price        string `json:"price"`         // Price in R$
	IsAccessible bool   `json:"is_accessible"` // Indica se um evento tem atenção à acessibilidade
	OrganizerID  int    `json:"organizer_id"`  // Contact information for the event
	Image        string `json:"image"`         //
	OpenDays     string `json:"open_days"`
	OpenTime     string `json:"open_time"`
}
