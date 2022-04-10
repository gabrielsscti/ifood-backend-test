package models

type CityTracksRequest struct {
	Service  string       `json:"service,omitempty"`
	Location CityLocation `json:"city"`
}

type CoordinateTracksRequest struct {
	Service     string             `json:"service,omitempty"`
	Coordinates CoordinateLocation `json:"coords"`
}
