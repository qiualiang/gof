package main

import (
	"fmt"
)

type Iterator interface {
	First() interface{}
	Next() interface{}
	HasNext() bool
}
type Aggregate interface {
	Add(obj interface{})
	Remove(obj interface{})
	GetIterator() Iterator
}
type ConcreteIterator struct {
	objs  []interface{}
	index int
}

func (it *ConcreteIterator) Init(objs []interface{}) {
	it.index = -1
	it.objs = make([]interface{}, len(objs))
	copy(it.objs, objs)
}
func (it *ConcreteIterator) First() interface{} {
	it.index = 0
	obj := it.objs[it.index]
	return obj
}
func (it *ConcreteIterator) HasNext() bool {
	if it.index < len(it.objs)-1 {
		return true
	}
	return false
}
func (it *ConcreteIterator) Next() interface{} {
	if it.HasNext() {
		it.index += 1
		return it.objs[it.index]
	}
	return nil
}

type ConcreteAggregate struct {
	objs []interface{}
}

func (a *ConcreteAggregate) Add(obj interface{}) {
	if a.objs == nil {
		a.objs = make([]interface{}, 0)
	}
	a.objs = append(a.objs, obj)
}
func (a *ConcreteAggregate) Remove(obj interface{}) {
	if a.objs == nil {
		return
	}
	for i, v := range a.objs {
		if v == obj {
			a.objs = append(a.objs[0:i], a.objs[i+1:]...)
		}
	}
}

func (a *ConcreteAggregate) GetIterator() Iterator {
	it := new(ConcreteIterator)
	it.Init(a.objs)
	return it
}
func main() {
	var ag Aggregate = new(ConcreteAggregate)
	ag.Add("first")
	ag.Add("second")
	ag.Add("third")
	var it Iterator = ag.GetIterator()
	for it.HasNext() {
		obj := it.Next()
		fmt.Println(obj.(string))
	}
	obj := it.First()
	fmt.Println("The firs one:", obj.(string))
}
