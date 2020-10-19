package main

import (
	"fmt"
)

type UnsharedConcreteFlyweight struct {
	info string
}

func (unshared *UnsharedConcreteFlyweight) SetInfo(info string) {
	unshared.info = info
}
func (unshared UnsharedConcreteFlyweight) Info() string {
	return unshared.info
}

type Flyweight interface {
	Operation(unshared UnsharedConcreteFlyweight)
}
type ConcreteFlyweight struct {
	Key string
}

func (concrete ConcreteFlyweight) Operation(unshared UnsharedConcreteFlyweight) {
	fmt.Println("concrete flyweight " + concrete.Key + " get!")
	fmt.Println("unshared info:" + unshared.Info())
}

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func (f FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := f.flyweights[key]; !ok {
		fmt.Println("create a new flyweight:", key)
		flyweight = ConcreteFlyweight{Key: key}
		f.flyweights[key] = flyweight
		return flyweight
	} else {
		fmt.Println("flyweight " + key + " has existed!")
		return flyweight
	}

}

func main() {
	factory := FlyweightFactory{flyweights: map[string]Flyweight{}}
	f1 := factory.GetFlyweight("A")
	f2 := factory.GetFlyweight("A")
	f3 := factory.GetFlyweight("A")
	f4 := factory.GetFlyweight("B")
	f5 := factory.GetFlyweight("B")

	f1.Operation(UnsharedConcreteFlyweight{info: "first allocate A"})
	f2.Operation(UnsharedConcreteFlyweight{info: "second allocate A"})
	f3.Operation(UnsharedConcreteFlyweight{info: "third allocate A"})
	f4.Operation(UnsharedConcreteFlyweight{info: "first allocate B"})
	f5.Operation(UnsharedConcreteFlyweight{info: "second allocate B"})

}
