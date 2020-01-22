package mix

import (
	"testing"
)

func TestRunShell(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 20},
	}
	for _, c := range cases {
		got := RunShell()
		if got != c.want {
			t.Errorf("RunShell(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// func TestLs() {

// }
