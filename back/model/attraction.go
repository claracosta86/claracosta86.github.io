package model

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
	Image string `json:"image"` // Image associated with the tourist attraction
}

type TouristAttractionCollection []TouristAttraction 

