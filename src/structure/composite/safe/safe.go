package main

import "fmt"

type Component interface {
	Operation()
}

type Leaf struct {
	name string
}

func (l *Leaf) SetName(name string) {
	l.name = name
}

func (l *Leaf) Operation() {
	fmt.Println("Leaf " + l.name + " visited!")
}

type Composite struct {
	children []Component
}

func (coms *Composite) NewChildren() {
	coms.children = make([]Component, 0)
}
func (coms *Composite) Add(c Component) {
	coms.children = append(coms.children, c)
}
func (coms *Composite) Remove(c Component) {
	switch c.(type) {
	case *Leaf:
		// fmt.Println("c is a leaf")
		children := coms.children
		for i, v := range children {
			if value, ok := v.(*Leaf); ok {
				if value == c {
					fmt.Println("c is a leaf,found!")
					// fmt.Println("i:", i)
					switch i {
					case 0: //at the first
						// fmt.Println("case first")
						coms.children = children[1:]
					case len(children) - 1:
						// fmt.Println("case last")
						coms.children = children[0 : len(children)-1]
					default:
						// fmt.Println("case middle")
						coms.children = append(children[:i], children[i+1:]...)

					}
					break
				}
			} else {
				if value, ok := v.(*Composite); ok {
					value.Remove(c)
				}

			}
		}
	case *Composite:
		// fmt.Println("c is a composite")
		children := coms.children
		for i, v := range children {
			if value, ok := v.(*Composite); ok {
				if value == c {
					fmt.Println("c is a composite,found!")
					// fmt.Println("i:", i)
					switch i {
					case 0: //at the first
						// fmt.Println("case first")
						coms.children = children[1:]
					case len(children) - 1:
						// fmt.Println("case last")
						coms.children = children[0 : len(children)-1]
					default:
						// fmt.Println("case middle")
						coms.children = append(children[:i], children[i+1:]...)

					}
					break
				} else {
					value.Remove(c)
				}
			}
		}

	default:
		fmt.Println("c is a component")
	}

}
func (coms *Composite) GetChild(i int) Component {
	return coms.children[i]
}
func (coms *Composite) Operation() {
	for _, child := range coms.children {
		child.Operation()
	}
}
func main() {
	fmt.Println("safe pattern")
	c0 := new(Composite)
	c0.NewChildren()
	c1 := new(Composite)
	c1.NewChildren()
	c2 := new(Composite)
	c2.NewChildren()
	l1 := Leaf{name: "1"}
	l2 := Leaf{name: "2"}

	l3 := Leaf{name: "3"}
	l4 := Leaf{name: "4"}
	l5 := Leaf{name: "5"}
	l6 := Leaf{name: "6"}
	c0.Add(&l1)
	c0.Add(c1)
	c0.Add(&l5)
	c0.Add(&l6)
	c1.Add(&l2)
	c1.Add(c2)
	c2.Add(&l3)
	c2.Add(&l4)
	c0.Operation()
	fmt.Println("remove.......")
	// c0.Remove(&l1)
	// c0.Remove(&l6)
	// c0.Remove(&l5)
	// c0.Remove(&l2)
	c0.Remove(c1)
	c0.Remove(&Leaf{name: "8"})
	c3 := new(Composite)
	c3.NewChildren()
	c0.Remove(c3)
	fmt.Println("after removed:")
	c0.Operation()
}
