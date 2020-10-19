package main

import (
	"fmt"
)

type HungrySingleton struct {
	name string
}

var instance *HungrySingleton

func (singleton *HungrySingleton) String() string {
	return singleton.name
}
func (singleton *HungrySingleton) Name() string {
	return singleton.name
}
func (singleton *HungrySingleton) SetName(name string) {
	singleton.name = name
}

func (singleton *HungrySingleton) GetInstance() *HungrySingleton {
	return instance
}
func init() {
	instance = &HungrySingleton{name: "gina"}
}
func main() {
	fmt.Printf("original instance:%s\n", instance)
	instance1 := new(HungrySingleton).GetInstance()
	instance2 := new(HungrySingleton).GetInstance()

	if instance1 == instance2 {
		fmt.Println("instance1 is equal to instance2")
	} else {
		fmt.Println("instance1 isn't equal to instance2")
	}
	instance1.SetName("haha")
	fmt.Printf("after change name :instance1:%s,instance2:%s\n", instance1, instance2)
}
