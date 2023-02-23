package main

import "fmt"

// ===> Abstract layerout  <===
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// ===> Implement layerout  <===
type Benz struct {
}

type Bmw struct {
}

func (bmw *Bmw) Run() {
	fmt.Println("Bmw is Running")
}

func (benz *Benz) Run() {
	fmt.Println("Benz is running")
}

type ZhangSan struct {
}

type WangEr struct {
}

func (zs *ZhangSan) Drive(car Car) {
	fmt.Println("Zhang san drive car")
}

func (we *WangEr) Drive(car Car) {
	fmt.Println("Wang er drive car")
}

func main() {
	bmw := new(Bmw)
	benz := new(Benz)

	zhangsan := &ZhangSan{}
	zhangsan.Drive(bmw)
	zhangsan.Drive(benz)

	wanger := &WangEr{}
	wanger.Drive(benz)

}
