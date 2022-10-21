package checker

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/flily/tisqli/syntax"
)

type PartialSQLTemplate struct {
	Template       string
	CorrectPayload string
}

func (t PartialSQLTemplate) Correct() string {
	return fmt.Sprintf(t.Template, t.CorrectPayload)
}

func (t PartialSQLTemplate) Build(input string) string {
	return fmt.Sprintf(t.Template, input)
}

type PartialSQLCheckResult struct {
	IsInjection bool
	Template    string
	Payload     string
	Reason      string
	Err         error
	AstCorrect  []*syntax.Node
	AstPartial  []*syntax.Node
}

func (r *PartialSQLCheckResult) SQL() string {
	return fmt.Sprintf(r.Template, r.Payload)
}

func (r *PartialSQLCheckResult) SQLInColour() string {
	if !r.IsInjection {
		payload := color.New(
			color.BgGreen,
		).Sprintf(r.Payload)
		return fmt.Sprintf(r.Template, payload)
	}

	partInjected := color.New(
		color.BgRed,
	)

	payload := partInjected.Sprintf(r.Payload)

	if i := strings.Index(r.Payload, "\x00"); i >= 0 {
		payloadAffected := partInjected.Sprintf(r.Payload[:i])

		partTerminated := color.New(color.BgYellow)
		payloadTerminated := partTerminated.Sprintf(r.Payload[i:])
		payload = payloadAffected + payloadTerminated
	}

	return fmt.Sprintf(r.Template, payload)
}

func (t PartialSQLTemplate) Check(payload string) *PartialSQLCheckResult {
	result := &PartialSQLCheckResult{
		Template: t.Template,
		Payload:  t.CorrectPayload,
	}

	parser := syntax.NewParser()
	astCorrect, _, err := parser.Parse(t.Correct())
	if err != nil {
		result.Reason = "template error"
		result.Err = err
		return result
	}

	result.AstCorrect = astCorrect
	result.Payload = payload
	astPartial, _, err := parser.Parse(t.Build(payload))
	if err != nil {
		result.Reason = "syntax error"
		result.Err = err
		return result
	}

	result.AstPartial = astPartial
	if len(astPartial) != 1 {
		result.IsInjection = true
		result.Reason = "none or multiple SQls"
		return result
	}

	if !astCorrect[0].TypeEqual(astPartial[0]) {
		result.IsInjection = true
		result.Reason = "AST modified"
		return result
	}

	return result
}

type PartialResult struct {
	Results []PartialSQLCheckResult
}

func (r *PartialResult) IsInjection() bool {
	for _, result := range r.Results {
		if result.IsInjection {
			return true
		}
	}

	return false
}

var sqlTemplates = []PartialSQLTemplate{
	{"SELECT * FROM users WHERE id = %s AND name = 'lorem'", "42"},
	{"SELECT * FROM users WHERE (id = %s) AND name = 'lorem'", "42"},
	{"SELECT * FROM users WHERE ((id = %s)) AND name = 'lorem'", "42"},
	{"SELECT * FROM users WHERE (((id = %s))) AND name = 'lorem'", "42"},
	{"SELECT * FROM users WHERE id = 42 AND name = '%s'", "ipsum"},
	{"SELECT * FROM users WHERE id = 42 AND (name = '%s')", "ipsum"},
	{"SELECT * FROM users WHERE id = 42 AND ((name = '%s'))", "ipsum"},
	{"SELECT * FROM users WHERE id = 42 AND (((name = '%s')))", "ipsum"},
	{"SELECT * FROM users WHERE id = 42 AND name = \"%s\"", "ipsum"},
	{"SELECT * FROM users WHERE id = 42 AND (name = \"%s\")", "ipsum"},
	{"SELECT * FROM users WHERE id = 42 AND ((name = \"%s\"))", "ipsum"},
	{"SELECT * FROM users WHERE id = 42 AND (((name = \"%s\")))", "ipsum"},
	{"SELECT * FROM %s WHERE id = 42 AND name = 'ipsum'", "users"},
	{"SELECT * FROM user JOIN college ON user.id = %s WHERE user.name = 'lorem'", "42"},
	{"SELECT * FROM user JOIN college ON user.name = '%s' WHERE user.name = 'lorem'", "lorem"},
}

type PartialChecker struct {
	Templates []PartialSQLTemplate
	Decoder   *Decoder
}

func NewPartialChecker(templates []PartialSQLTemplate, decoder *Decoder) *PartialChecker {
	c := &PartialChecker{
		Templates: templates,
		Decoder:   decoder,
	}
	return c
}

func DefaultPartialChecker() *PartialChecker {
	decoder := DefaultDecoders()
	return NewPartialChecker(sqlTemplates, decoder)
}

func (c *PartialChecker) Check(raw string) *PartialResult {
	payload := raw
	if c.Decoder != nil {
		payload = c.Decoder.Decode(raw)
	}

	result := &PartialResult{
		Results: make([]PartialSQLCheckResult, len(c.Templates)),
	}
	for i, template := range c.Templates {
		r := template.Check(payload)
		result.Results[i] = *r
	}

	return result
}
