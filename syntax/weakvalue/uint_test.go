package weakvalue

import (
	"testing"
)

func TestWeakUint(t *testing.T) {
	{
		w := NewUint(0)

		if w.Type() != ValueUint {
			t.Errorf("w.Type() = %v, want %v", w.Type(), ValueUint)
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
		w := NewUint(42)

		if w.Type() != ValueUint {
			t.Errorf("w.Type() = %v, want %v", w.Type(), ValueUint)
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

func TestUintAddNull(t *testing.T) {
	a := NewUint(42)
	b := NewNull()
	c := a.Add(b)

	if c.Type() != ValueNull {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueNull)
	}

	if !c.IsNull() {
		t.Errorf("c.IsNull() = %v, want %v", c.IsNull(), true)
	}
}

func TestUintAddInteger(t *testing.T) {
	a := NewUint(42)
	b := NewInteger(233)
	c := a.Add(b)

	if c.Type() != ValueUint {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueUint)
	}

	expected := uint64(275)
	if c.AsUint() != expected {
		t.Errorf("c.AsInteger() = %v, want %v", c.AsInteger(), expected)
	}
}

func TestUintAddUint(t *testing.T) {
	a := NewUint(42)
	b := NewUint(233)
	c := a.Add(b)

	if c.Type() != ValueUint {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueUint)
	}

	expected := uint64(275)
	if c.AsUint() != expected {
		t.Errorf("c.AsInteger() = %v, want %v", c.AsInteger(), expected)
	}
}

func TestUintAddFloat(t *testing.T) {
	a := NewUint(42)
	b := NewFloat(233.0)
	c := a.Add(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 275.0
	if c.AsFloat() != expected {
		t.Errorf("c.AsInteger() = %v, want %v", c.AsInteger(), expected)
	}
}

func TestUintAddRegularString(t *testing.T) {
	a := NewUint(42)
	b := NewString("lorem ipsum")
	c := a.Add(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := float64(42.0)
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestUintAddNumberString(t *testing.T) {
	a := NewUint(42)

	{
		b := NewString("233")
		c := a.Add(b)

		if c.Type() != ValueFloat {
			t.Errorf("c.Type() = %v, want %v", c.Type(), ValueUint)
		}

		expected := float64(275.0)
		if c.AsFloat() != expected {
			t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
		}
	}

	{
		b := NewString("3.1415926")
		c := a.Add(b)

		if c.Type() != ValueFloat {
			t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
		}

		expected := 45.1415926
		if c.AsFloat() != expected {
			t.Errorf("c.AsInteger() = %v, want %v", c.AsInteger(), expected)
		}
	}
}

func TestUintSubNull(t *testing.T) {
	a := NewUint(42)
	b := NewNull()
	c := a.Sub(b)

	if c.Type() != ValueNull {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueNull)
	}

	if !c.IsNull() {
		t.Errorf("c.IsNull() = %v, want %v", c.IsNull(), true)
	}
}

func TestUintSubInteger(t *testing.T) {
	a := NewUint(233)
	b := NewInteger(42)
	c := a.Sub(b)

	if c.Type() != ValueUint {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueUint)
	}

	expected := uint64(191)
	if c.AsUint() != expected {
		t.Errorf("c.AsInteger() = %v, want %v", c.AsInteger(), expected)
	}
}

func TestUintSubUint(t *testing.T) {
	a := NewUint(233)
	b := NewUint(42)
	c := a.Sub(b)

	if c.Type() != ValueUint {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueUint)
	}

	expected := uint64(191)
	if c.AsUint() != expected {
		t.Errorf("c.AsInteger() = %v, want %v", c.AsInteger(), expected)
	}
}

func TestUintSubFloat(t *testing.T) {
	a := NewUint(233)
	b := NewFloat(42)
	c := a.Sub(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 191.0
	if c.AsFloat() != expected {
		t.Errorf("c.AsInteger() = %v, want %v", c.AsInteger(), expected)
	}
}

func TestUintSubRegularString(t *testing.T) {
	a := NewUint(42)
	b := NewString("lorem ipsum")
	c := a.Sub(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := float64(42.0)
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestUintSubNumberString(t *testing.T) {
	a := NewUint(233)

	{
		b := NewString("42")
		c := a.Sub(b)

		if c.Type() != ValueFloat {
			t.Errorf("c.Type() = %v, want %v", c.Type(), ValueUint)
		}

		expected := float64(191.0)
		if c.AsFloat() != expected {
			t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
		}
	}

	{
		b := NewString("3.1415926")
		c := a.Sub(b)

		if c.Type() != ValueFloat {
			t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
		}

		expected := 229.8584074
		if c.AsFloat() != expected {
			t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
		}
	}
}
