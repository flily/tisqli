package syntax

import "github.com/pingcap/tidb/parser"

func Parse(sql string) ([]*Node, []error, error) {
	parser := parser.New()
	nodes, warns, err := parser.Parse(sql, "", "")
	if err != nil {
		return nil, warns, err
	}

	result := make([]*Node, len(nodes))
	for i, node := range nodes {
		visitor := NewVisitor()
		node.Accept(visitor)

		result[i] = visitor.Root
	}

	return result, warns, nil
}
