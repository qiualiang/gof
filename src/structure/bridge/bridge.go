package main

import (
	"fmt"
)

type Color interface {
	Color() string
}
type Yellow struct {
}

func (color Yellow) Color() string {
	return "yellow"
}

type Red struct {
}

func (color Red) Color() string {
	return "red"
}

type Kind interface {
	Kind() string
}
type Wallet struct {
}

func (wallet Wallet) Kind() string {
	return "wallet"
}

type HandBag struct {
}

func (handBag HandBag) Kind() string {
	return "hand bag"
}

type Bag interface {
	SetColor(color Color)
	SetKind(kind Kind)
	GetName()
}
type ChooseBag struct {
	color Color
	kind  Kind
}

func (bag *ChooseBag) SetColor(color Color) {
	bag.color = color
}
func (bag *ChooseBag) SetKind(kind Kind) {
	bag.kind = kind
}
func (bag *ChooseBag) GetName() {
	fmt.Println("color:", bag.color.Color(), "kind:", bag.kind.Kind())
}
func main() {
	var color Color = Yellow{}
	var kind Kind = HandBag{}
	var bag Bag
	bag = new(ChooseBag)
	bag.SetColor(color)
	bag.SetKind(kind)
	bag.GetName()
}
