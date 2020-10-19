package main

import "fmt"

type Motor interface {
	Drive()
}

type ElectricMotor struct {
}

func (em ElectricMotor) ElectricDrive() {
	fmt.Println("电能发动机驱动汽车！")
}

type OpticalMotor struct {
}

func (om OpticalMotor) opticalDrive() {
	fmt.Println("光能发动机驱动汽车！")
}

type ElectricAdapter struct {
	emotor ElectricMotor
}

func (ea ElectricAdapter) New() {
	ea.emotor = ElectricMotor{}
}
func (ea ElectricAdapter) Drive() {
	ea.emotor.ElectricDrive()
}

type OpticalAdapter struct {
	omotor OpticalMotor
}

func (op OpticalAdapter) New() {
	op.omotor = OpticalMotor{}
}
func (op OpticalAdapter) Drive() {
	op.omotor.opticalDrive()
}
func main() {
	fmt.Println("使用电能适配器")
	emotor := ElectricAdapter{}
	emotor.New()
	var motor Motor = emotor
	motor.Drive()

	fmt.Println("使用光能适配器")
	omotor := OpticalAdapter{}
	omotor.New()
	motor = omotor
	motor.Drive()
}
