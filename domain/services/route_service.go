package services

import (
	"github.com/fenr1s/back-end-take-home/config"
	"github.com/fenr1s/back-end-take-home/domain/models"
	"github.com/fenr1s/back-end-take-home/interfaces"
)

//RouteService contains logic of the route context
type RouteService struct {
	FileReader interfaces.FileReader
}

//GetRoutes return a array of routes from the csv and a possible error
func (r *RouteService) GetRoutes() ([]*models.Route, error) {
	routes := []*models.Route{}
	lines, err := r.FileReader.ReadFromFile(config.CSV_PATH + "/routes.csv")
	if err != nil {
		return nil, err
	}

	for i, line := range lines {
		if i != 0 {
			routes = append(routes, &models.Route{AirlineID: line[0], Origin: line[1], Destination: line[2]})
		}
	}
	return routes, err
}

//FindShortestPath Implements the BDS(breadth-first-search algorithm) to find the shortest path
//it requires an origin a destination and returns a string representing the shortest path
func (r *RouteService) FindShortestPath(origin string, destination string, routes []*models.Route) string {
	unvisiteds := []string{}
	unvisiteds = append(unvisiteds, origin)
	visiteds := []string{}
	prev := map[string]string{}

	for len(unvisiteds) != 0 {
		currentNode := unvisiteds[0]
		unvisiteds = unvisiteds[1:]
		visiteds = append(visiteds, currentNode)

		filteredRoutes, _ := r.FindRoutesByOrigin(currentNode, routes)
		for _, route := range filteredRoutes {
			if route.Destination == destination {
				prev[route.Origin] = route.Destination
				path := r.BuildPath(prev, origin, destination, "")
				return path
			}

			alreadyVisited := false
			for _, visited := range visiteds {
				if route.Destination == visited {
					alreadyVisited = true
					break
				}
			}

			if !alreadyVisited {
				prev[route.Origin] = route.Destination
				unvisiteds = append(unvisiteds, route.Destination)
			}
		}
	}
	return ""
}

//FindRoutesByOrigin filter all routes by a given origin
//returns a array of routes and a possible error
func (r *RouteService) FindRoutesByOrigin(origin string, routes []*models.Route) (routesByOrigin []*models.Route, err error) {
	for _, route := range routes {
		if route.Origin == origin {
			routesByOrigin = append(routesByOrigin, route)
		}
	}

	return routesByOrigin, err
}

//BuildPath a function to build a string representation of the path
//it requires a map of paths a origin and a destination, it uses recursion to build the path string and returns it
func (r *RouteService) BuildPath(prev map[string]string, origin string, destination string, path string) string {
	if origin == destination {
		return path
	}

	if prev[origin] == "" {
		return path + " -> " + destination
	}

	if path != "" {
		path = path + " -> " + prev[origin]
	} else {
		path = origin + " -> " + prev[origin]
	}

	return r.BuildPath(prev, prev[origin], destination, path)
}
