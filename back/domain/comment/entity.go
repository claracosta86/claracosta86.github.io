package comment

const (
	CulturalTypeEvent             = "event"
	CulturalTypeTouristAttraction = "tourist_attraction"
)

type Comment struct {
	CulturalType string
	Comment      CommentContent
	CreatedAt    string
	ID           int
	CulturalID   int
	UserName     string
}

type Comments []Comment
