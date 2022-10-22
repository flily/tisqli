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

func TestWeakStringToNumber(t *testing.T) {
	cases := []struct {
		s string
		i int64
		u uint64
		f float64
	}{
		{"0", 0, 0, 0.0},
		{"1", 1, 1, 1.0},
		{"233", 233, 233, 233.0},
		{"-1", -1, 0, -1.0},
	}

	for _, c := range cases {
		w := NewString(c.s)

		if v := w.AsInteger(); v != c.i {
			t.Errorf("w.AsInteger() = %v, want %v", v, c.i)
		}

		if v := w.AsUint(); v != c.u {
			t.Errorf("w.AsUint() = %v, want %v", v, c.u)
		}

		if v := w.AsFloat(); v != c.f {
			t.Errorf("w.AsFloat() = %v, want %v", v, c.f)
		}
	}
}

func TestStringAddNull(t *testing.T) {
	a := NewString("lorem")
	b := NewNull()
	c := a.Add(b)

	if c.Type() != ValueNull {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueNull)
	}

	if !c.IsNull() {
		t.Errorf("c.IsNull() = %v, want %v", c.IsNull(), true)
	}
}

func TestStringAddNumber(t *testing.T) {
	a := NewString("lorem")
	numbers := []WeakValue{
		NewInteger(42),
		NewUint(42),
		NewFloat(3.1415926),
	}

	for _, b := range numbers {
		c := a.Add(b)

		if c.Type() != ValueFloat {
			t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
		}

		if c.AsFloat() != b.AsFloat() {
			t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), b.AsFloat())
		}
	}
}

func TestStringAddString(t *testing.T) {
	a := NewString("lorem")
	b := NewString("ipsum")
	c := a.Add(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueInteger)
	}

	if c.AsFloat() != 0.0 {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), 0)
	}
}
