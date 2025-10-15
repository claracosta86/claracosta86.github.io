package commentary

type Commentary struct {
	CulturalType string `json:"cultural_type"`
	Commentary   string `json:"commentary"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	ID           int    `json:"id"`
	CulturalID   int    `json:"cultural_id"`
	UserID      int    `json:"user_id"`
}

type Commentaries []Commentary