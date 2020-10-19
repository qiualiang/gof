package main

import (
	"fmt"
)

type Cloneable interface {
	Clone() *SunWuKong
}
type SunWuKong struct {
	real  bool
	kind  string
	skill []string
}

func (wukong SunWuKong) Clone() *SunWuKong {
	return &SunWuKong{
		real:  false,
		kind:  wukong.kind,
		skill: wukong.skill,
	}
}
func main() {
	wukong := SunWuKong{real: true, kind: "monkey", skill: []string{"jump", "fly", "invisible"}}
	fake := wukong.Clone()
	fmt.Printf("wukong is real:%v\n", wukong.real)
	fmt.Printf("fake is real:%v\n", fake.real)
	_, ok := interface{}(wukong).(Cloneable)
	fmt.Println(ok)
}
