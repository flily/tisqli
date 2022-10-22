package weakvalue

import (
	"github.com/pingcap/tidb/parser/opcode"
)

func genericUnary(op opcode.Op, w WeakValue) WeakValue {
	switch op {
	default:
		return NewNull()
	}
}

func genericBinary(op opcode.Op, a WeakValue, b WeakValue) WeakValue {
	switch op {
	case opcode.Plus:
		return a.Add(b)
	case opcode.Minus:
		return a.Sub(b)
	default:
		return NewNull()
	}
}

func genericAdd(a WeakValue, b WeakValue) WeakValue {
	if a.IsNull() || b.IsNull() {
		return NewNull()
	}

	if a.Type() == ValueFloat || a.Type() == ValueString || b.Type() == ValueFloat || b.Type() == ValueString {
		return NewFloat(a.AsFloat() + b.AsFloat())
	}

	if a.Type() == ValueUint || b.Type() == ValueUint {
		return NewUint(a.AsUint() + b.AsUint())
	}

	// It MUST BE ValueInteger
	return NewInteger(a.AsInteger() + b.AsInteger())
}

func genericSub(a WeakValue, b WeakValue) WeakValue {
	if a.IsNull() || b.IsNull() {
		return NewNull()
	}

	if a.Type() == ValueFloat || a.Type() == ValueString || b.Type() == ValueFloat || b.Type() == ValueString {
		return NewFloat(a.AsFloat() - b.AsFloat())
	}

	if a.Type() == ValueUint || b.Type() == ValueUint {
		return NewUint(a.AsUint() - b.AsUint())
	}

	// It MUST BE ValueInteger
	return NewInteger(a.AsInteger() - b.AsInteger())
}
