package main

import (
	"flag"
	"fmt"
	"os"
)

//抽象产品
type Product interface {
	Show()
}

//具体产品1
type ConcreteProduct1 struct {
	name string
}

func (p1 ConcreteProduct1) Show() {
	fmt.Println("show product1:", p1.name)
}

//具体产品2
type ConcreteProduct2 struct {
	name string
}

func (p2 ConcreteProduct2) Show() {
	fmt.Println("show product2:", p2.name)
}

//抽象工厂：提供产品的生产方法
type AbstractFactory interface {
	NewProduct() Product
}

//具体工厂1：实现产品1的生成方法
type ConcreteFactory1 struct {
	name string
}

func (fc ConcreteFactory1) NewProduct() Product {
	fmt.Println("factory1 produces product1")
	return ConcreteProduct1{name: "product1"}
}

//具体工厂2：实现产品2的生成方法
type ConcreteFactory2 struct {
	name string
}

func (fc ConcreteFactory2) NewProduct() Product {
	fmt.Println("factory2 produces product2")
	return ConcreteProduct2{name: "product2"}
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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: input the concrete factory name to produce product.\n")
		fmt.Fprintf(os.Stderr, "Product 1:%s\n", "ConcreteFactory1")
		fmt.Fprintf(os.Stderr, "Product 2:%s\n", "ConcreteFactory2")
		flag.PrintDefaults()
	}
	flag.StringVar(&factory, "factory", "ConcreteFactory1", "The concrete factory name.")
}
func main() {
	// flag.Usage()
	flag.Parse()
	var p Product
	var fc AbstractFactory
	if fc = NewFactory(factory); fc != nil {
		p = fc.NewProduct()
		p.Show()
	}

}
