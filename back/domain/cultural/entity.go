package cultural

const (
	CulturalTypeEvent             = "event"
	CulturalTypeTouristAttraction = "tourist_attraction"
)

type Cultural struct {
	Events             []Event
	TouristAttractions []TouristAttraction
}

type TouristAttraction struct {
	ID             int
	Title          string
	Description    string
	WorkingHours   string
	Location       string
	Price          string
	IsAccessible   bool
	OrganizerID    int
	OrganizerEmail string
	Image          string
}
type Event struct {
	ID             int
	Title          string
	Description    string
	StartDate      string
	EndDate        string
	DurationHours  string
	Location       string
	Price          string
	IsAccessible   bool
	OrganizerID    int
	OrganizerEmail string
	Image          string
}
type CulturalList struct {
	ID   int
	Type string
}
