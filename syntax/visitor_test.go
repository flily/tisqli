package syntax

import (
	"testing"
)

func TestVisitorLead(t *testing.T) {
	visitor := NewVisitor()

	if l := visitor.Lead(); l != "" {
		t.Errorf("visitor.Lead() = '%s', want ''", l)
	}

	visitor.Level = 1
	if l := visitor.Lead(); l != "  +-" {
		t.Errorf("visitor.Lead() = '%s', want '  +-'", l)
	}

	visitor.Level = 2
	if l := visitor.Lead(); l != "  | +-" {
		t.Errorf("visitor.Lead() = '%s', want '  | +-'", l)
	}

	visitor.Level = 3
	if l := visitor.Lead(); l != "  | | +-" {
		t.Errorf("visitor.Lead() = '%s', want '  | | +-'", l)
	}
}
