package syntax

import (
	"testing"
)

func TestCStringStrip(t *testing.T) {
	cases := []struct {
		Input    string
		Expected string
	}{
		{"lorem ipsum", "lorem ipsum"},
		{"lorem\x00ipsum", "lorem"},
	}

	for _, c := range cases {
		got := CStringStrip(c.Input)
		if got != c.Expected {
			t.Errorf("CStringStrip('%s') = '%s', want '%s'", c.Input, got, c.Expected)
		}
	}
}
