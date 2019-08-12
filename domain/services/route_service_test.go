package services_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/fenr1s/back-end-take-home/config"
	"github.com/fenr1s/back-end-take-home/domain/services"
	"github.com/fenr1s/back-end-take-home/interfaces/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRouteService(t *testing.T) {
	t.Run("ShouldGetRoutes", shouldGetRoutes)
	t.Run("ShouldGetErrorGettingRoutes", shouldGetErrorGettingRoutes)
	t.Run("ShouldFindShortestPath", shouldFindShortestPath)
	t.Run("ShouldFindRouteByOrigin", shouldFindRouteByOrigin)
	t.Run("ShouldBuildPath", shouldBuildPath)
}

func shouldGetRoutes(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	line := []string{"test", "test", "test"}
	fakeLines := [][]string{}
	fakeLines = append(fakeLines, line)
	fakeLines = append(fakeLines, line)
	fakeFileReader.On("ReadFromFile", config.CSV_PATH+"/routes.csv").Return(fakeLines, nil)
	routeService := &services.RouteService{FileReader: fakeFileReader}
	routes, err := routeService.GetRoutes()

	assert.Equal(t, len(routes), 1)
	assert.Nil(t, err)
}

func shouldGetErrorGettingRoutes(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	line := []string{"test", "test", "test"}
	fakeLines := [][]string{}
	fakeLines = append(fakeLines, line)
	fakeLines = append(fakeLines, line)
	fakeFileReader.On("ReadFromFile", config.CSV_PATH+"/routes.csv").Return(fakeLines, errors.New("test"))
	routeService := &services.RouteService{FileReader: fakeFileReader}
	_, err := routeService.GetRoutes()
	assert.NotNil(t, err)
}

func shouldFindShortestPath(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	line := []string{"AC", "YYZ", "JFK"}
	fakeLines := [][]string{}
	fakeLines = append(fakeLines, line)
	fakeLines = append(fakeLines, line)
	fakeFileReader.On("ReadFromFile", config.CSV_PATH+"/routes.csv").Return(fakeLines, nil)
	routeService := &services.RouteService{FileReader: fakeFileReader}
	routes, err := routeService.GetRoutes()
	path := routeService.FindShortestPath("YYZ", "JFK", routes)

	assert.Equal(t, "YYZ -> JFK", path)
	assert.Nil(t, err)
}

func shouldFindRouteByOrigin(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	line := []string{"AC", "YYZ", "JFK"}
	fakeLines := [][]string{}
	fakeLines = append(fakeLines, line)
	fakeLines = append(fakeLines, line)
	fakeFileReader.On("ReadFromFile", config.CSV_PATH+"/routes.csv").Return(fakeLines, nil)
	routeService := &services.RouteService{FileReader: fakeFileReader}
	routes, err := routeService.GetRoutes()
	founds, err := routeService.FindRoutesByOrigin("YYZ", routes)

	assert.Equal(t, len(founds), 1)
	assert.Nil(t, err)
}

func shouldBuildPath(t *testing.T) {
	fakeFileReader := &mocks.FileReader{}
	line := []string{"AC", "YYZ", "JFK"}
	fakeLines := [][]string{}
	prev := map[string]string{}
	prev["YYZ"] = "JFK"
	fakeLines = append(fakeLines, line)
	fakeLines = append(fakeLines, line)
	fakeFileReader.On("ReadFromFile", config.CSV_PATH+"/routes.csv").Return(fakeLines, nil)
	routeService := &services.RouteService{FileReader: fakeFileReader}
	_, err := routeService.GetRoutes()
	path := routeService.BuildPath(prev, "YYZ", "JFK", "")
	fmt.Println(path)
	assert.Equal(t, path, "YYZ -> JFK")

	assert.Nil(t, err)
}
