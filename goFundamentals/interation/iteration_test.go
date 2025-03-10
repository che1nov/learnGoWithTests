package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		char     string
		count    int
		expected string
	}{
		{"d", 3, "ddd"},
		{"a", 5, "aaaaa"},
		{"f", 3, "fff"},
	}

	for _, test := range tests {
		repeated := Repeat(test.char, test.count)
		if repeated != test.expected {
			t.Errorf("expected %q but got %q", test.expected, repeated)
		}
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeatd := Repeat("a", 5)
	fmt.Println(repeatd)
	//Output: "aaaaa"
}
