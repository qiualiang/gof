package main

import (
	"fmt"
)

type Parlour struct {
	wall string
	tv   string
	sofa string
}

func (p *Parlour) SetTV(tv string) {
	p.tv = tv
}
func (p *Parlour) SetWall(wall string) {
	p.wall = wall
}
func (p *Parlour) SetSofa(sofa string) {
	p.sofa = sofa
}
func (p *Parlour) Show() {
	fmt.Printf("wall:%s,tv:%s,sofa:%s\n", p.wall, p.tv, p.sofa)
}

type Decorator interface {
	BuildWall()
	BuildTV()
	BuildSofa()
	GetResult() *Parlour
}

type ConcreteDecorator1 struct {
	parlour *Parlour
}

func (cd1 *ConcreteDecorator1) SetParlour(parlour *Parlour) {
	cd1.parlour = parlour
	fmt.Println("set the parlour")
}
func (cd1 ConcreteDecorator1) BuildWall() {
	cd1.parlour.SetWall("build by builder1")
	fmt.Println("Wall built by cd1!")
}
func (cd1 ConcreteDecorator1) BuildTV() {
	cd1.parlour.SetTV("build by builder1")
	fmt.Println("TV built by cd1!")
}
func (cd1 ConcreteDecorator1) BuildSofa() {
	cd1.parlour.SetSofa("build by builder1")
	fmt.Println("Sofa built by cd1!")
}
func (cd1 ConcreteDecorator1) GetResult() *Parlour {
	return cd1.parlour
}

type ConcreteDecorator2 struct {
	parlour *Parlour
}

func (cd2 *ConcreteDecorator2) SetParlour(parlour *Parlour) {
	cd2.parlour = parlour
	fmt.Println("Set the parlour")
}
func (cd2 ConcreteDecorator2) BuildWall() {
	cd2.parlour.SetWall("build by builder2")
	fmt.Println("Wall built by cd2!")
}
func (cd2 ConcreteDecorator2) BuildTV() {
	cd2.parlour.SetTV("build by builder2")
	fmt.Println("TV built by cd2!")
}
func (cd2 ConcreteDecorator2) BuildSofa() {
	cd2.parlour.SetSofa("build by builder2")
	fmt.Println("Sofa built by cd2!")
}
func (cd2 ConcreteDecorator2) GetResult() *Parlour {
	return cd2.parlour
}

type ProjectManager struct {
	builder Decorator
}

func (manager *ProjectManager) SetDecorator(builder Decorator) {
	manager.builder = builder
}
func (manager *ProjectManager) Decorate() *Parlour {
	builder := manager.builder
	builder.BuildSofa()
	builder.BuildTV()
	builder.BuildWall()
	return builder.GetResult()
}
func main() {
	fmt.Println("Manager ask builder1 to decorate the parlour:")
	builder1 := ConcreteDecorator1{parlour: new(Parlour)}
	manager := new(ProjectManager)
	manager.SetDecorator(builder1)
	p := manager.Decorate()
	p.Show()
	fmt.Println("Manager ask builder2 to decorate the parlour:")
	builder2 := ConcreteDecorator2{parlour: new(Parlour)}
	// manager := new(ProjectManager)
	manager.SetDecorator(builder2)
	p2 := manager.Decorate()
	p2.Show()

}
