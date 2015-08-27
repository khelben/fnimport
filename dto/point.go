package dto

import (
	"math"

	"github.com/khelben/gormGIS"
)

type Point struct {
	Lat float64
	Lng float64
}

func (g Point) ToGeoPoint() gormGIS.GeoPoint {
	lng, lat := g.Lng, g.Lat
	t := (lng - .5) * 2 * math.Pi
	i := math.Asin((math.Exp(2*math.Pi*(1-2*lat)) - 1) / (math.Exp(2*math.Pi*(1-2*lat)) + 1))
	return gormGIS.GeoPoint{
		Lng: t * 180 / math.Pi,
		Lat: i * 180 / math.Pi,
	}
}
