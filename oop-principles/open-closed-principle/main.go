package main

import "fmt"

type AbstractBanker interface {
	// Abstract busniess
	DoBusi()
}

// Save banker
type SaveBanker struct {
}

// Transfer banker
type TransferBanker struct {
}

// Pay banker
type PayBanker struct {
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("User starting save money")
}

func (tb *TransferBanker) DoBusi() {
	fmt.Println("User starting transfer money to other")
}

func (pb *PayBanker) DoBusi() {
	fmt.Println("User starting pay")
}

func BankerBusiness(banker AbstractBanker) {
	banker.DoBusi()
}

func main() {
	// Save money
	BankerBusiness(&SaveBanker{})

	// Tranfer money
	BankerBusiness(&TransferBanker{})

	// Pay money
	BankerBusiness(&PayBanker{})
}
