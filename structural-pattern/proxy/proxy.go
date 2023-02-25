package main

import "fmt"

type Goods struct {
	Kind string // Goods kind
	Fact bool   // Goods fact or in fact
}

// ======= Abstract Layout ========
type Shopping interface {
	Buy(good *Goods) // Something task
}

// ======= Implementation Layout ========

type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("Go to korea shooping, buy a ", goods.Kind)
}

type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("Go to American shopping, buy a ", goods.Kind)
}

type AfrikaShopping struct{}

func (as *AfrikaShopping) Buy(goods *Goods) {
	fmt.Println("Go to Afrika shopping, buy a ", goods.Kind)
}

// ---- Overseas proxy ------
type OverseasProxy struct {
	shopping Shopping // proxy someting subject
}

func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping}
}

func (op *OverseasProxy) Buy(goods *Goods) {
	// 1. Check goods
	if op.distinguish(goods) == true {
		// 2. Buy goods
		op.shopping.Buy(goods)

		// 3. Customs security check
		op.check(goods)
	}

}

func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("Start Verifycation [", goods.Kind, "] fact check")
	if goods.Fact == false {
		fmt.Println("Detected in fact goods ", goods.Kind, "Can not buy!")
	}
	return goods.Fact
}

func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("Check [", goods.Kind, "] customs inspection carried out successfully!")
}

func main() {
	g1 := Goods{
		Kind: "韩国白菜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "Iphone 14 pro",
		Fact: false,
	}

	// Customer self buy
	var shopping Shopping
	shopping = new(AmericanShopping)
	shopping.Buy(&g1)
	shopping.Buy(&g2)

	fmt.Println("---- Overseas proxy shopping -----")

	// Customer buy via proxy
	var overseasProxy Shopping
	overseasProxy = NewProxy(shopping)
	overseasProxy.Buy(&g1)
	overseasProxy.Buy(&g2)
}
