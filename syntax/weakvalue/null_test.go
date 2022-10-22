package weakvalue

import (
	"testing"
)

func TestWeakNull(t *testing.T) {
	w := NewNull()

	if w.Type() != ValueNull {
		t.Errorf("w.Type() = %v, want %v", w.Type(), ValueNull)
	}

	if s := w.String(); s != "NULL" {
		t.Errorf("w.String() = %v, want %v", s, "null")
	}

	if v := w.IsNull(); v != true {
		t.Errorf("w.IsNull() = %v, want %v", v, true)
	}

	if v := w.AsBoolean(); v != false {
		t.Errorf("w.AsBoolean() = %v, want %v", v, false)
	}

	if v := w.AsInteger(); v != 0 {
		t.Errorf("w.AsInteger() = %v, want %v", v, 0)
	}

	if v := w.AsUint(); v != 0 {
		t.Errorf("w.AsUint() = %v, want %v", v, 0)
	}

	if v := w.AsFloat(); v != 0.0 {
		t.Errorf("w.AsFloat() = %v, want %v", v, 0.0)
	}

	if v := w.AsString(); v != "" {
		t.Errorf("w.AsString() = %v, want %v", v, "null")
	}
}

func TestNullAdd(t *testing.T) {
	null := NewNull()

	oprands := []WeakValue{
		NewNull(),
		NewFalse(),
		NewTrue(),
		NewInteger(0),
		NewInteger(42),
		NewFloat(0.0),
		NewFloat(3.14),
		NewString(""),
		NewString("lorem ipsum"),
	}

	for _, oprand := range oprands {
		r := null.Add(oprand)
		if r.Type() != ValueNull {
			t.Errorf("r.Type() = %v, want %v", r.Type(), ValueNull)
		}

		if r.IsNull() != true {
			t.Errorf("r.IsNull() = %v, want %v", r.IsNull(), true)
		}
	}
}

func TestNullSub(t *testing.T) {
	null := NewNull()

	oprands := []WeakValue{
		NewNull(),
		NewFalse(),
		NewTrue(),
		NewInteger(0),
		NewInteger(42),
		NewFloat(0.0),
		NewFloat(3.14),
		NewString(""),
		NewString("lorem ipsum"),
	}

	for _, oprand := range oprands {
		r := null.Sub(oprand)
		if r.Type() != ValueNull {
			t.Errorf("r.Type() = %v, want %v", r.Type(), ValueNull)
		}

		if r.IsNull() != true {
			t.Errorf("r.IsNull() = %v, want %v", r.IsNull(), true)
		}
	}
}
