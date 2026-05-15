package atomic_variant

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Start() {
	evenChan := make(chan int)
	oddChan := make(chan int)

	var counter int64
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

					atomic.AddInt64(&counter, 1)
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

					atomic.AddInt64(&counter, -1)
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

	finalCounter := atomic.LoadInt64(&counter)
	fmt.Printf("Final value of counter in atomic variant: %d\n", finalCounter)
}
