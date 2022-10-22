package weakvalue

import (
	"testing"
)

func TestWeakInteger(t *testing.T) {
	{
		w := NewInteger(0)

		if w.Type() != ValueInteger {
			t.Errorf("w.Type() = %v, want %v", w.Type(), ValueInteger)
		}

		if s := w.String(); s != "0" {
			t.Errorf("w.String() = %v, want %v", s, "0")
		}

		if v := w.IsNull(); v != false {
			t.Errorf("w.IsNull() = %v, want %v", v, false)
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

		if v := w.AsString(); v != "0" {
			t.Errorf("w.AsString() = %v, want %v", v, "0")
		}
	}

	{
		w := NewInteger(42)

		if w.Type() != ValueInteger {
			t.Errorf("w.Type() = %v, want %v", w.Type(), ValueInteger)
		}

		if s := w.String(); s != "42" {
			t.Errorf("w.String() = %v, want %v", s, "42")
		}

		if v := w.IsNull(); v != false {
			t.Errorf("w.IsNull() = %v, want %v", v, false)
		}

		if v := w.AsBoolean(); v != true {
			t.Errorf("w.AsBoolean() = %v, want %v", v, true)
		}

		if v := w.AsInteger(); v != 42 {
			t.Errorf("w.AsInteger() = %v, want %v", v, 42)
		}

		if v := w.AsUint(); v != 42 {
			t.Errorf("w.AsUint() = %v, want %v", v, 42)
		}

		if v := w.AsFloat(); v != 42.0 {
			t.Errorf("w.AsFloat() = %v, want %v", v, 42.0)
		}

		if v := w.AsString(); v != "42" {
			t.Errorf("w.AsString() = %v, want %v", v, "42")
		}
	}
}

func TestWeakBooleans(t *testing.T) {
	{
		w := NewFalse()

		if w.Type() != ValueInteger {
			t.Errorf("w.Type() = %v, want %v", w.Type(), ValueInteger)
		}

		if s := w.String(); s != "0" {
			t.Errorf("w.String() = %v, want %v", s, "0")
		}

		if v := w.IsNull(); v != false {
			t.Errorf("w.IsNull() = %v, want %v", v, false)
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

		if v := w.AsString(); v != "0" {
			t.Errorf("w.AsString() = %v, want %v", v, "0")
		}
	}

	{
		w := NewTrue()

		if w.Type() != ValueInteger {
			t.Errorf("w.Type() = %v, want %v", w.Type(), ValueInteger)
		}

		if s := w.String(); s != "1" {
			t.Errorf("w.String() = %v, want %v", s, "1")
		}

		if v := w.IsNull(); v != false {
			t.Errorf("w.IsNull() = %v, want %v", v, false)
		}

		if v := w.AsBoolean(); v != true {
			t.Errorf("w.AsBoolean() = %v, want %v", v, true)
		}

		if v := w.AsInteger(); v != 1 {
			t.Errorf("w.AsInteger() = %v, want %v", v, 1)
		}

		if v := w.AsUint(); v != 1 {
			t.Errorf("w.AsUint() = %v, want %v", v, 1)
		}

		if v := w.AsFloat(); v != 1.0 {
			t.Errorf("w.AsFloat() = %v, want %v", v, 1.0)
		}

		if v := w.AsString(); v != "1" {
			t.Errorf("w.AsString() = %v, want %v", v, "1")
		}
	}
}
