package model

type SubAddress struct {
	Country   string `json:"country"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       int `json:"zip"`
	Street    string `json:"street"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}
