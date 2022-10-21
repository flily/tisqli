package checker

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/flily/tisqli/syntax"
)

type partialSQLTemplate struct {
	Template       string
	CorrectPayload string
}

func (t partialSQLTemplate) Correct() string {
	return fmt.Sprintf(t.Template, t.CorrectPayload)
}

func (t partialSQLTemplate) Build(input string) string {
	return fmt.Sprintf(t.Template, input)
}

type PartialSQLCheckResult struct {
	IsInjection bool
	Template    string
	Payload     string
	Reason      string
	Err         error
}

func (r *PartialSQLCheckResult) SQL() string {
	return fmt.Sprintf(r.Template, r.Payload)
}

func (r *PartialSQLCheckResult) SQLInColour() string {
	var c *color.Color
	if r.IsInjection {
		c = color.New(
			color.FgYellow,
			color.BgRed,
		)
	} else {
		c = color.New(
			color.FgYellow,
			color.BgGreen,
		)
	}

	payload := c.Sprint(r.Payload)
	return fmt.Sprintf(r.Template, payload)
}

func (t partialSQLTemplate) Check(payload string) *PartialSQLCheckResult {
	result := &PartialSQLCheckResult{
		Template: t.Template,
		Payload:  t.CorrectPayload,
	}
	astCorrect, _, err := syntax.Parse(t.Correct())
	if err != nil {
		result.Reason = "template error"
		result.Err = err
		return result
	}

	result.Payload = payload
	astPartial, _, err := syntax.Parse(t.Build(payload))
	if err != nil {
		result.Reason = "syntax error"
		result.Err = err
		return result
	}

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

var sqlTemplates = []partialSQLTemplate{
	{"SELECT * FROM users WHERE id = %s AND name = 'lorem'", "42"},
	{"SELECT * FROM users WHERE id = 42 AND name = '%s'", "ipsum"},
}

func OnPartial(part string) *PartialResult {
	templateList := sqlTemplates

	result := &PartialResult{
		Results: make([]PartialSQLCheckResult, len(templateList)),
	}
	for i, template := range templateList {
		r := template.Check(part)
		result.Results[i] = *r
	}

	return result
}
