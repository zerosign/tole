package template

import (
	"github.com/aymerick/raymond/ast"
	"github.com/zerosign/tole/base"
)

type Operation []string

type SourcePaths map[string][]Operation

type Populator struct {
	functions base.StrSet
	inner     SourcePaths
}

func NewPopulator() Populator {
	return Populator{base.EmptyStrSet(), SourcePaths{}}
}

func (p Populator) Lists() SourcePaths {
	return p.inner
}

func (p Populator) VisitProgram(node *ast.Program) interface{} {

	for _, n := range node.Body {
		n.Accept(p)
	}

	return nil
}

func (p Populator) VisitMustache(node *ast.MustacheStatement) interface{} {

	node.Expression.Accept(p)

	return nil
}

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

// expressions
func (p Populator) VisitExpression(node *ast.Expression) interface{} {

	node.Path.Accept(p)

	if len(node.Params) == 3 {

		if funcNode, ok := node.Params[0].(*ast.PathExpression); ok {
			var functionName, sourceName, queryPath string

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

func (p Populator) VisitSubExpression(node *ast.SubExpression) interface{} {
	node.Expression.Accept(p)
	return nil
}

// miscellaneous
func (p Populator) VisitHash(node *ast.Hash) interface{} {

	for _, v := range node.Pairs {
		v.Accept(p)
	}

	return nil
}

func (p Populator) VisitHashPair(node *ast.HashPair) interface{} {
	node.Val.Accept(p)

	return nil
}

func (p Populator) VisitContent(node *ast.ContentStatement) interface{} { return nil }
func (p Populator) VisitComment(node *ast.CommentStatement) interface{} { return nil }
func (p Populator) VisitPath(node *ast.PathExpression) interface{}      { return nil }

// literals
func (p Populator) VisitString(node *ast.StringLiteral) interface{}   { return nil }
func (p Populator) VisitBoolean(node *ast.BooleanLiteral) interface{} { return nil }
func (p Populator) VisitNumber(node *ast.NumberLiteral) interface{}   { return nil }
