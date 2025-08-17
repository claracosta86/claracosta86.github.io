package model

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

func (e Event) IsValid() bool {
	return e.ID > 0 && e.Title != "" && e.Description != "" && e.StartDate != "" && e.FinishDate != "" && e.Location != "" && e.Price >= 0 && e.DurationTime != ""
}

func (e *EventCollection) IsEmpty() bool {
	if e == nil || len(*e) == 0 {
		return true
	}
	return false
}