package main

import (
	"awesomeProject/lib"
	"fmt"
	"sync"
)

func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)

	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-evenChan:
				if !ok {
					return
				}
				if val%3 == 0 {
					mu.Lock()
					counter++
					mu.Unlock()
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-oddChan:
				if !ok {
					return
				}
				if val%33 == 0 {
					mu.Lock()
					counter--
					mu.Unlock()
				}
			}
		}
	}()

	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			evenChan <- i
		} else {
			oddChan <- i
		}
	}

	close(evenChan)
	close(oddChan)

	wg.Wait()

	fmt.Printf("Final counter value in mutex variant: %d\n", counter)
	atomic_variant.Start()
}
