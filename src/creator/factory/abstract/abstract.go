package main

import (
	"flag"
	"fmt"
	"os"
)

//抽象产品
type Product1 interface {
	Show1()
}
type Product2 interface {
	Show2()
}

//具体产品1-1
type ConcreteProduct11 struct {
	name string
}

func (p ConcreteProduct11) Show1() {
	fmt.Println("show product:", p.name)
}

//具体产品1-2
type ConcreteProduct12 struct {
	name string
}

func (p ConcreteProduct12) Show1() {
	fmt.Println("show product:", p.name)
}

//具体产品2-1
type ConcreteProduct21 struct {
	name string
}

func (p ConcreteProduct21) Show2() {
	fmt.Println("show product:", p.name)
}

//具体产品2-2
type ConcreteProduct22 struct {
	name string
}

func (p ConcreteProduct22) Show2() {
	fmt.Println("show product:", p.name)
}

//抽象工厂：提供产品的生产方法
type AbstractFactory interface {
	NewProduct1() Product1
	NewProduct2() Product2
}

//具体工厂1：实现产品1和2的生成方法
type ConcreteFactory1 struct {
	name string
}

func (fc ConcreteFactory1) NewProduct1() Product1 {
	fmt.Println("factory1 produces product1")
	return ConcreteProduct11{name: "product1"}
}
func (fc ConcreteFactory1) NewProduct2() Product2 {
	fmt.Println("factory1 produces product2")
	return ConcreteProduct21{name: "product2"}
}

//具体工厂2：实现产品1和2的生成方法
type ConcreteFactory2 struct {
	name string
}

func (fc ConcreteFactory2) NewProduct1() Product1 {
	fmt.Println("factory2 produces product1")
	return ConcreteProduct12{name: "product1"}
}
func (fc ConcreteFactory2) NewProduct2() Product2 {
	fmt.Println("factory2 produces product2")
	return ConcreteProduct22{name: "product2"}
}
func NewFactory(name string) AbstractFactory {
	switch name {
	case "ConcreteFactory1":
		return ConcreteFactory1{name: "factory1"}
	case "ConcreteFactory2":
		return ConcreteFactory2{name: "factory2"}
	default:
		fmt.Fprintf(os.Stderr, "Unsupported factory:%s\n", factory)
	}
	return nil
}

var factory string

func init() {
	flag.StringVar(&factory, "factory", "ConcreteFactory1", "The concrete factory name.")
	// flag.StringVar(&product, "product", "Product1", "The concrete product name of factory")
}
func main() {
	flag.Parse()
	var p1 Product1
	var p2 Product2
	var fc AbstractFactory
	if fc = NewFactory(factory); fc != nil {
		p1 = fc.NewProduct1()
		p1.Show1()
		p2 = fc.NewProduct2()
		p2.Show2()
	}
}
