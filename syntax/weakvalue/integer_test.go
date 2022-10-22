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

func TestIntegerAddNull(t *testing.T) {
	a := NewInteger(42)
	b := NewNull()
	c := a.Add(b)

	if c.Type() != ValueNull {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueNull)
	}

	if !c.IsNull() {
		t.Errorf("c.IsNull() = %v, want %v", c.IsNull(), true)
	}
}

func TestIntegerAddInteger(t *testing.T) {
	a := NewInteger(233)
	b := NewInteger(42)
	c := a.Add(b)
	if c.Type() != ValueInteger {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueInteger)
	}

	expected := int64(275)
	if c.AsInteger() != expected {
		t.Errorf("c.AsInteger() = %v, want %v", c.AsInteger(), expected)
	}
}

func TestIntegerAddUint(t *testing.T) {
	a := NewInteger(233)
	b := NewUint(42)
	c := a.Add(b)
	if c.Type() != ValueUint {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueInteger)
	}

	expected := uint64(275)
	if c.AsUint() != expected {
		t.Errorf("c.AsUint() = %v, want %v", c.AsUint(), expected)
	}
}

func TestIntegerAddFloat(t *testing.T) {
	a := NewInteger(233)
	b := NewFloat(3.1415926)
	c := a.Add(b)
	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 236.1415926
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestIntegerAddRegularString(t *testing.T) {
	a := NewInteger(233)
	b := NewString("lorem ipsum")
	c := a.Add(b)
	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 233.0
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestIntegerAddNumberString(t *testing.T) {
	a := NewInteger(233)

	{
		b := NewString("42")
		c := a.Add(b)
		if c.Type() != ValueFloat {
			t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
		}

		expected := 275.0
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

		expected := 236.1415926
		if c.AsFloat() != expected {
			t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
		}
	}
}

func TestIntegerSubNull(t *testing.T) {
	a := NewInteger(42)
	b := NewNull()
	c := a.Sub(b)

	if c.Type() != ValueNull {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueNull)
	}

	if !c.IsNull() {
		t.Errorf("c.IsNull() = %v, want %v", c.IsNull(), true)
	}
}

func TestIntegerSubInteger(t *testing.T) {
	a := NewInteger(233)
	b := NewInteger(42)
	c := a.Sub(b)
	if c.Type() != ValueInteger {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueInteger)
	}

	expected := int64(191)
	if c.AsInteger() != expected {
		t.Errorf("c.AsInteger() = %v, want %v", c.AsInteger(), expected)
	}
}

func TestIntegerSubUint(t *testing.T) {
	a := NewInteger(233)
	b := NewUint(42)
	c := a.Sub(b)
	if c.Type() != ValueUint {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueInteger)
	}

	expected := uint64(191)
	if c.AsUint() != expected {
		t.Errorf("c.AsUint() = %v, want %v", c.AsUint(), expected)
	}

}

func TestIntegerSubFloat(t *testing.T) {
	a := NewInteger(233)
	b := NewFloat(3.1415926)
	c := a.Sub(b)
	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 229.8584074
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestIntegerSubRegularString(t *testing.T) {
	a := NewInteger(233)
	b := NewString("lorem ipsum")
	c := a.Sub(b)
	if c.Type() != ValueFloat {
		t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
	}

	expected := 233.0
	if c.AsFloat() != expected {
		t.Errorf("c.AsFloat() = %v, want %v", c.AsFloat(), expected)
	}
}

func TestIntegerSubNumberString(t *testing.T) {
	a := NewInteger(233)

	{
		b := NewString("42")
		c := a.Sub(b)
		if c.Type() != ValueFloat {
			t.Errorf("c.Type() = %v, want %v", c.Type(), ValueFloat)
		}

		expected := 191.0
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
