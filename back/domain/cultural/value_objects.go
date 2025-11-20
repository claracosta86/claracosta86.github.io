package cultural

import "errors"

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
