package template

import (
	"github.com/aymerick/raymond/ast"
	"github.com/zerosign/tole/base"
)

// INFO: Most of the newtypes being created since
// sometimes using reflect.DeepEquals returns
// unexpected results.

// Operation : newtype for operation (function call)
//
type Operation []string

// Equal : check equality of Operation
//
// This method check equality by using `==` to both elements
//
func (o Operation) Equal(other Operation) bool {
	var result bool = true

	if len(o) != len(other) {
		return false
	}

	for ii := 0; ii < len(other); ii++ {
		result = result && (o[ii] == other[ii])
	}

	return result
}

// Operations : newtype for list of operation
//
type Operations []Operation

// Equal : check equality of Operations
//
// This method check equality by using `Equal` method to both elements
//
func (oo Operations) Equal(other Operations) bool {
	var result bool = true

	if len(oo) != len(other) {
		return false
	}

	for ii := 0; ii < len(other); ii++ {
		result = result && (oo[ii].Equal(other[ii]))
	}

	return result
}

// SourcePaths : newtype for string map of operations
//
//
type SourcePaths map[string]Operations

// Equal : check equality of SourcePaths
//
// This method check equality by using `Equal` method to both elements
//
func (s SourcePaths) Equal(other SourcePaths) bool {
	var result bool = true

	if len(s) != len(other) {
		return false
	}

	for key, op := range s {
		result = result && (op.Equal(other[key]))
	}

	return result
}

// Populator : populator for populating lookup paths in template
//
// Use it with ast.Node#Accept method and returns the
// lists of lookup paths by using Populator#lists. This struct
// are being used as once time call only since it holds states of
// SourcePaths.
//
// Technically, we only need interest for any function call statements.
// So, we need to visits(expands) at least :
// - ast.Program
// - ast.MustacheStatement
// - ast.BlockStatement
// - ast.PartialStatement
// - ast.Expression (since function call are expression)
// - ast.SubExpression
// - ast.Hash
// -
type Populator struct {
	functions base.StrSet
	inner     SourcePaths
}

// NewPopulator : create new populator by giving valid lookup func str.
//
//
func NewPopulator(validLookupFuncStr []string) Populator {
	return Populator{base.MakeStrSet(validLookupFuncStr), SourcePaths{}}
}

// Lists : list all lookup paths of Populator
//
//
func (p Populator) Lists() SourcePaths {
	return p.inner
}

//
// ast.Node Visitor section
//

// VisitProgram : visit program
//
// Actual entrypoint for entire ast.Node
//
func (p Populator) VisitProgram(node *ast.Program) interface{} {

	for _, n := range node.Body {
		n.Accept(p)
	}

	return nil
}

// VisitMustache : method that being called when Populator visit ast.MustacheStatement
//
// Since function calls could exists in a mustache statements too
//
func (p Populator) VisitMustache(node *ast.MustacheStatement) interface{} {

	node.Expression.Accept(p)

	return nil
}

// VisitBlock : method that being called when Populator visit ast.BlockStatement
//
// Since function calls could exists in a block too
//
func (p Populator) VisitBlock(node *ast.BlockStatement) interface{} {

	node.Expression.Accept(p)

	if node.Program != nil {
		node.Program.Accept(p)
	}

	if node.Inverse != nil {
		node.Inverse.Accept(p)
	}

	return nil
}

// VisitPartial : method that being called when Populator visit ast.PartialStatement
//
// Since function calls could exists in partial template too
//
func (p Populator) VisitPartial(node *ast.PartialStatement) interface{} {

	node.Name.Accept(p)

	if len(node.Params) > 0 {
		node.Params[0].Accept(p)
	}

	if node.Hash != nil {
		node.Hash.Accept(p)
	}

	return nil
}

// VisitExpression : method that being called when Populator visit ast.Expression
//
// Since function calls are basically one of expression (ast.Expression),
// this method are important for getting the actual information of lookup paths.
//
func (p Populator) VisitExpression(node *ast.Expression) interface{} {

	node.Path.Accept(p)

	if len(node.Params) == 3 {

		if funcNode, ok := node.Params[0].(*ast.PathExpression); ok {
			var functionName, sourceName, queryPath string

			//
			if p.functions.Contains(funcNode.Original) {
				functionName = funcNode.Original
			}

			if sourceNode, ok := node.Params[1].(*ast.StringLiteral); ok {
				sourceName = sourceNode.Value
			}

			if queryNode, ok := node.Params[2].(*ast.StringLiteral); ok {
				queryPath = queryNode.Value
			}

			if len(functionName) != 0 && len(sourceName) != 0 && len(queryPath) != 0 {
				var operations []Operation
				var ok bool

				if operations, ok = p.inner[sourceName]; ok {
					operations = append(operations, []string{functionName, queryPath})
				} else {
					operations = []Operation{}
					operations = append(operations, []string{functionName, queryPath})
				}

				p.inner[sourceName] = operations
			}
		}

	}

	for _, n := range node.Params {
		n.Accept(p)
	}

	if node.Hash != nil {
		node.Hash.Accept(p)
	}

	return nil
}

// VisitSubExpression : method that being called when Populator visit ast.SubExpression
//
func (p Populator) VisitSubExpression(node *ast.SubExpression) interface{} {
	node.Expression.Accept(p)
	return nil
}

// VisitHash : method being called when populator visits ast.Hash
//
// we do nothing in here since we don't need any information for ast.Hash
func (p Populator) VisitHash(node *ast.Hash) interface{} { return nil }

// VisitHashPair : method being called when populator visits ast.HashPair
//
// we do nothing in here since we don't need any information for ast.HashPair
func (p Populator) VisitHashPair(node *ast.HashPair) interface{} { return nil }

// VisitContent : method being called when populator visits ast.ContentStatement
//
// we do nothing in here since we don't need any information for ast.ContentStatement
func (p Populator) VisitContent(node *ast.ContentStatement) interface{} { return nil }

// VisitComment : method being called when populator visits ast.CommentStatement
//
// we do nothing in here since we don't need any information for ast.CommentStatement
func (p Populator) VisitComment(node *ast.CommentStatement) interface{} { return nil }

// VisitPath : method being called when populator visits ast.PathExpression
//
// we do nothing in here since we don't need any information for ast.PathExpression
func (p Populator) VisitPath(node *ast.PathExpression) interface{} { return nil }

// VisitString : method being called when populator visits ast.StringLiteral
//
// we do nothing in here since we don't need any information for ast.StringLiteral
func (p Populator) VisitString(node *ast.StringLiteral) interface{} { return nil }

// VisitBoolean : method being called when populator visits ast.BooleanLiteral
//
// we do nothing in here since we don't need any information for ast.BooleanLiteral
func (p Populator) VisitBoolean(node *ast.BooleanLiteral) interface{} { return nil }

// VisitNumber : method being called when populator visits ast.NumberLiteral
//
// we do nothing in here since we don't need any information for ast.NumberLiteral
func (p Populator) VisitNumber(node *ast.NumberLiteral) interface{} { return nil }
