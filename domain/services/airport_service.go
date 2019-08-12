package services

import (
	"errors"
	"strconv"

	"github.com/fenr1s/back-end-take-home/config"
	"github.com/fenr1s/back-end-take-home/domain/models"
	"github.com/fenr1s/back-end-take-home/interfaces"
)

//AirportService service layer from airport data
type AirportService struct {
	FileReader interfaces.FileReader
}

//GetAirports get airports from csv
func (a *AirportService) GetAirports() (airports []*models.Airport, err error) {
	lines, err := a.FileReader.ReadFromFile(config.CSV_PATH + "/airports.csv")
	if err != nil {
		return airports, err
	}
	for i, line := range lines {
		if i != 0 {
			lat, err := strconv.ParseFloat(line[4], 64)
			if err != nil {
				return nil, err
			}
			lng, err := strconv.ParseFloat(line[5], 64)
			if err != nil {
				return nil, err
			}
			airports = append(airports, &models.Airport{
				Name:      line[0],
				City:      line[1],
				Country:   line[2],
				Iata3:     line[3],
				Latitude:  lat,
				Longitude: lng,
			})
		}
	}

	return airports, err
}

//CheckExistance check if airport exists by its iata3
func (a *AirportService) CheckExistance(airports []*models.Airport, iata3 string) (exists bool, err error) {
	if len(airports) == 0 {
		return exists, errors.New("Cannot check airport existance")
	}

	for _, airport := range airports {
		if airport.Iata3 == iata3 {
			return true, err
		}
	}

	return exists, err
}
