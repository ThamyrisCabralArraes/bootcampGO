package main

import (
	"GO-TESTES-I/calculator"
	"fmt"
)

func main() {
	sum := calculator.Sum(2, 2)
	sub := calculator.Sub(2, 2)
	div, err := calculator.Div(10, 2)
	numbers := []int{9, 5, 2, 7, 1, 3}
	calculator.SortAscending(numbers)

	fmt.Println(numbers)
	fmt.Println(err)
	fmt.Println(sum, sub, div)
}
