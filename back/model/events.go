package model

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
	Image          string     `json:"image"` // Image associated with the event
}

type EventCollection []Event