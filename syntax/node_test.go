package syntax

import (
	"reflect"
	"testing"

	"github.com/pingcap/tidb/parser/ast"
)

func newTestNode(key int) ast.Node {
	n := &ast.SelectField{
		Offset: key,
	}
	return n
}

func TestNodeChildren(t *testing.T) {
	root := &Node{
		Node: newTestNode(1),
	}

	if last := root.Last(); last != nil {
		t.Errorf("root.Last() = %v, want nil", last)
	}

	root.NewChild(newTestNode(10))
	root.NewChild(newTestNode(20))
	root.NewChild(newTestNode(30))
	root.NewChild(newTestNode(40))

	last := root.Last()
	if last == nil {
		t.Fatal("root.Last() = nil, want not nil")
	}

	if node, ok := last.Node.(*ast.SelectField); !ok {
		t.Errorf("root.Last().Node.(type) = %v, want *ast.SelectField", reflect.TypeOf(node))
	} else {
		if node.Offset != 40 {
			t.Errorf("root.Last().Node.(*ast.SelectField).Offset = %d, want 40", node.Offset)
		}
	}
}
