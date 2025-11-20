package model

type CreateCulturalRequest struct {
	ID                int                    `json:"id"`
	Title             string                 `json:"title"`
	Type              string                 `json:"type"`
	Description       string                 `json:"description"`
	Event             EventHours             `json:"event"`
	TouristAttraction TouristAttractionHours `json:"touristAttraction"`
	Location          string                 `json:"location"`
	Price             string                 `json:"price"`
	IsAccessible      bool                   `json:"isAccessible"`
	OrganizerID       int                    `json:"organizerID"`
	Image             string                 `json:"image"`
}
type CreateCulturalResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type EventHours struct {
	StartDate     string `json:"startDate"`     // YYYY-MM-DD HH:MM
	EndDate       string `json:"endDate"`       // YYYY-MM-DD HH:MM
	DurationHours string `json:"durationHours"` // HH:MM
}

type TouristAttractionHours struct {
	WorkingHours string `json:"workingHours"` // HH:MM
}
type GetCulturalResponse struct {
	ID                int                    `json:"id"`
	Title             string                 `json:"title"`
	Description       string                 `json:"description"`
	Event             EventHours             `json:"event"`
	TouristAttraction TouristAttractionHours `json:"touristAttraction"`
	Location          string                 `json:"location"`
	Price             string                 `json:"price"`
	IsAccessible      bool                   `json:"isAccessible"`
	Organizer         Organizer              `json:"organizer"`
	Image             string                 `json:"image"`
}
type UpdateCulturalRequest struct {
	ID                int                    `json:"id"`
	Title             string                 `json:"title"`
	Type              string                 `json:"type"`
	Description       string                 `json:"description"`
	Event             EventHours             `json:"event"`
	TouristAttraction TouristAttractionHours `json:"touristAttraction"`
	Location          string                 `json:"location"`
	Price             string                 `json:"price"`
	IsAccessible      bool                   `json:"isAccessible"`
	OrganizerID       int                    `json:"organizerID"`
	Image             string                 `json:"image"`
}

type Event struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	DurationHours string `json:"durationHours"`
	Location      string `json:"location"`
	Price         string `json:"price"`
	IsAccessible  bool   `json:"isAccessible"`
	OrganizerID   int    `json:"organizerID"`
	Image         string `json:"image"`
}

type TouristAttraction struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	Location     string `json:"location"`
	Price        string `json:"price"`
	IsAccessible bool   `json:"isAccessible"`
	OrganizerID  int    `json:"organizerID"`
	Image        string `json:"image"`
	WorkingHours string `json:"workingHours"`
}

type GetAllCulturaisResponse struct {
	Events             []Event             `json:"events"`
	TouristAttractions []TouristAttraction `json:"touristAttractions"`
}
