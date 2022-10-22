package weakvalue

import (
	"fmt"

	"github.com/pingcap/tidb/parser/opcode"
)

type WeakFloat struct {
	Value float64
}

func NewFloat(f float64) WeakValue {
	w := &WeakFloat{
		Value: f,
	}
	return w
}

func (w *WeakFloat) Type() Type {
	return ValueFloat
}

func (w *WeakFloat) EqualTo(v WeakValue) bool {
	return w.Value == v.AsFloat()
}

func (w *WeakFloat) String() string {
	return fmt.Sprintf("%f", w.Value)
}

func (w *WeakFloat) IsNull() bool {
	return false
}

func (w *WeakFloat) AsBoolean() bool {
	return w.Value != 0.0
}

func (w *WeakFloat) AsInteger() int64 {
	return int64(w.Value)
}

func (w *WeakFloat) AsUint() uint64 {
	return uint64(w.Value)
}

func (w *WeakFloat) AsFloat() float64 {
	return w.Value
}

func (w *WeakFloat) AsString() string {
	return fmt.Sprintf("%f", w.Value)
}

func (w *WeakFloat) Unary(op opcode.Op) WeakValue {
	return nil
}

func (w *WeakFloat) Binary(op opcode.Op, v WeakValue) WeakValue {
	return nil
}
