package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var y float32
	x := rand.Int32N(100)

	if x <= 20 {
		y = float32((x * x) - 4 + x)
	} else {
		y = (5.0 * float32(x)) / (4 * (float32(x) * float32(x)))
	}
	fmt.Printf("Змінна х: %d\n", x)
	fmt.Printf("Результат: %f\n", y)
}
