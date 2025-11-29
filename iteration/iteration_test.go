package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times test", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, got, expected)
	})
	t.Run("repeat 7 times test", func(t *testing.T) {
		got := Repeat("a", 7)
		expected := "aaaaaaa"
		assertCorrectMessage(t, got, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	//Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
