package models

//Airline model representation
type Airline struct {
	Name           string `json:"name"`
	TwoDigitCode   string `json:"two_digit_code"`
	ThreeDigitCode string `json:"three_digit_code"`
	Country        string `json:"country"`
}
