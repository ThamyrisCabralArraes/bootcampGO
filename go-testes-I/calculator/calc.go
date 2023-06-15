package calculator

import (
	"errors"
	"sort"
)

// sum two values
func Sum(n1, n2 int) int {
	return n1 + n2
}

// sub two values
func Sub(n1, n2 int) int {
	return n1 - n2
}

// Div two values
func Div(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, errors.New("nenhum numero pode ser divisivel por 0")
	}

	div := n1 / n2
	return div, nil
}

func SortAscending(numbers []int) {
	sort.Ints(numbers)
}
