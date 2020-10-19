package main

import (
	"fmt"
)

type Subject interface {
	Add(ob Observer)
	Remove(ob Observer)
	NotifyObserver()
}
type Observer interface {
	Response()
}
type ConcreteSubject struct {
	observers []Observer
}

func (cs *ConcreteSubject) Add(ob Observer) {
	if cs.observers == nil {
		cs.observers = make([]Observer, 0)
	}
	cs.observers = append(cs.observers, ob)
}
func (cs *ConcreteSubject) Remove(ob Observer) {
	for i, v := range cs.observers {
		if ob == v {
			cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
		}
	}
}
func (cs *ConcreteSubject) NotifyObserver() {
	fmt.Println("Subject change,notify the observers!")
	for _, v := range cs.observers {
		v.Response()
	}
}

type ConcreteObserver1 struct {
}

func (ob *ConcreteObserver1) Response() {
	fmt.Println("concrete observer 1 response.")

}

type ConcreteObserver2 struct {
}

func (ob *ConcreteObserver2) Response() {
	fmt.Println("concrete observer 2 response.")

}
func main() {
	var subject Subject = new(ConcreteSubject)
	obs1 := new(ConcreteObserver1)
	obs2 := new(ConcreteObserver2)
	subject.Add(obs1)
	subject.Add(obs2)
	// subject.Remove(obs2)
	subject.NotifyObserver()
}
