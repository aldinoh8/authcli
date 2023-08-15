package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("test helpers start")

	m.Run()

	fmt.Println("test helpers finish ..")
}

func TestSumToN(t *testing.T) {
	testCases := []struct {
		actual   int
		expected int
		message  string
	}{
		{
			actual:   SumToN(3),
			expected: 6,
			message:  "result should be equal 6",
		},
		{
			actual:   SumToN(4),
			expected: 10,
			message:  "result should be equal 10",
		},
		{
			actual:   SumToN(5),
			expected: 15,
			message:  "result should be equal 15",
		},
		{
			actual:   SumToN(6),
			expected: 21,
			message:  "result should be equal 21",
		},
	}

	for _, v := range testCases {
		t.Run(v.message, func(t *testing.T) {
			assert.Equal(t, v.expected, v.actual)
			assert.Greater(t, v.actual, 0)
			assert.Less(t, 0, v.actual)
			assert.NotNil(t, v.actual)
		})
	}
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, 9, Multiply(3, 3), "result should be equal 9")
	assert.Equal(t, 12, Multiply(4, 3), "result should be equal 12")
}

// func TestAnotherFunc2(t *testing.T) {

// }

// func TestAnotherFunc3(t *testing.T) {

// }
