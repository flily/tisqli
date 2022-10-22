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

func TestFloatAddNull(t *testing.T) {
	a := NewFloat(3.1415926)
	b := NewNull()
	c := a.Add(b)

	if c.Type() != ValueNull {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueNull)
	}

	if !c.IsNull() {
		t.Errorf("c.IsNull() = %v, want %v", c.IsNull(), true)
	}
}

func TestFloatAddInteger(t *testing.T) {
	a := NewFloat(3.1415926)
	b := NewInteger(42)
	c := a.Add(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 45.1415926
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestFloatAddUint(t *testing.T) {
	a := NewFloat(3.1415926)
	b := NewUint(42)
	c := a.Add(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 45.1415926
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestFloatAddFloat(t *testing.T) {
	a := NewFloat(3.1415926)
	b := NewFloat(42.0)
	c := a.Add(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 45.1415926
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestFloatAddRegularString(t *testing.T) {
	a := NewFloat(3.1415926)
	b := NewString("lorem ipsum")
	c := a.Add(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 3.1415926
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestFloatAddNumberString(t *testing.T) {
	a := NewFloat(3.1415926)

	{
		b := NewString("42")
		c := a.Add(b)

		if c.Type() != ValueFloat {
			t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
		}

		expected := 45.1415926
		if c.AsFloat() != expected {
			t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
		}
	}

	{
		b := NewString("2.71828")
		c := a.Add(b)

		if c.Type() != ValueFloat {
			t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
		}

		expected := 5.8598726
		if c.AsFloat() != expected {
			t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
		}
	}
}

func TestFloatSubNull(t *testing.T) {
	a := NewFloat(3.1415926)
	b := NewNull()
	c := a.Sub(b)

	if c.Type() != ValueNull {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueNull)
	}

	if !c.IsNull() {
		t.Errorf("c.IsNull() = %v, want %v", c.IsNull(), true)
	}
}

func TestFloatSubInteger(t *testing.T) {
	a := NewFloat(3.1415926)
	b := NewInteger(42)
	c := a.Sub(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := -38.8584074
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestFloatSubUint(t *testing.T) {
	a := NewFloat(3.1415926)
	b := NewUint(42)
	c := a.Sub(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := -38.8584074
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestFloatSubFloat(t *testing.T) {
	a := NewFloat(3.1415926)
	b := NewFloat(42.0)
	c := a.Sub(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := -38.8584074
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestFloatSubRegularString(t *testing.T) {
	a := NewFloat(3.1415926)
	b := NewString("lorem ipsum")
	c := a.Sub(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 3.1415926
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestFloatSubNumberString(t *testing.T) {
	a := NewFloat(3.1415926)

	b := NewString("42")
	c := a.Sub(b)

	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := -38.8584074
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}
