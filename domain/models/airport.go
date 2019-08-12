package models

//Airport model representation
type Airport struct {
	Name      string  `json:"name"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Iata3     string  `json:"iata3"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
