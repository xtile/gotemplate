package main

import (
	"fmt"
	"io"
	"net/http"
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

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, getStringToPrint())
	}

	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Listing for requests at http://address:8000/hello")
	http.ListenAndServe(":8000", nil)

	for {
		wg.Add(1)
		go worker(&wg)

		wg.Wait()

		time.Sleep(time.Second * 10)
	}
}
