package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Adder(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %q, instead we got %q", expected, sum)
	}
}

func ExampleAdder() {
	sum := Adder(1, 5)
	fmt.Println(sum)
	// Output: 6
}
