package weakvalue

import (
	"fmt"

	"github.com/pingcap/tidb/parser/opcode"
)

// WeakInteger represents integer value.
type WeakInteger struct {
	Value int64
}

// NewInteger creates an instance of integer.
func NewInteger(i int64) WeakValue {
	w := &WeakInteger{
		Value: i,
	}
	return w
}

// WeakFalse creates an instance of false, as a integer.
func NewFalse() WeakValue {
	return NewInteger(0)
}

// WeakTrue creates an instance of true, as a integer.
func NewTrue() WeakValue {
	return NewInteger(1)
}

func (w *WeakInteger) Type() Type {
	return ValueInteger
}

// EqualTo
func (w *WeakInteger) EqualTo(v WeakValue) bool {
	return w.Value == v.AsInteger()
}

// String returns a string representation of the value.
func (w *WeakInteger) String() string {
	return fmt.Sprintf("%d", w.Value)
}

func (w *WeakInteger) IsNull() bool {
	return false
}

func (w *WeakInteger) AsBoolean() bool {
	return w.Value != 0
}

func (w *WeakInteger) AsInteger() int64 {
	return w.Value
}

func (w *WeakInteger) AsUint() uint64 {
	return uint64(w.Value)
}

func (w *WeakInteger) AsFloat() float64 {
	return float64(w.Value)
}

func (w *WeakInteger) AsString() string {
	return fmt.Sprintf("%d", w.Value)
}

func (w *WeakInteger) Unary(op opcode.Op) WeakValue {
	return genericUnary(op, w)
}

func (w *WeakInteger) Binary(op opcode.Op, v WeakValue) WeakValue {
	return genericBinary(op, w, v)
}

func (w *WeakInteger) Add(v WeakValue) WeakValue {
	return genericAdd(w, v)
}

func (w *WeakInteger) Sub(v WeakValue) WeakValue {
	return genericSub(w, v)
}
