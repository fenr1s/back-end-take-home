package services_test

import (
	"errors"
	"testing"

	"github.com/fenr1s/back-end-take-home/config"
	"github.com/fenr1s/back-end-take-home/domain/models"
	"github.com/fenr1s/back-end-take-home/domain/services"
	"github.com/fenr1s/back-end-take-home/interfaces/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAirportService(t *testing.T) {
	t.Run("ShouldGetAirports", shouldGetAirports)
	t.Run("ShouldReturnErrorGettingAirports", shouldReturnErrorGettingAirports)
	t.Run("ShouldCheckExistance", shouldCheckExistance)
	t.Run("ShouldCheckExistanceFalse", shouldCheckExistanceFalse)
	t.Run("InvalidateExistanceCheck", invalidateExistanceCheck)
}

func shouldGetAirports(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	line := []string{"test", "test", "test", "test", "40.63980103", "40.63980103"}
	fakeLines := [][]string{}
	fakeLines = append(fakeLines, line)
	fakeLines = append(fakeLines, line)
	fakeFileReader.On("ReadFromFile", config.CSV_PATH+"/airports.csv").Return(fakeLines, nil)
	airportService := &services.AirportService{FileReader: fakeFileReader}
	airports, err := airportService.GetAirports()

	assert.Equal(t, len(airports), 1)
	assert.Nil(t, err)
}
func shouldReturnErrorGettingAirports(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	line := []string{"test", "test", "test", "test", "40.63980103", "40.63980103"}
	fakeLines := [][]string{}
	fakeLines = append(fakeLines, line)
	fakeLines = append(fakeLines, line)
	fakeFileReader.On("ReadFromFile", config.CSV_PATH+"/airports.csv").Return(fakeLines, errors.New("test error"))
	airportService := &services.AirportService{FileReader: fakeFileReader}
	_, err := airportService.GetAirports()

	assert.NotNil(t, err)
}

func shouldCheckExistance(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	airportService := &services.AirportService{FileReader: fakeFileReader}
	fakeAirports := []*models.Airport{}
	fakeAirport := &models.Airport{
		Name:      "test",
		City:      "test",
		Country:   "test",
		Iata3:     "test",
		Latitude:  40.63980103,
		Longitude: 40.63980103,
	}
	fakeAirports = append(fakeAirports, fakeAirport)
	exists, err := airportService.CheckExistance(fakeAirports, "test")
	assert.Equal(t, true, exists)
	assert.Nil(t, err)
}

func shouldCheckExistanceFalse(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	airportService := &services.AirportService{FileReader: fakeFileReader}
	fakeAirports := []*models.Airport{}
	fakeAirport := &models.Airport{
		Name:      "test",
		City:      "test",
		Country:   "test",
		Iata3:     "test",
		Latitude:  40.63980103,
		Longitude: 40.63980103,
	}
	fakeAirports = append(fakeAirports, fakeAirport)
	exists, err := airportService.CheckExistance(fakeAirports, "zxc")
	assert.Equal(t, false, exists)
	assert.Nil(t, err)
}

func invalidateExistanceCheck(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	airportService := &services.AirportService{FileReader: fakeFileReader}
	fakeAirports := []*models.Airport{}
	_, err := airportService.CheckExistance(fakeAirports, "zxc")
	assert.NotNil(t, err)
}
