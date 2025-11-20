package comment

type Comment struct {
	CulturalType string `json:"cultural_type"`
	Comment      string `json:"comment"`
	CreatedAt    string `json:"created_at"`
	ID           int    `json:"id"`
	CulturalID   int    `json:"cultural_id"`
	UserName     string `json:"user_name"`
}

type Comments []Comment
