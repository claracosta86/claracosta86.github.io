package model

type Comment struct {
	ID           int    `json:"id"`
	CulturalID   int    `json:"culturalID"`
	UserName     string `json:"userName"`
	CulturalType string `json:"culturalType"`
	Comment      string `json:"comment"`
	CreatedAt    string `json:"createdAt"`
}

type GetCommentsResponse struct {
	Comments []Comment `json:"comments"`
}

type CreateCommentRequest struct {
	Comment      string `json:"comment"`
	CulturalType string `json:"culturalType"`
	CulturalID   int    `json:"culturalID"`
	UserID       int    `json:"userID"`
}
