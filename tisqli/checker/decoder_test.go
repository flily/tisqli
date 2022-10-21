package checker

import (
	"testing"
)

func TestURLDecode(t *testing.T) {
	cases := []struct {
		Input    string
		Expected string
	}{
		{"lorem ipsum", "lorem ipsum"},
		{"c%2b%2b", "c++"},
		{"c%loremipsum", "c%loremipsum"},
	}

	for _, c := range cases {
		result := URLDecode(c.Input)
		if result != c.Expected {
			t.Errorf("URLDecode(%s) == %s, expected %s", c.Input, result, c.Expected)
		}
	}
}

func TestNilDecoder(t *testing.T) {
	var decoder *Decoder
	result := decoder.Decode("lorem ipsum")
	if result != "lorem ipsum" {
		t.Errorf("decoder.Decode(lorem ipsum) == %s", result)
	}
}

func TestDecoder(t *testing.T) {
	decoder := NewDecoder(
		URLDecode,
	)
	result := decoder.Decode("lorem%20ipsum")
	if result != "lorem ipsum" {
		t.Errorf("decoder.Decode(lorem%%20ipsum) == %s", result)
	}
}

func TestDefaultDecoder(t *testing.T) {
	decoder := DefaultDecoders()
	result := decoder.Decode("lorem%20ipsum")
	if result != "lorem ipsum" {
		t.Errorf("decoder.Decode(lorem%%20ipsum) == %s", result)
	}
}
