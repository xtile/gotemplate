package main

import (
	"fmt"
	"sync"
	"time"
)

func getStringToPrint() string {
	return string("Hello, World!")
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(getStringToPrint())
}

func main() {
	var wg sync.WaitGroup

	for {
		wg.Add(1)
		go worker(&wg)

		wg.Wait()

		time.Sleep(time.Second * 10)
	}
}
