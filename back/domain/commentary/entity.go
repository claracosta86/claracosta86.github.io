package commentary

type Commentary struct {
	CulturalType string `json:"cultural_type"`
	Commentary   string `json:"commentary"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	ID           int    `json:"id"`
	CulturalID   int    `json:"cultural_id"`
	UserName     string `json:"user_name"`
}

type Commentaries []Commentary