package main

import (
	"fmt"
)

type Target interface {
	Request()
}
type Adaptee interface {
	SepcificRequest()
}
type AdapteeImpl struct {
	name string
}

func (adaptee AdapteeImpl) SepcificRequest() {
	fmt.Printf("适配者%s中的业务代码被调用！\n", adaptee.name)
}

type Adapter struct {
	name    string
	adaptee Adaptee
}

func (adapter *Adapter) New() {
	adapter.adaptee = AdapteeImpl{name: "adaptee"}
}
func (adapter Adapter) Request() {
	fmt.Printf("适配器%s中的业务代码被调用！\n", adapter.name)
	adapter.adaptee.SepcificRequest()
}

func main() {
	fmt.Println("适配器模式测试：")
	var adapter Adapter = Adapter{name: "adapter"}
	adapter.New()
	var target Target = adapter
	target.Request()
}
