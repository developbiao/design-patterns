package main

import "fmt"

// private singleton
type singleton struct{}

// private instance
var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("Never do anyting just for the money.")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
