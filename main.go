package main

import (
	"fmt"
	"sync"
)

var (
	mu          = &sync.Mutex{}
	balance int = 0
)

func deposit(money int, w *sync.WaitGroup) {
	mu.Lock()
	defer func() {
		mu.Unlock()
		w.Done()
	}()
	balance += money
}

func finalBalance() int {
	return balance
}

func main() {
	w := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go deposit(100, w)
	}
	w.Wait()

	fmt.Println("Balance : ", finalBalance())
}
