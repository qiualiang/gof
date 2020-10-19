package main

import (
	"fmt"
)

type Stock struct {
	name     string
	quantity int
}

func (stock *Stock) Set(name string, quantity int) {
	stock.name = name
	stock.quantity = quantity
}

func (stock *Stock) Buy(s Stock) {
	stock.name = s.name
	stock.quantity = s.quantity
	fmt.Printf("Buy stock %s, quantity:%d \n", s.name, s.quantity)
}
func (stock Stock) Sell(s Stock) {
	stock.name = s.name
	stock.quantity = s.quantity
	fmt.Printf("Sell stock %s, quantity:%d \n", s.name, s.quantity)
}

type Order interface {
	Execute()
}
type BuyStock struct {
	stock Stock
}

func (buy BuyStock) Execute() {
	fmt.Println("execute in buy command")
	buy.stock.Buy(buy.stock)
}

type SellStock struct {
	stock Stock
}

func (sell SellStock) Execute() {
	fmt.Println("execute in sell command")
	sell.stock.Sell(sell.stock)
}

type Broker struct {
	order Order
}

func (broker *Broker) SetOrder(order Order) {
	fmt.Printf("set order:%v\n", order)
	broker.order = order
}
func (broker Broker) Call() {
	fmt.Println("call in broker")
	broker.order.Execute()
}

func main() {
	stock := new(Stock)
	stock.Set("icbc", 200)
	var order Order
	order = BuyStock{*stock}
	broker := new(Broker)
	broker.SetOrder(order)
	broker.Call()

	order = SellStock{*stock}
	broker.SetOrder(order)
	broker.Call()
}
