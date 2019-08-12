package controllers

import (
	"net/http"

	"github.com/fenr1s/back-end-take-home/interfaces"
	"github.com/gin-gonic/gin"
)

// @title Flight Service Swagger
// @version 1.0
// @description This is a swagger for guestlogix challenge.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email andrews8.wu@gmail.com

// @host localhost:8081
// @BasePath /api/

//FlightController controller handles flight resources
type FlightController struct {
	RouteService   interfaces.Router
	AirportService interfaces.Airporter
}

// @Summary Search
// @Description given an origin and a destination search for the shortest path
// @Produce  json
// @Param origin query string false "origin iata containing 3 character" minlength(0) maxlength(3)
// @Param destination query string false "destination iata containing 3 character" minlength(0) maxlength(3)
// @Router /api/routes?origin={origin}&destination={destination} [get]
// @Success 200 {string} string "shortest path"
// @Failure 500 {string} string ""
// @Failure 400 {string} string ""
func (r *FlightController) Search(c *gin.Context) {
	origin := c.Query("origin")
	destination := c.Query("destination")
	airports, err := r.AirportService.GetAirports()
	exists, err := r.AirportService.CheckExistance(airports, origin)
	if !exists {
		c.JSON(http.StatusBadRequest, "Invalid Origin")
		return
	}
	exists, err = r.AirportService.CheckExistance(airports, destination)
	if !exists {
		c.JSON(http.StatusBadRequest, "Invalid Destination")
		return
	}
	routes, err := r.RouteService.GetRoutes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Unexpected Error Ocurred")
		return
	}
	path := r.RouteService.FindShortestPath(origin, destination, routes)
	if path == "" {
		c.JSON(http.StatusOK, "No Route")
		return
	}

	c.JSON(http.StatusOK, path)
}
