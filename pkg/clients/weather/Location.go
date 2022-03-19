package weather

import "strconv"

type CityLocation struct {
	cityName string
}

func (c *CityLocation) GETParameter() string {
	return "q=" + c.cityName
}

type CoordinateLocation struct {
	x float64
	y float64
}

func (c *CoordinateLocation) GETParameter() string {
	return "lat=" + strconv.FormatFloat(c.x, 'f', -1, 64) + "&lon=" + strconv.FormatFloat(c.y, 'f', -1, 64)
}
