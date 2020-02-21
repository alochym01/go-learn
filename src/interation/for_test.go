package interation

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("want '%q' but got '%q'", want, got)
		}
	}
	got := Repeat("a")
	want := "aaaaa"
	assertCorrectMessage(t, got, want)
}

func BenchmarkRepeatFactor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatFactor("a")
	}
}
