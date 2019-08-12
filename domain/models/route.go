package models

//Route model representation
type Route struct {
	AirlineID   string `json:"airline_id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}
