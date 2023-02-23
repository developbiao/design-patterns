package main

import "fmt"

type GunShop struct{}

type GunDance struct{}

func (gs *GunShop) onShopping() {
	fmt.Println("Gun for shopping")
}

func (gd *GunDance) onDancing() {
	fmt.Println("Gun for on dancing")
}

func main() {
	gs := new(GunShop)
	gs.onShopping()

	gd := new(GunDance)
	gd.onDancing()
}
