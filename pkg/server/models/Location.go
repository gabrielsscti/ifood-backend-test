package models

import "strconv"

type CityLocation struct {
	CityName string `json:"name"`
}

func (c *CityLocation) GETParameter() string {
	return "q=" + c.CityName
}

type CoordinateLocation struct {
	X float64 `json:"X"`
	Y float64 `json:"Y"`
}

func (c *CoordinateLocation) GETParameter() string {
	return "lat=" + strconv.FormatFloat(c.X, 'f', -1, 64) + "&lon=" + strconv.FormatFloat(c.Y, 'f', -1, 64)
}
