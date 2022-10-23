package syntax

import (
	"fmt"
	"reflect"

	"github.com/pingcap/tidb/parser/ast"
	driver "github.com/pingcap/tidb/types/parser_driver"
)

// Node is a wrap for ast.Node, add pointers to parent and children
type Node struct {
	ast.Node

	Parent     *Node
	Children   []*Node
	IsConstant bool
}

func (n *Node) NewChild(node ast.Node) *Node {
	child := &Node{
		Node:     node,
		Parent:   n,
		Children: make([]*Node, 0),
	}

	n.Children = append(n.Children, child)
	return child
}

func (n *Node) Last() *Node {
	if len(n.Children) <= 0 {
		return nil
	}

	return n.Children[len(n.Children)-1]
}

func (n *Node) AllChildrenConstant() bool {
	if len(n.Children) <= 0 {
		return n.IsConstant
	}

	for _, child := range n.Children {
		if !child.IsConstant {
			return false
		}
	}

	return true
}

func (n *Node) VerifyConstant() {
	for _, child := range n.Children {
		child.VerifyConstant()
	}

	if n.AllChildrenConstant() {
		n.IsConstant = true
	}

	switch n.Node.(type) {
	case *driver.ValueExpr:
		n.IsConstant = true

	case *ast.FuncCallExpr:
		if len(n.Children) <= 0 {
			n.IsConstant = true
		}
	}
}

func (n *Node) Info() string {
	switch a := n.Node.(type) {
	case *ast.ColumnName:
		return a.Name.O

	case *ast.FuncCallExpr:
		return fmt.Sprintf("%s([%d])", a.FnName.O, len(a.Args))

	case *ast.SelectField:
		if a.WildCard != nil {
			return fmt.Sprintf("%s.*", a.WildCard.Table.O)
		}

		return a.AsName.O

	case *ast.BinaryOperationExpr:
		return fmt.Sprintf("<%s>", a.Op)

	case *driver.ValueExpr:
		return fmt.Sprintf("%s <%v>", a.Type.String(), a.GetValue())

	default:
		return a.Text()
	}
}

func (n *Node) Message() string {
	if n.IsConstant {
		return "const"
	}

	return ""
}

func (n *Node) PrintTree(level int, lead string, last bool) {
	if level <= 0 {
		fmt.Printf("ROOT\n")

	} else {
		next := lead + "+-+ "
		fmt.Printf("%s%#T (%d)  '%s': %s\n",
			next, n.Node, n.Node.OriginTextPosition(),
			n.Info(), n.Message())
	}

	for i, child := range n.Children {
		nextLast := i == len(n.Children)-1
		if last {
			child.PrintTree(level+1, lead+"  ", nextLast)
		} else {
			child.PrintTree(level+1, lead+"| ", nextLast)
		}
	}
}

func (n *Node) TypeEqual(other *Node) bool {
	typeMe := reflect.TypeOf(n.Node)
	typeOther := reflect.TypeOf(other.Node)
	if typeMe != typeOther {
		return false
	}

	if len(n.Children) != len(other.Children) {
		return false
	}

	for i, child := range n.Children {
		if !child.TypeEqual(other.Children[i]) {
			return false
		}
	}

	return true
}
