package model

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	OpenDays    string    `json:"open_days"`
	Location    string    `json:"location"`
	Price      float64    `json:"price"` // Price in R$
	OpenTime        string    `json:"open_time"` // Time in HH:MM format
	IsAccessible    bool      `json:"is_accessible"` // Indica se um evento tem atenção à acessibilidade
	Contact         string    `json:"contact"` // Contact information for the event
	Image          string     `json:"image"` // Image associated with the event
}

type EventCollection []Event

func (e Event) IsValid() bool {
	return e.ID > 0 && e.Title != "" && e.Description != "" && e.OpenDays != "" && e.Location != "" && e.Price >= 0 && e.OpenTime != ""
}

func (e *EventCollection) IsEmpty() bool {
	if e == nil || len(*e) == 0 {
		return true
	}
	return false
}