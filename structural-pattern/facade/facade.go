package main

import "fmt"

type SubSystemA struct{}

func (sa *SubSystemA) MethodA() {
	fmt.Println("Sub system method-A")
}

type SubSystemB struct{}

func (sa *SubSystemB) MethodB() {
	fmt.Println("Sub system method-B")
}

type SubSystemC struct{}

func (sa *SubSystemC) MethodC() {
	fmt.Println("Sub system method-C")
}

type SubSystemD struct{}

func (sa *SubSystemD) MethodD() {
	fmt.Println("Sub system method-D")
}

// Facade  Reduced to a simple interface for use
type Facade struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
	d *SubSystemD
}

func (f *Facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f *Facade) MethodTwo() {
	f.c.MethodC()
	f.d.MethodD()
}

func main() {
	// If not using facade call MethodA() and Method B
	sa := new(SubSystemA)
	sa.MethodA()

	sb := new(SubSystemB)
	sb.MethodB()

	fmt.Println("------ Using Facade-------")
	f := Facade{
		a: new(SubSystemA),
		b: new(SubSystemB),
		c: new(SubSystemC),
		d: new(SubSystemD),
	}
	// Call facade wapper methods
	f.MethodOne()
}
