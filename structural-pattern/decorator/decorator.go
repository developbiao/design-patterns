package main

import "fmt"

// ===== Abstract Layout  ========
// Abstract Component
type Phone interface {
	Show()
}

// Decorator shoud be interface but Golang interface synax can not
// Has membmer attribute
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

// ----- Implementation Layout ------
// Concreate Component
type HuaWei struct{}

func (hw *HuaWei) Show() {
	fmt.Println("Show HuaWei meta Phone")
}

type Iphone struct{}

func (ip *Iphone) Show() {
	fmt.Println("Show Iphone 14 pro")
}

// Concreate deorator
type MoDecorator struct {
	Decorator // Inherit the basic decorator class (mainly inherit the Phone member attributes)
}

func (md *MoDecorator) Show() {
	md.phone.Show() // Call the original method of the decorated component
	fmt.Println("Film mobile phone")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("Phone with case")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

// ----- Logic ------
func main() {
	var huawei Phone
	huawei = new(HuaWei)
	huawei.Show() // call original component method
	fmt.Println("----------")

	// Film mobile
	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei)
	moHuawei.Show()
	fmt.Println("----------")

	// Case mobile
	var keHuaWei Phone
	keHuaWei = NewKeDecorator(huawei)
	keHuaWei.Show()
	fmt.Println("----------")

	// File and case mobile
	var KeAndMoHuaWei Phone
	KeAndMoHuaWei = NewMoDecorator(keHuaWei)
	KeAndMoHuaWei.Show()
}
