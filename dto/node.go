package dto

type Node struct {
	Sw       Point
	Ne       Point
	Children []Node
	Text     string
}
