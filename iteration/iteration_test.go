package iteration

import (
	"fmt"
	"testing"
)

func TestIterate(t *testing.T) {
	got := Repeat("a", 3)
	want := "aaa"

	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	out := Repeat("a", 10)
	fmt.Println(out)
	// Output: aaaaaaaaaa
}
