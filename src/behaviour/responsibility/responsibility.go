package main

import (
	"fmt"
)

type Handle interface {
	SetNext(next Handle)
	GetNext() Handle
	HandleRequest(request string)
}
type ConcreteHandle1 struct {
	next Handle
}

func (handle *ConcreteHandle1) SetNext(next Handle) {
	handle.next = next
}
func (handle *ConcreteHandle1) GetNext() Handle {
	return handle.next
}
func (handle *ConcreteHandle1) HandleRequest(request string) {
	if request == "one" {
		fmt.Println("concrete handler 1 response for the request")
	} else {
		if handle.GetNext() != nil {
			handle.GetNext().HandleRequest(request)
		} else {
			fmt.Println("No one handle the request.")
		}
	}
}

type ConcreteHandle2 struct {
	next Handle
}

func (handle *ConcreteHandle2) SetNext(next Handle) {
	handle.next = next
}
func (handle *ConcreteHandle2) GetNext() Handle {
	return handle.next
}
func (handle *ConcreteHandle2) HandleRequest(request string) {
	if request == "two" {
		fmt.Println("concrete handler 2 response for the request")
	} else {
		if handle.GetNext() != nil {
			handle.GetNext().HandleRequest(request)
		} else {
			fmt.Println("No one handle the request.")
		}
	}
}
func main() {
	var handle1, handle2 Handle
	handle1 = new(ConcreteHandle1)
	handle2 = new(ConcreteHandle2)
	handle1.SetNext(handle2)
	handle1.HandleRequest("one")
	handle1.HandleRequest("two")
	handle1.HandleRequest("three")

}
