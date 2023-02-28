package main

import "fmt"

// Sell strategy
// Shopping malls have strategy A (20 % off)
// and strategy B (consumption over 200, cash back 100),
// use the strategy mode to simulate the scene
type SellStrategy interface {
	GetPrice(price float64) float64
}

type StrategyA struct{}

func (sa *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("Execute strategy-A, all product 80%")
	return price * 0.8
}

type StrategyB struct{}

func (sb *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("Execute strategy-B, consumption over 200, cash back 100")
	if price >= 200 {
		price -= 100
	}
	return price
}

// Environment
type Goods struct {
	Price    float64
	Strategy SellStrategy
}

func (g *Goods) SetStrategy(s SellStrategy) {
	g.Strategy = s
}

func (g *Goods) SellPrice() float64 {
	fmt.Println("Original Price ", g.Price, ".")
	return g.Strategy.GetPrice(g.Price)
}

func main() {
	adidas := Goods{
		Price: 320.0,
	}
	// AM, shopping mall strategy-A
	adidas.SetStrategy(new(StrategyA))
	fmt.Println("AM adidas sell shoes price ", adidas.SellPrice())

	// PM, shopping mall strategy-B
	adidas.SetStrategy(new(StrategyB))
	fmt.Println("PM adidas sell shoes price ", adidas.SellPrice())
	fmt.Println("Main func done!")
}
