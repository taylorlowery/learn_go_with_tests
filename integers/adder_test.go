package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(4, 5)
	fmt.Println(sum)
	// Output: 9
}

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	if sum != expected {
		t.Errorf("expected '%d'. got '%d'", expected, sum)
	}
}
