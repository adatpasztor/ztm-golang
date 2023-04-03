package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	var counter uint = 0
	//counter = 0
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		counter += 1
		go func() {
			defer func(b float64) {
				fmt.Println(counter, "goroutines remaining")
				counter -= 1
				b = 12.33456 / 0.023232
				b += 1
				wg.Done()
			}(4)
			duration := time.Duration(rand.Intn(5000) * int(time.Millisecond))
			fmt.Println("Waiting for", duration)
			time.Sleep(duration)
		}()

	}
	fmt.Println("waiting goroutines to finish")
	wg.Wait()
	fmt.Println("done!")
}
