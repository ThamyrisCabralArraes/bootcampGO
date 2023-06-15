package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	num1 := 3
	num2 := 3
	expectedResults := 6

	result := Sum(num1, num2)

	if result != expectedResults {
		t.Errorf("A funçao sum retornou o resultado %v, mas deveria ser %v", expectedResults, result)
	}
}

func TestSub(t *testing.T) {
	num1 := 3
	num2 := 0
	expectedResults := 3

	result := Sub(num1, num2)

	assert.Equal(t, expectedResults, result, "devem ser iguais")
}

func TestDiv(t *testing.T) {
	num1 := 3
	num2 := 3
	expectedResults := 1
	// expectedError := "nenhum numero pode ser divisivel por 0"

	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		assert.Equal(t, expectedError, err, "nao pode usar zero")
	// 	}
	// }()

	result, err := Div(num1, num2)
	if num2 == 0 {
		assert.Error(t, err, "nao pode usar zero")
	}

	assert.Equal(t, expectedResults, result, "devem ser iguais")
	assert.NotNil(t, result)
}

func TestDiv2(t *testing.T) {
	num1 := 3
	num2 := 0
	expectedResults := "nenhum numero pode ser divisivel por 0"

	_, err := Div(num1, num2)

	assert.Equal(t, expectedResults, err.Error())
}

func TestSortAscending(t *testing.T) {
	numbers := []int{9, 5, 2, 7, 1, 3}
	numbersOk := []int{1, 2, 3, 5, 7, 9}

	SortAscending(numbers)

	assert.Equal(t, numbersOk, numbers, "sejam iguais")

}
