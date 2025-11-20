package model

type Comment struct {
	ID           int    `json:"id"`
	CulturalID   int    `json:"cultural_id"`
	UserName     string `json:"user_name"`
	CulturalType string `json:"cultural_type"`
	Comment      string `json:"comment"`
	CreatedAt    string `json:"created_at"`
}

type GetCommentsResponse struct {
	Comments []Comment `json:"commentaries"`
}

type CreateCommentRequest struct {
	Comment      string `json:"comment"`
	CulturalType string `json:"cultural_type"`
	CulturalID   int    `json:"cultural_id"`
	UserID       int    `json:"user_id"`
}

type UpdateCommentsRequest struct {
	Comment string `json:"comment"`
	ID      int    `json:"id"`
}

type DeleteCommentsRequest struct {
	ID int `json:"id"`
}
