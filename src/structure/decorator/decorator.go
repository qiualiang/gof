package main

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type Circle struct {
	radius float64
}

func (c Circle) Draw() {
	fmt.Println("Draw a circle with radius:", c.radius)
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Draw() {
	fmt.Printf("Draw a rectangle with length:%f,width:%f\n", r.length, r.width)
}

type ShapeDecorator interface {
	Shape
}
type RedShapeDecorator struct {
	S Shape
}

func (d RedShapeDecorator) Draw() {
	d.S.Draw()
	d.setRedBorder()
}
func (d RedShapeDecorator) setRedBorder() {
	fmt.Println("Decorator:set red border!")
}

func main() {
	var s Shape = Circle{radius: 12.11}
	s.Draw()
	fmt.Println("===========================")
	fmt.Println("Draw shap with decorator:")
	var d Shape = RedShapeDecorator{S: s}
	d.Draw()
}
