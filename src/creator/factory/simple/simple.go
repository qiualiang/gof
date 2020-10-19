package main

import (
	"fmt"
)

const (
	SQUARE    = "square"
	CIRCLE    = "circle"
	RECTANGLE = "rectangle"
)

type Shape interface {
	draw()
}

type Square struct{}
type Circle struct{}
type Rectangle struct{}
type ShapeFactory struct{}

func (s Square) draw() {
	fmt.Println("Draw a square")
}

func (c Circle) draw() {
	fmt.Println("Draw a circle")
}

func (r Rectangle) draw() {
	fmt.Println("Draw a rectangle")
}
func (factory ShapeFactory) getShape(shape string) Shape {
	switch shape {
	case RECTANGLE:
		return new(Rectangle)
	case SQUARE:
		return new(Square)
	case CIRCLE:
		return new(Circle)
	default:
		fmt.Println("Unsupported shape:", shape)
		return nil
	}
}

func main() {
	var s Shape
	factory := new(ShapeFactory)
	if s = factory.getShape("circle"); s != nil {
		s.draw()
	}

	if s = factory.getShape("square"); s != nil {
		s.draw()
	}

	if s = factory.getShape("rectangle"); s != nil {
		s.draw()
	}

	if s = factory.getShape("triangle"); s != nil {
		s.draw()
	}
}
