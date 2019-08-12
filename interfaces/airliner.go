package interfaces

import "github.com/fenr1s/back-end-take-home/domain/models"

//Airliner contract
type Airliner interface {
	GetAirlines() (airlines []*models.Airline, err error)
}
