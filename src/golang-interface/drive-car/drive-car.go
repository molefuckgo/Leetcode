package main

import "fmt"

func main() {
	// 张三开宝马
	bmw := &Bmw{}
	zhang3 := &Zhang3{}
	zhang3.Drive(bmw)
}

// ========>	抽象层	<========
type Car interface {
	Run()
}

type Driver interface {
	Driver(car Car)
}

// ========>	实现层	<========
type Benz struct {
}

func (benz *Benz) Run() {
	fmt.Println("Benz is running...")
}

type Bmw struct {
}

func (bmw Bmw) Run() {
	fmt.Println("Bmw is running...")
}

type Zhang3 struct {
}

func (zhang3 Zhang3) Drive(car Car) {
	fmt.Println("zhang3 drive car")
	car.Run()
}

type Li4 struct {
}

func (li4 Li4) Drive(car Car) {
	fmt.Println("li4 drive car")
	car.Run()
}
