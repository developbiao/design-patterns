package main

import "fmt"

// ====== Abstract layout =======
type AbstractCherry interface {
	ShowCherry()
}

type AbstractApple interface {
	ShowApple()
}

type AbstractPear interface {
	ShowPear()
}

// Abstract factory
type AbstractFactory interface {
	CreateCherry() AbstractCherry
	CreateApple() AbstractApple
	CreatePear() AbstractPear
}

// ======= Implement Layout ======
// China product family
type ChinaCherry struct{}

func (cc *ChinaCherry) ShowCherry() {
	fmt.Println("Made in china cherry")
}

type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("Made in china apple")
}

type ChinaPear struct{}

func (cp *ChinaPear) ShowPear() {
	fmt.Println("Made in china pear")
}

// China factory
type ChinaFactory struct{}

func (cf *ChinaFactory) CreateCherry() AbstractCherry {
	var cherry AbstractCherry
	cherry = new(ChinaCherry)
	return cherry
}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinaPear)
	return pear
}

// ===== Japanese product family
type JapanCherry struct{}

func (jc *JapanCherry) ShowCherry() {
	fmt.Println("Made in japan cherry")
}

type JapanApple struct{}

func (ja *JapanApple) ShowApple() {
	fmt.Println("Made in japan apple")
}

type JapanPear struct{}

func (jp *JapanPear) ShowPear() {
	fmt.Println("Made in japan pear")
}

// Japanese factory
type JapanFactory struct{}

func (jf *JapanFactory) CreateCherry() AbstractCherry {
	var cherry AbstractCherry
	cherry = new(JapanCherry)
	return cherry
}

func (jf *JapanFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(JapanApple)
	return apple
}

func (jf *JapanFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(JapanPear)
	return pear
}

func main() {
	// Need china apple, cherry, pear
	// 1. Create Factory
	var cFactory AbstractFactory
	cFactory = new(ChinaFactory)

	// 2. Make apple
	var cApple AbstractApple
	cApple = cFactory.CreateApple()
	cApple.ShowApple()

	var cCherry AbstractCherry
	cCherry = cFactory.CreateCherry()
	cCherry.ShowCherry()

	var cPear AbstractPear
	cPear = cFactory.CreatePear()
	cPear.ShowPear()

	// Need japan cherry and pear
	var jFactory AbstractFactory
	jFactory = new(JapanFactory)

	var jCherry AbstractCherry
	jCherry = jFactory.CreateCherry()
	jCherry.ShowCherry()

	var jPear AbstractPear
	jPear = jFactory.CreatePear()
	jPear.ShowPear()

}
