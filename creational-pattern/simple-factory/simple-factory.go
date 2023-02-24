package main

import "fmt"

// ===== Abstract Layout =====

// Abstract fruit
type Fruit interface {
	Show()
}

// ===== Base Modules =====

// Fruit base class
type BaseFruit struct {
	Name  string
	Price int32
}

func (b *BaseFruit) Sell(weight int32) int32 {
	return b.Price * weight
}

type Apple struct {
	base BaseFruit
}

func (a *Apple) Show() {
	fmt.Printf("I am %s\n", a.base.Name)
}

type Banana struct {
	base BaseFruit
}

func (b *Banana) Show() {
	fmt.Printf("I am %s\n", b.base.Name)
}

type Pear struct {
	base BaseFruit
}

func (p *Pear) Show() {
	fmt.Printf("I am %s \n", p.base.Name)
}

type Factory struct {
}

func (f *Factory) Create(kind string) Fruit {
	var fruit Fruit
	if kind == "Apple" {
		fruit = &Apple{
			base: BaseFruit{
				Price: 33,
				Name:  kind,
			},
		}
	} else if kind == "Banana" {
		fruit = &Banana{
			base: BaseFruit{
				Price: 21,
				Name:  kind,
			},
		}
	} else if kind == "Pear" {
		fruit = &Pear{
			base: BaseFruit{
				Price: 21,
				Name:  kind,
			},
		}

	}
	return fruit
}

func main() {
	factory := &Factory{}
	apple := factory.Create("Apple")
	apple.Show()
	// Try to conver to apple kind
	if a, ok := (apple).(*Apple); ok {
		fmt.Println("Apple sell fee", a.base.Sell(33))
	}

	banana := factory.Create("Banana")
	banana.Show()

	pear := factory.Create("Pear")
	pear.Show()

}
