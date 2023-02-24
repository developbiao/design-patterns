package main

import "fmt"

// ===== Abstract layout =======

// Abstarct produt
type Fruit interface {
	Show()
}

// Abstract factory
type AbstractFactory interface {
	CreateFruit() Fruit
}

// ===== Base layout =======

// Concreate product
// Cheery
type Cherry struct {
	Fruit
}

func (c *Cherry) Show() {
	fmt.Println("I' am cherry")
}

type Apple struct {
	Fruit
}

func (a *Apple) Show() {
	fmt.Println("I' am apple")
}

type Pear struct {
	Fruit
}

func (p *Pear) Show() {
	fmt.Println("I' am pear")
}

// Conecreate Factory
type CherryFactory struct {
	AbstractFactory
}

func (fac *CherryFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Cherry)
	return fruit
}

type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Pear)
	return fruit
}

func main() {
	// Need a concreate apple
	var cherryFactory AbstractFactory
	cherryFactory = new(CherryFactory)
	var cherry Fruit
	cherry = cherryFactory.CreateFruit()
	cherry.Show()

	// New a concreate pear
	var pearFactory AbstractFactory
	pearFactory = new(PearFactory)
	var pear Fruit
	pear = pearFactory.CreateFruit()
	pear.Show()

	// New a concreate  apple
	var appleFactory AbstractFactory
	appleFactory = new(AppleFactory)
	var apple Fruit
	apple = appleFactory.CreateFruit()
	apple.Show()

}
