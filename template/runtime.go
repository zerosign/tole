package template

import (
	"github.com/zerosign/tole/template/runtime"
	"reflect"
)

// Runtime : represents runtime stack
//
// we use slice as reversed stack
// root point to index 0, top of the stack are
// index len(inner)-1
//
type Runtime struct {
	inner []reflect.Value
	// block runtime (stack & its indexed heap)
	block runtime.Blocks
	// expr runtime (stack & its indexed heap)
	expr runtime.Expressions
}

func (r Runtime) Push(value reflect.Value) {
	r.inner = append(r.inner, value)
}

// Pop : Pop current stack
//
//
func (r Runtime) Pop() reflect.Value {
	if len(r.inner) == 0 {
		return zero
	}

	var result reflect.Value

	// returns top of the stack
	// and returns others element of the stack
	result, r.inner = r.inner[len(r.inner)-1], r.inner[:len(r.inner)-1]

	return result
}

//
//
func (r Runtime) Root() reflect.Value {
	return r.inner[0]
}
