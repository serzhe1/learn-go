package _for

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expect := 6

		if got != expect {
			t.Errorf("expect %d, got %d, numbers %v", expect, got, numbers)
		}
	})
}
