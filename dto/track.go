package dto

type Track struct {
	Id          int
	Bounds      Bounds
	StartVertex int
	EndVertex   int
	Path        []Point
}
