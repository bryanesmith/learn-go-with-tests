package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat positive number of times", func(t *testing.T) {
		received := Repeat("a", 3)
		expected := "aaa"
		assertCorrectMessage(t, received, expected)
	})

	t.Run("Repeat 0 times", func(t *testing.T) {
		received := Repeat("a", 0)
		expected := ""
		assertCorrectMessage(t, received, expected)
	})

	t.Run("Repeat negative number of times", func(t *testing.T) {
		received := Repeat("a", -1)
		expected := ""
		assertCorrectMessage(t, received, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleRepeat() {
	str := Repeat("ㅎ", 5)
	fmt.Println(str)
	// Output: ㅎㅎㅎㅎㅎ
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
