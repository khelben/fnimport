package db

import (
	"github.com/khelben/gormGIS"
)

type FietsNetTrack struct {
	Id        int
	Name      string
	StartNode int
	EndNode   int
	Point     gormGIS.GeoLineString `sql:"type:geometry(Linestring,4326)"`
}
