package main

import (
	"fmt"
	"sync"
)

var once sync.Once

// private singleton
type singleton struct{}

// private instance
var instance *singleton = new(singleton)

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("Never do anyting just for the money.")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
