package weakvalue

import "github.com/pingcap/tidb/parser/opcode"

// WeakNull represents a null value.
type WeakNull struct {
}

func NewNull() WeakValue {
	return &WeakNull{}
}

func (w *WeakNull) Type() Type {
	return ValueNull
}

func (w *WeakNull) EqualTo(v WeakValue) bool {
	return v.IsNull()
}

func (w *WeakNull) String() string {
	return "NULL"
}

func (w *WeakNull) IsNull() bool {
	return true
}

func (w *WeakNull) AsBoolean() bool {
	return false
}

func (w *WeakNull) AsInteger() int64 {
	return 0
}

func (w *WeakNull) AsUint() uint64 {
	return 0
}

func (w *WeakNull) AsFloat() float64 {
	return 0
}

func (w *WeakNull) AsString() string {
	return ""
}

func (w *WeakNull) Unary(op opcode.Op) WeakValue {
	return nil
}

func (w *WeakNull) Binary(op opcode.Op, v WeakValue) WeakValue {
	return nil
}
