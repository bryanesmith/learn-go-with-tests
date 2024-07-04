package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("Sum an array of len 5", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		got := Sum(input)
		want := 15
		assertCorrectMessage(t, got, want, input)
	})

	t.Run("Sum an array of len 3", func(t *testing.T) {
		input := []int{1, 2, 3}
		got := Sum(input)
		want := 6
		assertCorrectMessage(t, got, want, input)
	})

	t.Run("Sum an array of len 0", func(t *testing.T) {
		input := []int{}
		got := Sum(input)
		want := 0
		assertCorrectMessage(t, got, want, input)
	})

}

func assertCorrectMessage(t testing.TB, got, want int, input []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given %v", got, want, input)
	}
}
