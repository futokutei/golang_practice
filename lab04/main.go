package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := make(chan int, 100)
	evenNumbers := make(chan int, 50)
	squares := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(4)

	go generate(numbers, &wg)
	go filterEven(evenNumbers, numbers, &wg)
	go square(squares, evenNumbers, &wg)
	go sum(squares, &wg)

	wg.Wait()

}

func generate(num chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(num)

	for i := 1; i <= 100; i++ {
		num <- i
	}
}

func filterEven(ev chan int, nums chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ev)

	for num := range nums {
		if num%2 == 0 {
			ev <- num
		}
	}
}

func square(sqr chan int, evs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(sqr)

	for ev := range evs {
		sqr <- ev * ev
	}
}

func sum(sqr chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sumOf := 0
	for sq := range sqr {
		sumOf += sq
	}

	fmt.Printf("Sum of squares: %d", sumOf)
}
