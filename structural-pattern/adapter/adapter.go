package main

import "fmt"

// Adapter target
type V5 interface {
	Use5V()
}

// Buz dependency V5 interface
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("Phone chargeing...")
	p.v.Use5V()
}

// Adaptee
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("Using 220V")
}

// Powr adapter
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("Using adapter charge")
	// call adaptee method
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

func main() {
	iphone := NewPhone(NewAdapter(new(V220)))
	iphone.Charge()
}
