package checker

import (
	"testing"
)

func TestFullResultIsInjection(t *testing.T) {
	result := &FullResult{
		AllowMultipleStatements: false,
	}

	if result.IsInjection() {
		t.Errorf("result.IsInjection() == true, expected false")
	}

	result.Reason = FullCheckReasonMoreStatements
	if !result.IsInjection() {
		t.Errorf("result.IsInjection() == false, expected true")
	}

	result.AllowMultipleStatements = true
	if result.IsInjection() {
		t.Errorf("result.IsInjection() == true, expected false")
	}

	result.AllowMultipleStatements = false
	if !result.IsInjection() {
		t.Errorf("result.IsInjection() == false, expected true")
	}

	result.Elements = []FullElementResult{
		{},
	}

	if !result.IsInjection() {
		t.Errorf("result.IsInjection() == false, expected true")
	}
}

func TestFullCheckerWithNormalSQL(t *testing.T) {
	checker := DefaultFullChecker()
	payload := "SELECT * FROM users"

	result := checker.Check(payload)
	if result == nil {
		t.Fatalf("NewFullChecker().Check(%s) should return result", payload)
	}

	if result.IsInjection() {
		t.Errorf("result.IsInjection() == true, expected false")
	}
}

func TestFullCheckerWithTautology(t *testing.T) {
	checker := DefaultFullChecker()
	payload := "SELECT * FROM users WHERE 1 = 1"

	result := checker.Check(payload)
	if result == nil {
		t.Fatalf("NewFullChecker().Check(%s) should return result", payload)
	}

	if !result.IsInjection() {
		t.Errorf("result.IsInjection() == false, expected true")
	}
}

func TestFullCheckerWithMultipleSQLs(t *testing.T) {
	checker := DefaultFullChecker()
	payload := "SELECT * FROM users; SELECT * FROM users"

	result := checker.Check(payload)
	if result == nil {
		t.Fatalf("NewFullChecker().Check(%s) should return result", payload)
	}

	if !result.IsInjection() {
		t.Errorf("result.IsInjection() == false, expected true")
	}

	result.AllowMultipleStatements = true
	if result.IsInjection() {
		t.Errorf("result.IsInjection() == true, expected false")
	}
}

func TestFullCheckerWithStaticSelectList(t *testing.T) {
	checker := DefaultFullChecker()
	payload := "SELECT * FROM users UNION SELECT 1"

	result := checker.Check(payload)
	if result == nil {
		t.Fatalf("NewFullChecker().Check(%s) should return result", payload)
	}

	if !result.IsInjection() {
		t.Errorf("result.IsInjection() == false, expected true")
	}
}

func TestFullCheckerWithConstantFunctions(t *testing.T) {
	checker := DefaultFullChecker()
	payload := "SELECT * FROM users UNION SELECT NOW()"

	result := checker.Check(payload)
	if result == nil {
		t.Fatalf("NewFullChecker().Check(%s) should return result", payload)
	}

	if !result.IsInjection() {
		t.Errorf("result.IsInjection() == false, expected true")
	}
}

func TestFullCheckerWithSyntaxError(t *testing.T) {
	checker := DefaultFullChecker()
	payload := "SELECT"

	result := checker.Check(payload)
	if result == nil {
		t.Fatalf("NewFullChecker().Check(%s) should return result", payload)
	}

	if result.Reason != FullCheckReasonSyntaxError {
		t.Errorf("result.Reason == %s, expected %s", result.Reason, FullCheckReasonSyntaxError)
	}

	if result.IsInjection() {
		t.Errorf("result.IsInjection() == true, expected false")
	}
}
