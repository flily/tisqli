package weakvalue

import (
	"github.com/pingcap/tidb/parser/opcode"
)

type Type int

const (
	ValueNull    Type = 0
	ValueInteger Type = 1
	ValueUint    Type = 2
	ValueFloat   Type = 3
	ValueString  Type = 4
)

// WeakValue is a value of weak type
type WeakValue interface {
	Type() Type
	EqualTo(WeakValue) bool
	String() string

	IsNull() bool
	AsBoolean() bool
	AsInteger() int64
	AsUint() uint64
	AsFloat() float64
	AsString() string

	Unary(opcode.Op) WeakValue
	Binary(opcode.Op, WeakValue) WeakValue
}
