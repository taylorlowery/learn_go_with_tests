package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}

func TestRepeat(t *testing.T) {
	t.Run("3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		if repeated != expected {
			t.Errorf("got '%s', want '%s'", repeated, expected)
		}
	})
	t.Run("0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		if repeated != expected {
			t.Errorf("got '%s', want '%s'", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
