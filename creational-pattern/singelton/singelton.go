package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// flag for initializaed
var initialized uint32

var lock sync.Mutex

// private singleton
type singleton struct{}

// private instance
var instance *singleton = new(singleton)

func GetInstance() *singleton {
	// Check flag
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		instance = new(singleton)
		// Set flag
		atomic.StoreUint32(&initialized, 1)
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
