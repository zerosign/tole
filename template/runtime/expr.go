package runtime

import (
	"github.com/aymerick/raymond/ast"
)

type Expressions struct {
	stacks []*ast.Expression
	heap   map[*ast.Expression]bool
}
