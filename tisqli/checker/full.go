package checker

import (
	"github.com/flily/tisqli/syntax"
	"github.com/pingcap/tidb/parser/ast"
)

const (
	FullCheckReasonOK                 = "ok"
	FullCheckReasonModified           = "AST modified"
	FullCheckMoreStatements           = "none or multiple SQls"
	FullCheckSyntaxError              = "syntax error"
	FullCheckConstantBinaryExpression = "constant binary expression"
	FullCheckConstantSelectStatement  = "constant select statement"
)

type FullElementResult struct {
	Reason string
	Text   string
}

type FullResult struct {
	Err      error
	Reason   string
	Elements []FullElementResult
}

func (r *FullResult) IsInjection() bool {
	return len(r.Elements) > 0
}

type FullChecker struct {
	Decoder *Decoder
}

func NewFullChecker(decoder *Decoder) *FullChecker {
	c := &FullChecker{
		Decoder: decoder,
	}
	return c
}

func DefaultFullChecker() *FullChecker {
	decoder := NewDecoder()
	return NewFullChecker(decoder)
}

func fullWalkNode(node *syntax.Node, result *FullResult) {
	for _, child := range node.Children {
		fullWalkNode(child, result)
	}

	if !node.IsConstant {
		return
	}

	element := &FullElementResult{}
	switch node.Node.(type) {
	case *ast.BinaryOperationExpr:
		element.Reason = FullCheckConstantBinaryExpression

	case *ast.SelectStmt:
		element.Reason = FullCheckConstantSelectStatement
	}

	if len(element.Reason) > 0 {
		result.Elements = append(result.Elements, *element)
	}
}

func IsFullInjection(node *syntax.Node, result *FullResult) {
	fullWalkNode(node, result)
}

func (c *FullChecker) Check(raw string) *FullResult {
	sql := c.Decoder.Decode(raw)
	result := &FullResult{
		Reason: "",
	}

	parser := syntax.NewParser()
	nodes, _, err := parser.Parse(sql)
	if err != nil {
		result.Reason = FullCheckSyntaxError
		result.Err = err
		return result
	}

	if len(nodes) != 1 {
		result.Reason = FullCheckMoreStatements
		return result
	}

	root := nodes[0]
	root.VerifyConstant()
	IsFullInjection(root, result)
	return result
}
