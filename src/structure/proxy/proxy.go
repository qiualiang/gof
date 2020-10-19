package main

import "fmt"

type Image interface {
	Display()
}
type RealImage struct {
	FileName string
}

func (real RealImage) Display() {
	fmt.Println("Display the real image")
}
func (real *RealImage) LoadFromDisk(fileName string) {
	fmt.Println("Load image from disk.")
	real.FileName = fileName
}

type ProxyImage struct {
	Real     *RealImage
	FileName string
}

func (proxy *ProxyImage) Display() {
	fmt.Println("Display the proxy image")
	if proxy.Real == nil {
		proxy.Real = new(RealImage)
		proxy.Real.LoadFromDisk(proxy.FileName)
	}
	proxy.Real.Display()
}
func main() {
	proxy := ProxyImage{FileName: "/images/1.png"}
	proxy.Display()
	fmt.Println("=======================")
	fmt.Println("Display the second time")
	proxy.Display()
}
