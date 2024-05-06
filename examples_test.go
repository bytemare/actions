package actions_test

import (
	"fmt"

	"github.com/bytemare/actions/internal"
)

// Example_Addition shows how to add numbers.
func Example_addition() {
	a := 2
	b := 3

	fmt.Printf("%d + %d = %d\n", a, b, internal.Addition(a, b))
}
