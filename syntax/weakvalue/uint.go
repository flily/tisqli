package weakvalue

import (
	"fmt"

	"github.com/pingcap/tidb/parser/opcode"
)

type Uint struct {
	Value uint64
}

func NewUint(u uint64) WeakValue {
	w := &Uint{
		Value: u,
	}
	return w
}

func (w *Uint) Type() Type {
	return ValueUint
}

func (w *Uint) EqualTo(v WeakValue) bool {
	return w.Value == v.AsUint()
}

func (w *Uint) String() string {
	return fmt.Sprintf("%d", w.Value)
}

func (w *Uint) IsNull() bool {
	return false
}

func (w *Uint) AsBoolean() bool {
	return w.Value != 0
}

func (w *Uint) AsInteger() int64 {
	return int64(w.Value)
}

func (w *Uint) AsUint() uint64 {
	return w.Value
}

func (w *Uint) AsFloat() float64 {
	return float64(w.Value)
}

func (w *Uint) AsString() string {
	return fmt.Sprintf("%d", w.Value)
}

func (w *Uint) Unary(op opcode.Op) WeakValue {
	return nil
}

func (w *Uint) Binary(op opcode.Op, v WeakValue) WeakValue {
	return nil
}
