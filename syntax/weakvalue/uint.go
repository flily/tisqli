package weakvalue

import (
	"fmt"

	"github.com/pingcap/tidb/parser/opcode"
)

type WeakUint struct {
	Value uint64
}

func NewUint(u uint64) WeakValue {
	w := &WeakUint{
		Value: u,
	}
	return w
}

func (w *WeakUint) Type() Type {
	return ValueUint
}

func (w *WeakUint) EqualTo(v WeakValue) bool {
	return w.Value == v.AsUint()
}

func (w *WeakUint) String() string {
	return fmt.Sprintf("%d", w.Value)
}

func (w *WeakUint) IsNull() bool {
	return false
}

func (w *WeakUint) AsBoolean() bool {
	return w.Value != 0
}

func (w *WeakUint) AsInteger() int64 {
	return int64(w.Value)
}

func (w *WeakUint) AsUint() uint64 {
	return w.Value
}

func (w *WeakUint) AsFloat() float64 {
	return float64(w.Value)
}

func (w *WeakUint) AsString() string {
	return fmt.Sprintf("%d", w.Value)
}

func (w *WeakUint) Unary(op opcode.Op) WeakValue {
	return genericUnary(op, w)
}

func (w *WeakUint) Binary(op opcode.Op, v WeakValue) WeakValue {
	return genericBinary(op, w, v)
}

func (w *WeakUint) Add(v WeakValue) WeakValue {
	return genericAdd(w, v)
}

func (w *WeakUint) Sub(v WeakValue) WeakValue {
	return genericSub(w, v)
}
