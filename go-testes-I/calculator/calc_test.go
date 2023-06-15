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
	num2 := 3
	expectedResults := 0

	result := Sub(num1, num2)

	assert.Equal(t, expectedResults, result, "devem ser iguais")
}
