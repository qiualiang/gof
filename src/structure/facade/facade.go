package main

import "fmt"

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

type Square struct {
	width float64
}

func (s Square) Draw() {
	fmt.Printf("Draw a square with width:%f\n", s.width)
}

type ShapeMaker struct {
	circle    Circle
	rectangle Rectangle
	square    Square
}

func (maker ShapeMaker) DrawCircle() {
	maker.circle.Draw()
}
func (maker ShapeMaker) DrawRectangle() {
	maker.rectangle.Draw()
}
func (maker ShapeMaker) DrawSquare() {
	maker.square.Draw()
}

func main() {
	c := Circle{radius: 5}
	r := Rectangle{width: 10, length: 5}
	s := Square{width: 10}
	maker := ShapeMaker{circle: c, rectangle: r, square: s}
	maker.DrawCircle()
	maker.DrawRectangle()
	maker.DrawSquare()
}
