package checker

import (
	"testing"
)

func TestPartialSQLTemplate(t *testing.T) {
	template := &PartialSQLTemplate{
		Template:       "SELECT %s",
		CorrectPayload: "42",
	}

	{
		correct := template.Correct()
		expected := "SELECT 42"
		if correct != expected {
			t.Errorf("template.Correct() == '%s', expected '%s'", correct, expected)
		}
	}

	{
		sql := template.Build("233")
		expected := "SELECT 233"
		if sql != expected {
			t.Errorf("template.Build(233) == '%s', expected '%s'", sql, expected)
		}
	}
}

func TestBuildResultSQL(t *testing.T) {
	result := &PartialSQLCheckResult{
		IsInjection: false,
		Template:    "SELECT %s LIMIT 10",
		Payload:     "42",
	}

	{
		text := result.SQL()
		textExpected := "SELECT 42 LIMIT 10"
		if text != textExpected {
			t.Errorf("result.SQL() == '%s', expected '%s'", text, textExpected)
		}

		coloured := result.SQLInColour()
		colouredExpected := "SELECT " + colourCorrect.Sprintf("42") + " LIMIT 10"
		if coloured != colouredExpected {
			t.Errorf("result.SQLInColour() == '%s', expected '%s'", coloured, colouredExpected)
		}
	}

	result.IsInjection = true
	{
		text := result.SQL()
		textExpected := "SELECT 42 LIMIT 10"
		if text != textExpected {
			t.Errorf("result.SQL() == '%s', expected '%s'", text, textExpected)
		}

		coloured := result.SQLInColour()
		colouredExpected := "SELECT " + colourInjected.Sprintf("42") + " LIMIT 10"
		if coloured != colouredExpected {
			t.Errorf("result.SQLInColour() == '%s', expected '%s'", coloured, colouredExpected)
		}
	}
}

func TestBuildResultSQLWithTerminatedPayload(t *testing.T) {
	result := &PartialSQLCheckResult{
		IsInjection: true,
		Template:    "SELECT '%s' LIMIT 10",
		Payload:     "lorem\x00ipsum",
	}

	{
		text := result.SQL()
		textExpected := "SELECT 'lorem\x00ipsum' LIMIT 10"
		if text != textExpected {
			t.Errorf("result.SQL() == '%s', expected '%s'", text, textExpected)
		}

		coloured := result.SQLInColour()
		colouredExpected := "SELECT '" + colourInjected.Sprintf("lorem\x00") +
			colourTerminated.Sprint("ipsum") + "' LIMIT 10"
		if coloured != colouredExpected {
			t.Errorf("result.SQLInColour() == '%s', expected '%s'", coloured, colouredExpected)
		}
	}
}

func TestTemplateCheck(t *testing.T) {
	template := &PartialSQLTemplate{
		Template:       "SELECT * FROM t WHERE id = %s",
		CorrectPayload: "42",
	}

	cases := []struct {
		payload     string
		reason      string
		isInjection bool
		hasError    bool
	}{
		{"42", PartialCheckReasonOK, false, false},
		{"42 OR 1=1", PartialCheckReasonModified, true, false},
		{"42; SELECT *", PartialCheckReasonMoreStatements, true, false},
		{"42'", PartialCheckReasonSyntaxError, false, true},
	}

	for _, c := range cases {
		result := template.Check(c.payload)
		if c.hasError && result.Err == nil {
			t.Errorf("template.Check(%s) should return error", c.payload)
		}

		if result.IsInjection != c.isInjection {
			t.Errorf("template.Check(%s).IsInjection == %t, expected %t", c.payload, result.IsInjection, c.isInjection)
		}

		if result.Reason != c.reason {
			t.Errorf("template.Check(%s).Reason == %s, expected %s", c.payload, result.Reason, c.reason)
		}
	}
}

func TestTemplateCheckWithTemplateError(t *testing.T) {
	template := &PartialSQLTemplate{
		Template:       "SECT %s",
		CorrectPayload: "42",
	}

	cases := []struct {
		payload     string
		reason      string
		isInjection bool
		hasError    bool
	}{
		{"42", PartialCheckReasonTemplateError, false, true},
	}

	for _, c := range cases {
		result := template.Check(c.payload)
		if c.hasError && result.Err == nil {
			t.Errorf("template.Check(%s) should return error", c.payload)
		}

		if result.IsInjection != c.isInjection {
			t.Errorf("template.Check(%s).IsInjection == %t, expected %t", c.payload, result.IsInjection, c.isInjection)
		}

		if result.Reason != c.reason {
			t.Errorf("template.Check(%s).Reason == %s, expected %s", c.payload, result.Reason, c.reason)
		}
	}
}

func TestDefaultPartialCheckerWithInjection(t *testing.T) {
	checker := DefaultPartialChecker()
	payload := "1' or '1' = '1"

	result := checker.Check(payload)
	if result == nil {
		t.Fatalf("DefaultPartialChecker().Check(%s) should return result", payload)
	}

	if result.IsInjection() != true {
		t.Errorf("result.IsInjection == %t, expected %t", result.IsInjection(), true)
	}
}

func TestDefaultPartialCheckerWithCorrect(t *testing.T) {
	checker := DefaultPartialChecker()
	payload := "lorem ipsum"

	result := checker.Check(payload)
	if result == nil {
		t.Fatalf("DefaultPartialChecker().Check(%s) should return result", payload)
	}

	if result.IsInjection() != false {
		t.Errorf("result.IsInjection == %t, expected %t", result.IsInjection(), false)
	}
}
