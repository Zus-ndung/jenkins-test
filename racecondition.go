package main

import (
	"fmt"
	"sync"
)

var lock = make(chan bool, 1)
var x = 0

func add(lock chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	lock <- true
	x = x + 1
	<-lock
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add(lock, wg)
	}
	wg.Wait()
	fmt.Println(x)
}
