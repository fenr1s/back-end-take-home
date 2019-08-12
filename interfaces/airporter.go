package interfaces

import "github.com/fenr1s/back-end-take-home/domain/models"

//Airporter contract
type Airporter interface {
	GetAirports() (airports []*models.Airport, err error)
	CheckExistance(airports []*models.Airport, iata3 string) (exists bool, err error)
}
