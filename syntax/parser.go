package syntax

import (
	"github.com/pingcap/tidb/parser"
)

type Parser struct {
	parser       *parser.Parser
	StripCString bool
}

func NewParser() *Parser {
	p := &Parser{
		parser:       parser.New(),
		StripCString: true,
	}

	return p
}

func (p *Parser) Parse(sql string) ([]*Node, []error, error) {
	if p.StripCString {
		sql = CStringStrip(sql)
	}

	nodes, warns, err := p.parser.Parse(sql, "", "")
	if err != nil {
		return nil, warns, NewParserError(sql, err)
	}

	result := make([]*Node, len(nodes))
	for i, node := range nodes {
		visitor := NewVisitor()
		node.Accept(visitor)

		result[i] = visitor.Root
	}

	return result, warns, nil
}
