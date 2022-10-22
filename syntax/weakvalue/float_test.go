package weakvalue

import "testing"

func TestWeakFloat(t *testing.T) {
	{
		w := NewFloat(0.0)

		if w.Type() != ValueFloat {
			t.Errorf("w.Type() = %v, want %v", w.Type(), ValueFloat)
		}

		if s := w.String(); s != "0.000000" {
			t.Errorf("w.String() = %v, want %v", s, "0.000000")
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

		if v := w.AsString(); v != "0.000000" {
			t.Errorf("w.AsString() = %v, want %v", v, "0.000000")
		}
	}

	{
		w := NewFloat(42.0)

		if w.Type() != ValueFloat {
			t.Errorf("w.Type() = %v, want %v", w.Type(), ValueFloat)
		}

		if s := w.String(); s != "42.000000" {
			t.Errorf("w.String() = %v, want %v", s, "42.000000")
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

		if v := w.AsString(); v != "42.000000" {
			t.Errorf("w.AsString() = %v, want %v", v, "42.000000")
		}
	}
}
