package services_test

import (
	"errors"
	"testing"

	"github.com/fenr1s/back-end-take-home/config"
	"github.com/fenr1s/back-end-take-home/domain/services"
	"github.com/fenr1s/back-end-take-home/interfaces/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAirlineService(t *testing.T) {
	t.Run("ShouldGetAirlines", shouldGetAirlines)
	t.Run("ShouldReturnError", shouldReturnError)
}

func shouldGetAirlines(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	line := []string{"test", "test", "test", "test"}
	fakeLines := [][]string{}
	fakeLines = append(fakeLines, line)
	fakeLines = append(fakeLines, line)
	fakeFileReader.On("ReadFromFile", config.CSV_PATH+"/airlines.csv").Return(fakeLines, nil)
	airlineService := &services.AirlineService{FileReader: fakeFileReader}
	airlines, err := airlineService.GetAirlines()

	assert.Equal(t, len(airlines), 1)
	assert.Nil(t, err)
}
func shouldReturnError(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	line := []string{"test", "test", "test", "test"}
	fakeLines := [][]string{}
	fakeLines = append(fakeLines, line)
	fakeLines = append(fakeLines, line)
	fakeFileReader.On("ReadFromFile", config.CSV_PATH+"/airlines.csv").Return(fakeLines, errors.New("test error"))
	airlineService := &services.AirlineService{FileReader: fakeFileReader}
	_, err := airlineService.GetAirlines()

	assert.NotNil(t, err)
}
