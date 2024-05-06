package addition_test

import (
	"github.com/bytemare/actions/internal"
	"testing"
)

type test struct {
	a, b, result int
}

var tests = []test{
	{2, 3, 5},
	{0, 0, 0},
}

func TestAddition(t *testing.T) {
	for i, addition := range tests {
		result := internal.Addition(addition.a, addition.b)
		if result != addition.result {
			t.Fatalf("%d: invalid result. Expected %d + %d = %d, got %d",
				i, addition.a, addition.b, addition.result, result)
		}
	}
}

func FuzzAddition(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		expected := a + b
		result := internal.Addition(a, b)
		if result != expected {
			t.Errorf("invalid result. Expected %d + %d = %d, got %d", a, b, expected, result)
		}
	})
}
