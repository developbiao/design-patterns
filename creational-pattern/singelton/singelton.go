package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

// private singleton
type singleton struct{}

// private instance
var instance *singleton = new(singleton)

func GetInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("Never do anyting just for the money.")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
