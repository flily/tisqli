package checker

import (
	"fmt"

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

func (t partialSQLTemplate) Check(part string) bool {
	astCorrect, _, err := syntax.Parse(t.Correct())
	if err != nil {
		return false
	}

	astPartial, _, err := syntax.Parse(t.Build(part))
	if err != nil {
		return false
	}

	if len(astPartial) != 1 {
		return true
	}

	if !astCorrect[0].TypeEqual(astPartial[0]) {
		return true
	}

	return false
}

var sqlTemplates = []partialSQLTemplate{
	{"SELECT * FROM users WHERE id = %s AND name = 'lorem'", "42"},
	{"SELECT * FROM users WHERE id = 42 AND name = '%s'", "ipsum"},
}

func OnPartial(part string) bool {
	for _, template := range sqlTemplates {
		if template.Check(part) {
			return true
		}
	}

	return false
}
