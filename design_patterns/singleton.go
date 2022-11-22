package main

import (
	"fmt"
	"sync"
	"time"
)

type Singleton struct {
}

var lock = &sync.Mutex{}
var singleton *Singleton

func GetInstance() *Singleton {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleton == nil {
			singleton = &Singleton{}
			fmt.Println("Singleton object created")
		} else {
			fmt.Println("Singleton object already created")
		}
	} else {
		fmt.Println("Singleton object already created.")
	}
	return singleton
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(wg *sync.WaitGroup) {
			GetInstance()
			wg.Done()
		}(wg)

		go func(wg *sync.WaitGroup) {
			GetInstance()
			wg.Done()
		}(wg)
	}
	time.Sleep(100 * time.Millisecond)
	wg.Wait()
}
