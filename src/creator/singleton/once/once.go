package main

import (
	"fmt"
	"sync"
)

type HungrySingleton struct {
	name string
}

var instance *HungrySingleton
var once sync.Once

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
	once.Do(func() {
		fmt.Println("Creating instance!")
		instance = &HungrySingleton{name: "Gina"}
	})
	return instance
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
	fmt.Printf("instance1:%s,instance2:%s\n", instance1, instance2)
}
