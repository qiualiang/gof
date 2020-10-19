package main

import "fmt"

type IPerson interface {
	SetName(name string)
	BeforeOut()
	Out()
}
type Person struct {
	Specific IPerson
	name     string
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Out() {
	p.BeforeOut()
	fmt.Println(p.name + " go out...")
}

func (p *Person) BeforeOut() {
	if p.Specific == nil {
		return
	}

	p.Specific.BeforeOut()
}

type Boy struct {
	Person
}

func (_ *Boy) BeforeOut() {
	fmt.Println("get up..")
}

type Girl struct {
	Person
}

func (_ *Girl) BeforeOut() {
	fmt.Println("dress up..")
}
func main() {
	var p *Person = &Person{}

	p.Specific = &Boy{}
	p.SetName("Tony")
	p.Out()

	p.Specific = &Girl{}
	p.SetName("Lily")
	p.Out()
}
