package main

import (
	"github.com/fenr1s/back-end-take-home/api"
	"github.com/fenr1s/back-end-take-home/api/controllers"
	"github.com/fenr1s/back-end-take-home/domain/services"
)

func main() {
	//Inversion of controll (di)
	fileReader := &services.FileReader{}
	routeService := &services.RouteService{FileReader: fileReader}
	airportService := &services.AirportService{FileReader: fileReader}
	flightController := &controllers.FlightController{RouteService: routeService, AirportService: airportService}
	server := &api.Server{FlightController: flightController}

	server.Run()
}
