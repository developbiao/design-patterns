package main

import "fmt"

// Abstract class, make drink, wrapper all the implementation steps of the template
type Beverage interface {
	BoilWater() // boil water
	Brew()      // brew
	PourInCup() // pour in cup
	AddThings() // Add things

	WantAddThings() bool // Add things hook
}

// Template let the specific production process be inhertied adn realized
type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	// Subclasses can override this method to decide whether to perform the follwing actions
	if t.b.WantAddThings() {
		// b is Beverage, which is the interface of MakeCaffee, here you need to assign a value to the interface, pointing to the specific subclass object
		// Trigger the polymorphic feature of all interface methods of b.
		t.b.AddThings()
	}
}

// Concreate template make coffee
type MakeCaffee struct {
	template
}

func NewMakeCaffee() *MakeCaffee {
	makeCaffee := new(MakeCaffee)
	//
	makeCaffee.b = makeCaffee
	return makeCaffee
}

func (mc *MakeCaffee) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

func (mc *MakeCaffee) Brew() {
	fmt.Println("用水冲咖啡豆")
}

func (mc *MakeCaffee) PourInCup() {
	fmt.Println("将冲好的咖啡倒入陶瓷杯中")
}

func (mc *MakeCaffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func (mc *MakeCaffee) WantAddThings() bool {
	return true //enable hook
}

// Concreate template Make tea
type MakeTea struct {
	template
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func (mt *MakeTea) BoilWater() {
	fmt.Println("将水煮到80摄氏度")
}

func (mt *MakeTea) Brew() {
	fmt.Println("用水冲茶叶")
}

func (mt *MakeTea) PourInCup() {
	fmt.Println("将冲好的茶叶放入个有情调的玻璃杯中")
}

func (mt *MakeTea) AddThings() {
	fmt.Println("添加茉莉花")
}

func (mt *MakeTea) WantAddThings() bool {
	return false
}

func main() {
	makeTea := NewMakeTea()
	makeTea.MakeBeverage()

	fmt.Println("---------")

	makeCaffee := NewMakeCaffee()
	makeCaffee.MakeBeverage()
}
