package main

import (
	"fmt"
	"lab_03/calc"
)

func init() {
	fmt.Println("Preparing to calculation")
	fmt.Println()
}

func main() {
	defer func() {
		fmt.Println()
		fmt.Println("End of calculating")
	}()

	array := []float64{1, 2, 3, 4, 5, 6, -1, 10, 100}
	fmt.Printf("%.2f\n", calc.Sum(array...))
	fmt.Printf("%.2f\n", calc.Max(array...))
	fmt.Printf("%.2f\n", calc.Min(array...))
	res1, err1 := calc.Divide(100.25, 5)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Printf("%.2f\n", res1)
	}

	fmt.Println()
	fmt.Println("Structure operations")
	fmt.Println()

	c := calc.Calc{}

	array_structure := []float64{2, 45, 10, 500, 10, -200, -222.22}

	fmt.Printf("%.2f\n", c.Sum(array_structure...))
	fmt.Printf("%.2f\n", c.Max(array_structure...))
	fmt.Printf("%.2f\n", c.Min(array_structure...))
	res2, err2 := c.Divide(100.25, 5)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Printf("%.2f\n", res2)
	}
}
