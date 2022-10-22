package weakvalue

import (
	"fmt"
	"strconv"

	"github.com/pingcap/tidb/parser/opcode"
	"github.com/pingcap/tidb/util/sqlexec"
)

type WeakString struct {
	Value string
}

func NewString(s string) WeakValue {
	w := &WeakString{
		Value: s,
	}
	return w
}

func (w *WeakString) Type() Type {
	return ValueString
}

func (w *WeakString) EqualTo(v WeakValue) bool {
	return w.Value == v.AsString()
}

func (w *WeakString) String() string {
	return fmt.Sprintf("'%s'", sqlexec.EscapeString(w.Value))
}

func (w *WeakString) IsNull() bool {
	return false
}

func (w *WeakString) AsBoolean() bool {
	return w.Value != ""
}

func (w *WeakString) AsInteger() int64 {
	i, _ := w.ToInteger()
	return i
}

func (w *WeakString) AsUint() uint64 {
	u, _ := w.ToUint()
	return u
}

func (w *WeakString) AsFloat() float64 {
	f, _ := w.ToFloat()
	return f
}

func (w *WeakString) AsString() string {
	return w.Value
}

func (w *WeakString) Unary(op opcode.Op) WeakValue {
	return genericUnary(op, w)
}

func (w *WeakString) Binary(op opcode.Op, v WeakValue) WeakValue {
	return genericBinary(op, w, v)
}

func (w *WeakString) ToInteger() (int64, bool) {
	i, err := strconv.ParseInt(w.Value, 10, 64)
	if err != nil {
		return 0, false
	}

	return i, true
}

func (w *WeakString) ToUint() (uint64, bool) {
	u, err := strconv.ParseUint(w.Value, 10, 64)
	if err != nil {
		return 0, false
	}

	return u, true
}

func (w *WeakString) ToFloat() (float64, bool) {
	f, err := strconv.ParseFloat(w.Value, 64)
	if err != nil {
		return 0.0, false
	}

	return f, true
}

func (w *WeakString) Add(v WeakValue) WeakValue {
	return genericAdd(w, v)
}

func (w *WeakString) Sub(v WeakValue) WeakValue {
	return genericSub(w, v)
}
