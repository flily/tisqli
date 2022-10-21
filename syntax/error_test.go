package syntax

import (
	"fmt"
	"testing"
)

func TestParseInt(t *testing.T) {
	cases := []struct {
		Input    string
		Expected int
	}{
		{"", 0},
		{"0", 0},
		{"233", 233},
		{"-233", -233},
		{"not a number", 0},
	}

	for _, c := range cases {
		got := parseInt(c.Input)
		if got != c.Expected {
			t.Errorf("parseInt(%s) = %d, want %d", c.Input, got, c.Expected)
		}
	}
}

func TestParserErrorParseWithFlagFound(t *testing.T) {
	sql := "the quick brown fox jumps over the lazy dog"
	err := fmt.Errorf("TestError line 1 column 42: lorem ipsum")
	parseError := NewParserError(sql, err)
	if parseError == nil {
		t.Errorf("NewParserError(err) = nil, want not nil: %v", err)
	}

	e, ok := parseError.(*ParserError)
	if !ok {
		t.Errorf("err.(type) = %v, want *ParserError", err)
	} else {
		if e.Line != 1 || e.Column != 42 {
			t.Errorf("err.(*ParserError) = %v, want {Line: 1, Column: 42}", e)
		}
	}

	hint := e.Hint()
	expected := "" +
		"the quick brown fox jumps over the lazy dog\n" +
		"                                         ^\n" +
		"                                         |\n" +
		"                                         +-- TestError line 1 column 42: lorem ipsum"
	if hint != expected {
		t.Errorf("wrong err.Hint()")
		t.Errorf(hint)
	}

	if e.Error() != expected {
		t.Errorf("wrong err.Error()")
	}

	if e.Unwrap() != err {
		t.Errorf("wrong err.Unwrap()")
	}
}

func TestParserErrorParseWithFlagNotFound(t *testing.T) {
	sql := "the quick brown fox jumps over the lazy dog"
	err := fmt.Errorf("lorem ipsum")
	parseError := NewParserError(sql, err)
	if parseError != nil {
		t.Errorf("NewParserError(err) = nil, want not nil: %v", err)
	}
}
