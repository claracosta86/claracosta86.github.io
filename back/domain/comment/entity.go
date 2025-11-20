package comment

type Comment struct {
	CulturalType string
	Comment      string
	CreatedAt    string
	ID           int
	CulturalID   int
	UserName     string
}

type Comments []Comment
