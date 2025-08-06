package model

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

func (t TouristAttraction) IsValid() bool {
	return t.ID > 0 && t.Title != "" && t.Description != "" && t.OpenDays != "" && t.Location != "" && t.Price >= 0 && t.OpenTime != ""
}

func (t *TouristAttractionCollection) IsEmpty() bool {
	if t == nil || len(*t) == 0 {
		return true
	}
	return false
}