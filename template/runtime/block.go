package runtime

import (
	"github.com/aymerick/raymond/ast"
)

type Blocks struct {
	heaps  []map[string]interface{}
	stacks []*ast.BlockStatement
}
