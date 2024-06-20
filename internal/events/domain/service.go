package domain

import (
	"errors"
	"fmt"
)

type spotService struct{}

func NewSpotService() *spotService {
	return &spotService{}
}

func (s *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	for i := range quantity {
		name := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		spot, err := NewSpot(event, name)
		if err != nil {
			return err
		}
		event.Spots = append(event.Spots, *spot)
	}

	return nil
}
