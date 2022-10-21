package syntax

import (
	"testing"
)

func TestParserOnSimpleError(t *testing.T) {
	sql := "select"

	parser := NewParser()
	nodes, _, err := parser.Parse(sql)
	if len(nodes) > 0 {
		t.Errorf("len(nodes) = %d, want 0", len(nodes))
	}

	if err == nil {
		t.Errorf("err = nil, want not nil")
	}

	if parseError, ok := err.(*ParserError); !ok {
		t.Errorf("err.(type) = %v, want *ParserError", err)
	} else {
		if parseError.Line != 1 || parseError.Column != 6 {
			t.Errorf("err.(*ParserError) = %v, want {Line: 1, Column: 6}", parseError)
		}
	}
}

func TestParserOnSimpleSelect(t *testing.T) {
	sql := "select *"

	parser := NewParser()
	nodes, _, err := parser.Parse(sql)
	if len(nodes) != 1 {
		t.Errorf("len(nodes) = %d, want 1", len(nodes))
	}

	if err != nil {
		t.Errorf("expect err is nil, but got:")
		t.Errorf(err.Error())
	}
}
