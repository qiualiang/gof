package main

import (
	"fmt"
)

type Strategy interface {
	StrategyMethod()
}
type ConcreteStrategyA struct {
}

func (s ConcreteStrategyA) StrategyMethod() {
	fmt.Println("concrte strategy A's method is visited.")
}

type ConcreteStrategyB struct {
}

func (s ConcreteStrategyB) StrategyMethod() {
	fmt.Println("concrte strategy B's method is visited.")
}

type Context struct {
	strategy Strategy
}

func (cx *Context) SetSrategy(strategy Strategy) {
	cx.strategy = strategy
}
func (cx *Context) GetSrategy() Strategy {
	return cx.strategy
}
func (cx *Context) StrategyMethod() {
	cx.strategy.StrategyMethod()
}

func main() {
	context := new(Context)
	var s Strategy = new(ConcreteStrategyA)
	context.SetSrategy(s)
	context.StrategyMethod()
	fmt.Println("================================")
	s = new(ConcreteStrategyB)
	context.SetSrategy(s)
	context.StrategyMethod()
}
