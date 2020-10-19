package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	// SetData(data []string)
	Interpret(info string) bool
}
type TerminalExpression struct {
	nodes map[string]string
}

func (te *TerminalExpression) SetData(data []string) {
	te.nodes = make(map[string]string)
	for _, v := range data {
		te.nodes[v] = v
	}
}
func (te *TerminalExpression) Interpret(ifo string) bool {
	if _, ok := te.nodes[ifo]; ok {
		return true
	}
	return false
}

type AndExpression struct {
	city   Expression
	people Expression
}

func (and *AndExpression) Interpret(info string) bool {
	ss := strings.Split(info, "的")
	return and.city.Interpret(ss[0]) && and.people.Interpret(ss[1])
}

type Context struct {
	cities     []string
	people     []string
	cityPerson Expression
}

func (cx *Context) Init() {
	cx.cities = []string{"广州", "深圳"}
	cx.people = []string{"老人", "妇女", "儿童"}
	// var city, people Expression
	city := new(TerminalExpression)
	people := new(TerminalExpression)
	city.SetData(cx.cities)
	people.SetData(cx.people)
	cx.cityPerson = &AndExpression{city, people}
}
func (cx Context) FreeRide(info string) {
	if ok := cx.cityPerson.Interpret(info); ok {
		fmt.Println("您是" + info + ",您本次乘车免费")
	} else {
		fmt.Println(info + ",乘车扣费2元.")
	}

}
func main() {
	bus := new(Context)
	bus.Init()
	bus.FreeRide("广州的老人")
	bus.FreeRide("广州的妇女")
	bus.FreeRide("广州的年轻人")
	bus.FreeRide("广州的儿童")
	bus.FreeRide("深圳的老人")
	bus.FreeRide("河北的老人")
}
