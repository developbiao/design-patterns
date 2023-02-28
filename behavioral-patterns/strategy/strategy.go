package main

import "fmt"

// Weapon strategy abstarct
type WeaponStrategy interface {
	UseWeapon()
}

// Concreate strategy
type Ak47 struct{}

func (ak *Ak47) UseWeapon() {
	fmt.Println("Using Ak47 fighting")
}

type Knife struct{}

func (k *Knife) UseWeapon() {
	fmt.Println("UsingKnife fighting")
}

// Environment
type Hero struct {
	strategy WeaponStrategy //  Have an abstract strategy
}

// Set strategy
func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

// Call strategy
func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}

func main() {
	hero := Hero{}
	// Using Knife
	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()

	// Using Ak47
	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	fmt.Println("Main func done!")
}
