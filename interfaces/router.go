package interfaces

import "github.com/fenr1s/back-end-take-home/domain/models"

//Router contract
type Router interface {
	GetRoutes() ([]*models.Route, error)
	FindShortestPath(origin string, destination string, routes []*models.Route) string
	FindRoutesByOrigin(origin string, routes []*models.Route) (routesByOrigin []*models.Route, err error)
	BuildPath(prev map[string]string, origin string, destination string, path string) string
}
