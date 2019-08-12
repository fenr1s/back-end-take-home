package controllers_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fenr1s/back-end-take-home/api/controllers"

	"github.com/fenr1s/back-end-take-home/domain/models"
	"github.com/fenr1s/back-end-take-home/interfaces/mocks"

	"github.com/gin-gonic/gin"

	"github.com/fenr1s/back-end-take-home/api"
	"github.com/stretchr/testify/assert"
)

var g *gin.Engine

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	var req *http.Request
	if body != nil {
		req, _ = http.NewRequest(method, path, bytes.NewBuffer(body))
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestFlightController(t *testing.T) {
	t.Run("ShouldSearchOk", shouldSearchOk)
	t.Run("ShouldValidateOrigin", shouldValidateOrigin)
	t.Run("ShouldValidateDestination", shouldValidateDestination)
	t.Run("ShouldReturnNoRoute", shouldReturnNoRoute)
}

func shouldSearchOk(t *testing.T) {
	fakeList := []*models.Airport{}
	fakeAirport := &models.Airport{Name: "test", City: "test", Country: "test", Iata3: "tes", Latitude: 123.21, Longitude: 123.1}
	fakeList = append(fakeList, fakeAirport)
	fakeRoutes := []*models.Route{}
	fakeRoute := &models.Route{}
	fakeRoutes = append(fakeRoutes, fakeRoute)
	fakeAirportService := &mocks.Airporter{}
	fakeAirportService.On("GetAirports").Return(fakeList, nil)
	fakeAirportService.On("CheckExistance", fakeList, "test").Return(true, nil)
	fakeRouteService := &mocks.Router{}
	fakeRouteService.On("GetRoutes").Return(fakeRoutes, nil)
	fakeRouteService.On("FindShortestPath", "test", "test", fakeRoutes).Return("YYZ -> JFK")
	flightController := &controllers.FlightController{AirportService: fakeAirportService, RouteService: fakeRouteService}
	s := &api.Server{
		FlightController: flightController,
	}
	g = s.SetupRoutes()
	w := performRequest(g, "GET", "/api/routes?origin=test&destination=test", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "\"YYZ -\\u003e JFK\"", w.Body.String())
}

func shouldValidateOrigin(t *testing.T) {
	fakeList := []*models.Airport{}
	fakeAirport := &models.Airport{Name: "test", City: "test", Country: "test", Iata3: "tes", Latitude: 123.21, Longitude: 123.1}
	fakeList = append(fakeList, fakeAirport)
	fakeRoutes := []*models.Route{}
	fakeRoute := &models.Route{}
	fakeRoutes = append(fakeRoutes, fakeRoute)
	fakeAirportService := &mocks.Airporter{}
	fakeAirportService.On("GetAirports").Return(fakeList, nil)
	fakeAirportService.On("CheckExistance", fakeList, "test").Return(false, nil)
	fakeRouteService := &mocks.Router{}
	fakeRouteService.On("GetRoutes").Return(fakeRoutes, nil)
	fakeRouteService.On("FindShortestPath", "test", "test", fakeRoutes).Return("Invalid Origin")
	flightController := &controllers.FlightController{AirportService: fakeAirportService, RouteService: fakeRouteService}
	s := &api.Server{
		FlightController: flightController,
	}
	g = s.SetupRoutes()
	w := performRequest(g, "GET", "/api/routes?origin=test&destination=test", nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "\"Invalid Origin\"", w.Body.String())
}

func shouldValidateDestination(t *testing.T) {
	fakeList := []*models.Airport{}
	fakeAirport := &models.Airport{Name: "test", City: "test", Country: "test", Iata3: "tes", Latitude: 123.21, Longitude: 123.1}
	fakeList = append(fakeList, fakeAirport)
	fakeRoutes := []*models.Route{}
	fakeRoute := &models.Route{}
	fakeRoutes = append(fakeRoutes, fakeRoute)
	fakeAirportService := &mocks.Airporter{}
	fakeAirportService.On("GetAirports").Return(fakeList, nil)
	fakeAirportService.On("CheckExistance", fakeList, "test").Return(true, nil).Once()
	fakeAirportService.On("CheckExistance", fakeList, "test").Return(false, nil).Once()
	fakeRouteService := &mocks.Router{}
	fakeRouteService.On("GetRoutes").Return(fakeRoutes, nil)
	fakeRouteService.On("FindShortestPath", "test", "test", fakeRoutes).Return("Invalid Destination")
	flightController := &controllers.FlightController{AirportService: fakeAirportService, RouteService: fakeRouteService}
	s := &api.Server{
		FlightController: flightController,
	}
	g = s.SetupRoutes()
	w := performRequest(g, "GET", "/api/routes?origin=test&destination=test", nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "\"Invalid Destination\"", w.Body.String())
}

func shouldReturnNoRoute(t *testing.T) {
	fakeList := []*models.Airport{}
	fakeAirport := &models.Airport{Name: "test", City: "test", Country: "test", Iata3: "tes", Latitude: 123.21, Longitude: 123.1}
	fakeList = append(fakeList, fakeAirport)
	fakeRoutes := []*models.Route{}
	fakeRoute := &models.Route{}
	fakeRoutes = append(fakeRoutes, fakeRoute)
	fakeAirportService := &mocks.Airporter{}
	fakeAirportService.On("GetAirports").Return(fakeList, nil)
	fakeAirportService.On("CheckExistance", fakeList, "test").Return(true, nil).Twice()
	fakeRouteService := &mocks.Router{}
	fakeRouteService.On("GetRoutes").Return(fakeRoutes, nil)
	fakeRouteService.On("FindShortestPath", "test", "test", fakeRoutes).Return("")
	flightController := &controllers.FlightController{AirportService: fakeAirportService, RouteService: fakeRouteService}
	s := &api.Server{
		FlightController: flightController,
	}
	g = s.SetupRoutes()
	w := performRequest(g, "GET", "/api/routes?origin=test&destination=test", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "\"No Route\"", w.Body.String())
}

func shouldReturnInterServerError(t *testing.T) {
	fakeList := []*models.Airport{}
	fakeAirport := &models.Airport{Name: "test", City: "test", Country: "test", Iata3: "tes", Latitude: 123.21, Longitude: 123.1}
	fakeList = append(fakeList, fakeAirport)
	fakeRoutes := []*models.Route{}
	fakeRoute := &models.Route{}
	fakeRoutes = append(fakeRoutes, fakeRoute)
	fakeAirportService := &mocks.Airporter{}
	fakeAirportService.On("GetAirports").Return(fakeList, nil)
	fakeAirportService.On("CheckExistance", fakeList, "test").Return(true, nil).Twice()
	fakeRouteService := &mocks.Router{}
	fakeRouteService.On("GetRoutes").Return(fakeRoutes, errors.New("test error"))
	fakeRouteService.On("FindShortestPath", "test", "test", fakeRoutes).Return("")
	flightController := &controllers.FlightController{AirportService: fakeAirportService, RouteService: fakeRouteService}
	s := &api.Server{
		FlightController: flightController,
	}
	g = s.SetupRoutes()
	w := performRequest(g, "GET", "/api/routes?origin=test&destination=test", nil)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "\"Unexpected Error Ocurred\"", w.Body.String())
}
