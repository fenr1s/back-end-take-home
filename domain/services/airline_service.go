package services

import (
	"github.com/fenr1s/back-end-take-home/config"
	"github.com/fenr1s/back-end-take-home/domain/models"
	"github.com/fenr1s/back-end-take-home/interfaces"
)

//AirlineService contains logic of the airline context
type AirlineService struct {
	FileReader interfaces.FileReader
}

//GetAirlines return a array of airlines based on the csv and a possible error
func (a *AirlineService) GetAirlines() (airlines []*models.Airline, err error) {
	lines, err := a.FileReader.ReadFromFile(config.CSV_PATH + "/airlines.csv")
	if err != nil {
		return airlines, err
	}

	for i, line := range lines {
		if i != 0 {
			airlines = append(airlines, &models.Airline{
				Name:           line[0],
				TwoDigitCode:   line[1],
				ThreeDigitCode: line[2],
				Country:        line[3],
			})
		}
	}

	return airlines, err
}
