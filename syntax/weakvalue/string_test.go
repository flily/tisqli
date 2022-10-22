package weakvalue

import (
	"testing"
)

func TestWeakString(t *testing.T) {
	{
		w := NewString("")

		if w.Type() != ValueString {
			t.Errorf("w.Type() = %v, want %v", w.Type(), ValueString)
		}

		if s := w.String(); s != `''` {
			t.Errorf("w.String() = %v, want %v", s, `''`)
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

		if v := w.AsString(); v != "" {
			t.Errorf("w.AsString() = %v, want %v", v, "")
		}
	}

	{
		w := NewString("lorem ipsum")

		if w.Type() != ValueString {
			t.Errorf("w.Type() = %v, want %v", w.Type(), ValueString)
		}

		if s := w.String(); s != `'lorem ipsum'` {
			t.Errorf("w.String() = %v, want %v", s, `'lorem ipsum'`)
		}

		if v := w.IsNull(); v != false {
			t.Errorf("w.IsNull() = %v, want %v", v, false)
		}

		if v := w.AsBoolean(); v != true {
			t.Errorf("w.AsBoolean() = %v, want %v", v, true)
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

		if v := w.AsString(); v != "lorem ipsum" {
			t.Errorf("w.AsString() = %v, want %v", v, "lorem ipsum")
		}

		if v := w.AsString(); v != "lorem ipsum" {
			t.Errorf("w.AsString() = %v, want %v", v, "lorem ipsum")
		}

		if v := w.AsString(); v != "lorem ipsum" {
			t.Errorf("w.AsString() = %v, want %v", v, "lorem ipsum")
		}
	}
}
