package model

type Commentary struct {
	ID         int    `json:"id"`
	CulturalID int    `json:"cultural_id"`
	UserName   string `json:"user_name"`
	CulturalType string `json:"cultural_type"`
	Commentary string `json:"commentary"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type GetCommentariesResponse struct {
	Commentaries []Commentary `json:"commentaries"`
}

type CreateCommentaryRequest struct {
	Commentary  string `json:"commentary"`
	CulturalType string `json:"cultural_type"`
	CulturalID  int    `json:"cultural_id"`
	UserID      int    `json:"user_id"`
}

type UpdateCommentaryRequest struct {
	Commentary string `json:"commentary"`
	ID         int    `json:"id"`
}

type DeleteCommentaryRequest struct {
	ID int `json:"id"`
}
