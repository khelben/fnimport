package db

import (
	"github.com/khelben/gormGIS"
)

type FietsNetNode struct {
	Id    int
	Name  string
	Point gormGIS.GeoPoint `sql:"type:geometry(Geometry,4326)"`
}
