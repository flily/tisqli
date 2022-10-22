package weakvalue

import (
	"fmt"

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
	return 0
}

func (w *WeakString) AsUint() uint64 {
	return 0
}

func (w *WeakString) AsFloat() float64 {
	return 0
}

func (w *WeakString) AsString() string {
	return w.Value
}

func (w *WeakString) Unary(op opcode.Op) WeakValue {
	return nil
}

func (w *WeakString) Binary(op opcode.Op, v WeakValue) WeakValue {
	return nil
}
