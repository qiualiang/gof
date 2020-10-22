package main

import "fmt"

type AbstractClass interface {
	TemplateMethod()
	SpecificMethod()
	AbstractMethod1()
	AbstractMethod2()
}
type ConcreteClass struct {
	Specific AbstractClass
}

func (cc *ConcreteClass) AbstractMethod1() {
	if cc.Specific == nil {
		return
	}
	cc.Specific.AbstractMethod1()
}
func (cc *ConcreteClass) AbstractMethod2() {
	if cc.Specific == nil {
		return
	}
	cc.Specific.AbstractMethod2()

}
func (cc *ConcreteClass) SpecificMethod() {
	fmt.Println("running specificMethod")
}
func (cc *ConcreteClass) TemplateMethod() {
	cc.SpecificMethod()
	cc.AbstractMethod1()
	cc.AbstractMethod2()
}

type ConcreteA struct {
	ConcreteClass
}

func (_ *ConcreteA) AbstractMethod1() {
	fmt.Println("running abstract method 1 of concretA..")
}
func (_ *ConcreteA) AbstractMethod2() {
	fmt.Println("running abstract method 2 of concretA..")
}

type ConcreteB struct {
	ConcreteClass
}

func (_ *ConcreteB) AbstractMethod1() {
	fmt.Println("running abstract method 1 of concretB..")
}
func (_ *ConcreteB) AbstractMethod2() {
	fmt.Println("running abstract method 2 of concretB..")
}

func main() {
	var p *ConcreteClass = &ConcreteClass{}
	fmt.Println("the object is A")
	p.Specific = &ConcreteA{}
	p.TemplateMethod()
	fmt.Println("=========================")
	fmt.Println("the object is B")
	p.Specific = &ConcreteB{}
	p.TemplateMethod()
}
