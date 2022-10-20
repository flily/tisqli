package syntax

import (
	"strings"

	"github.com/pingcap/tidb/parser/ast"
)

type Visitor struct {
	Level int
	Root  *Node
	Node  *Node
}

func NewVisitor() *Visitor {
	root := &Node{
		Children: make([]*Node, 0),
	}
	v := &Visitor{
		Level: 0,
		Root:  root,
		Node:  root,
	}
	return v
}

func (v *Visitor) Lead() string {
	if v.Level == 0 {
		return ""
	}

	parts := []string{"  "}
	if v.Level > 1 {
		parts = append(parts, strings.Repeat("| ", v.Level-1))
	}

	parts = append(parts, "+-")
	return strings.Join(parts, "")
}

func (v *Visitor) Enter(in ast.Node) (ast.Node, bool) {
	v.Level += 1
	node := v.Node.NewChild(in)
	v.Node = node

	return in, false
}

func (v *Visitor) Leave(in ast.Node) (ast.Node, bool) {
	v.Level -= 1
	v.Node = v.Node.Parent.Last().Parent
	return in, true
}
