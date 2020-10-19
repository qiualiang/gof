package main

import (
	"fmt"
)

type Visitor interface {
	VisitA(a *ConcreteElementA)
	VisitB(b *ConcreteElementB)
}
type VisitorA struct {
}
type VisitorB struct {
}
type Element interface {
	Accept(visitor Visitor)
}

func (visotor VisitorA) VisitA(a *ConcreteElementA) {
	fmt.Println("Visitor A visit element a")
	a.OperationA()
}
func (visotor VisitorA) VisitB(b *ConcreteElementB) {
	fmt.Println("Visitor A visit element b")
	b.OperationB()
}

func (visotor VisitorB) VisitA(a *ConcreteElementA) {
	fmt.Println("Visitor B visit element a")
	a.OperationA()
}
func (visotor VisitorB) VisitB(b *ConcreteElementB) {
	fmt.Println("Visitor B visit element b")
	b.OperationB()
}

type ConcreteElementA struct {
}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitA(e)
}
func (e *ConcreteElementA) OperationA() {
	fmt.Println(" operation on A")
}

type ConcreteElementB struct {
}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitB(e)
}
func (e *ConcreteElementB) OperationB() {
	fmt.Println(" operation on B")
}

type ObjectStructure struct {
	elements []Element
}

func (o *ObjectStructure) Accept(visitor Visitor) {
	for _, v := range o.elements {
		v.Accept(visitor)
	}
}
func (o *ObjectStructure) Add(e Element) {
	if o.elements == nil {
		o.elements = []Element{}
	}
	o.elements = append(o.elements, e)
}
func (o *ObjectStructure) Remove(e Element) {
	if o.elements == nil {
		return
	}
	for i, v := range o.elements {
		if v == e {
			o.elements = append(o.elements[0:i], o.elements[i+1:]...)
		}
	}
}
func main() {
	obj := new(ObjectStructure)
	obj.Add(new(ConcreteElementA))
	obj.Add(new(ConcreteElementB))
	var visitor Visitor = VisitorA{}
	obj.Accept(visitor)
	fmt.Println("===========================")
	visitor = VisitorB{}
	obj.Accept(visitor)

}
