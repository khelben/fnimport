package dto

type Poi struct {
	Id         int
	GraphId    int
	Coordinate Point
	Label      string
	Category   string
}
