package main

import (
	"fmt"
)

type State interface {
	Handle(context *Context)
}
type ConcreteStateA struct {
}

func (state *ConcreteStateA) Handle(context *Context) {
	fmt.Println("current state is A")
	context.SetState(new(ConcreteStateA))
}

type ConcreteStateB struct {
}

func (state *ConcreteStateB) Handle(context *Context) {
	fmt.Println("current state is B")
	context.SetState(new(ConcreteStateB))

}

type Context struct {
	state State
}

func (context *Context) Init() {
	context.state = new(ConcreteStateA)
}
func (context *Context) SetState(state State) {
	context.state = state
}
func (context *Context) GetState() State {
	return context.state
}
func (context *Context) Handle() {
	context.state.Handle(context)
}
func main() {
	var context *Context = new(Context)
	context.Init()
	context.Handle()
	context.SetState(new(ConcreteStateB))
	context.Handle()
	context.Init()
	context.Handle()
}
