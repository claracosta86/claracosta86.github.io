package model

type TouristAttraction struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"` // YYYY-MM-DD
	Location    string `json:"location"`
	Price      float64 `json:"price"` // Price in R$
	DurationTime    string `json:"duration_time"` // HH:MM 
	IsAccessible bool   `json:"is_accessible"` 
	Contact string `json:"contact"` 
	Image string `json:"image"`
}

type TouristAttractionCollection []TouristAttraction 

func (t TouristAttraction) IsValid() bool {
	return t.ID > 0 && t.Title != "" && t.Description != "" && t.Date != "" && t.Location != "" && t.Price >= 0 && t.DurationTime != ""
}

func (t *TouristAttractionCollection) IsEmpty() bool {
	if t == nil || len(*t) == 0 {
		return true
	}
	return false
}