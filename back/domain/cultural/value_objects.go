package cultural

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type Price struct {
	value string
}

func NewPrice(value string) Price {
	if value == "" {
		return Price{value: "R$0,00"}
	}
	return Price{value: value}
}

func (p Price) String() string {
	return p.value
}

// Value implements the driver.Valuer interface
func (p Price) Value() (driver.Value, error) {
	return p.value, nil
}

// Scan implements the sql.Scanner interface
func (p *Price) Scan(value interface{}) error {
	if value == nil {
		p.value = ""
		return nil
	}
	switch v := value.(type) {
	case []byte:
		p.value = string(v)
	case string:
		p.value = v
	default:
		return fmt.Errorf("failed to scan Price: %v", value)
	}
	return nil
}

type Location struct {
	value string
}

func NewLocation(value string) (Location, error) {
	if value == "" {
		return Location{}, errors.New("location cannot be empty")
	}
	return Location{value: value}, nil
}

func (l Location) String() string {
	return l.value
}

// Value implements the driver.Valuer interface
func (l Location) Value() (driver.Value, error) {
	return l.value, nil
}

// Scan implements the sql.Scanner interface
func (l *Location) Scan(value interface{}) error {
	if value == nil {
		l.value = ""
		return nil
	}
	switch v := value.(type) {
	case []byte:
		l.value = string(v)
	case string:
		l.value = v
	default:
		return fmt.Errorf("failed to scan Location: %v", value)
	}
	return nil
}
