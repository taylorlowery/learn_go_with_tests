package arrays_and_slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		want := 10
		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, numbers)
		}
	})
}
