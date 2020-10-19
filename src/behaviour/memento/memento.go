package main

import (
	"fmt"
)

type Originator struct {
	state string
}
type Memento struct {
	state string
}

func (mem *Memento) SetState(state string) {
	mem.state = state

}
func (mem Memento) GetState() string {
	return mem.state
}
func (originator *Originator) SetState(state string) {
	originator.state = state

}
func (originator Originator) GetState() string {
	return originator.state
}
func (originator *Originator) CreateMemento() Memento {
	return Memento{state: originator.state}
}
func (originator *Originator) RestoreMemento(mem Memento) {
	originator.SetState(mem.GetState())
}

// The manager
type Caretaker struct {
	memento Memento
}

func (caretaker *Caretaker) SetMemento(mem Memento) {
	caretaker.memento = mem
}
func (caretaker Caretaker) GetMemento() Memento {
	return caretaker.memento
}
func main() {
	or := new(Originator)
	cr := new(Caretaker)
	or.SetState("S0")
	fmt.Println("The original state:", or.GetState())
	cr.SetMemento(or.CreateMemento()) // save state
	or.SetState("s1")
	fmt.Println("The new state:", or.GetState())
	or.RestoreMemento(cr.GetMemento()) // restore original state
	fmt.Println("After restore,the state is:", or.GetState())
}
