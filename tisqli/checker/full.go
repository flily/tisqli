package checker

import (
	"github.com/flily/tisqli/syntax"
	"github.com/pingcap/tidb/parser/ast"
)

const (
	FullCheckReasonOK                       = "ok"
	FullCheckReasonModified                 = "AST modified"
	FullCheckReasonMoreStatements           = "none or multiple SQls"
	FullCheckReasonSyntaxError              = "syntax error"
	FullCheckReasonConstantBinaryExpression = "constant binary expression"
	FullCheckReasonConstantSelectStatement  = "constant select statement"
)

type FullElementResult struct {
	Reason string
	Text   string
}

type FullResult struct {
	Err                     error
	Reason                  string
	Elements                []FullElementResult
	AllowMultipleStatements bool
}

func (r *FullResult) IsInjection() bool {
	if r.Reason == FullCheckReasonMoreStatements && !r.AllowMultipleStatements {
		return true
	}

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
		element.Reason = FullCheckReasonConstantBinaryExpression

	case *ast.SelectStmt:
		element.Reason = FullCheckReasonConstantSelectStatement
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
		result.Reason = FullCheckReasonSyntaxError
		result.Err = err
		return result
	}

	if len(nodes) != 1 {
		result.Reason = FullCheckReasonMoreStatements
		return result
	}

	root := nodes[0]
	root.VerifyConstant()
	IsFullInjection(root, result)
	return result
}
