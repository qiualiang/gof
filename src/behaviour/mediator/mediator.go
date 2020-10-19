package main

import (
	"fmt"
)

type Mediator interface {
	Register(cl Colleague)
	Reply(cl Colleague)
}
type ConcreteMediator struct {
	colleagues []Colleague
}

func (mediator *ConcreteMediator) Register(cl Colleague) {
	if !mediator.Exist(cl) {
		mediator.colleagues = append(mediator.colleagues, cl)
		cl.SetMediator(mediator)
	}
}
func (mediator *ConcreteMediator) Reply(cl Colleague) {
	for _, v := range mediator.colleagues {
		if v != cl {
			v.Receive()
		}
	}
}
func (mediator *ConcreteMediator) Exist(cl Colleague) bool {
	for _, v := range mediator.colleagues {
		if v == cl {
			return true
		}
	}
	return false
}

type Colleague interface {
	SetMediator(m Mediator)
	Receive()
	Send()
}
type ConcreteColleague1 struct {
	mediator Mediator
}

func (cl *ConcreteColleague1) SetMediator(m Mediator) {
	cl.mediator = m
}
func (cl *ConcreteColleague1) Receive() {
	fmt.Println("concrete colleague 1 receive message")
}
func (cl *ConcreteColleague1) Send() {
	fmt.Println("concrete colleague 1 sendmessage")
	cl.mediator.Reply(cl)
}

type ConcreteColleague2 struct {
	mediator Mediator
}

func (cl *ConcreteColleague2) SetMediator(m Mediator) {
	cl.mediator = m
}
func (cl *ConcreteColleague2) Receive() {
	fmt.Println("concrete colleague 2 receive message")
}
func (cl *ConcreteColleague2) Send() {
	fmt.Println("concrete colleague 2 sendmessage")
	cl.mediator.Reply(cl)
}

func main() {
	var md Mediator = &ConcreteMediator{colleagues: make([]Colleague, 0)}
	var c1, c2 Colleague
	c1 = new(ConcreteColleague1)
	c2 = new(ConcreteColleague2)
	md.Register(c1)
	md.Register(c2)
	c1.Send()
	c2.Send()
}
