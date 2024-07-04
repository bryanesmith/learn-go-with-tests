package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum an array of len 5", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		got := Sum(input)
		want := 15
		assertSum(t, got, want, input)
	})

	t.Run("Sum an array of len 3", func(t *testing.T) {
		input := []int{1, 2, 3}
		got := Sum(input)
		want := 6
		assertSum(t, got, want, input)
	})

	t.Run("Sum an array of len 0", func(t *testing.T) {
		input := []int{}
		got := Sum(input)
		want := 0
		assertSum(t, got, want, input)
	})

}

func TestSumAll(t *testing.T) {
	t.Run("Sum two arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertSumAll(t, got, want)
	})
	t.Run("Sum one arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3})
		want := []int{6}
		assertSumAll(t, got, want)
	})
	t.Run("Sum no arrays", func(t *testing.T) {
		got := SumAll()
		want := []int{}
		assertSumAll(t, got, want)
	})
}

func assertSum(t testing.TB, got, want int, input []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d, input: %v", got, want, input)
	}
}

func assertSumAll(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
