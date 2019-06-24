package template

// import (
// 	"bytes"
// 	"github.com/aymerick/raymond/ast"
// 	"reflect"
// )

// TODO: @zerosign
//
// I need to comment the evaluator first, since
// it seems that modifying & refactoring the flow
// won't be that much easy.
//
// var (
// 	emptyBlockParams = make(BlockParams, 0) // Heap<String, BlockParams>
// 	emptyBlocks      = make(Blocks, 0)      // Stack<BlockStatement>
// 	emptyExpressions = make(Expressions, 0) // Stack<Expression>
// 	emptyExprFunc    = make(ExprFunc)       // Heap<Expression, bool>
// )

// type Evaluator struct {
// 	blockParams BlockParams
// 	blocks      Blocks
// 	exprs       Expressions
// 	exprFunc    ExprFunc
// 	currentNode ast.Node
// 	runtime     Runtime
// 	functions   *FunctionTables
// }

// func NewEvaluator(program *ast.Program, functions *FunctionTables) Evaluator {
// 	return Evaluator{emptyBlockParams, emptyBlocks, emptyExpressions, emptyExprFunc, program, functions}
// }

// func EmptyEvaluator(program *ast.Program) Evaluator {
// 	return Evaluator(emptyBlockParams, emptyBlocks, emptyExpressions, emptyExprFunc, program, EmptyTables())
// }

// func (v Evaluator) setNode(node ast.Node) {
// 	v.currentNode = node
// }

// //
// // Visitor interface
// //

// // Statements

// // VisitProgram implements corresponding Visitor interface method
// func (v Evaluator) VisitProgram(node *ast.Program) interface{} {
// 	v.at(node)

// 	buf := new(bytes.Buffer)

// 	for _, n := range node.Body {
// 		if str := Str(n.Accept(v)); str != "" {
// 			if _, err := buf.Write([]byte(str)); err != nil {
// 				log.Fatal("Evaluation error: %s\nCurrent node:\n\t%s", err, v.currentNode)
// 			}
// 		}
// 	}

// 	return buf.String()
// }

// // VisitMustache implements corresponding Visitor interface method
// func (v Evaluator) VisitMustache(node *ast.MustacheStatement) interface{} {
// 	v.at(node)

// 	// evaluate expression
// 	expr := node.Expression.Accept(v)

// 	// check if this is a safe string
// 	isSafe := isSafeString(expr)

// 	// get string value
// 	str := Str(expr)
// 	if !isSafe && !node.Unescaped {
// 		// escape html
// 		str = Escape(str)
// 	}

// 	return str
// }

// // VisitBlock implements corresponding Visitor interface method
// func (v Evaluator) VisitBlock(node *ast.BlockStatement) interface{} {
// 	v.at(node)

// 	v.pushBlock(node)

// 	var result interface{}

// 	// evaluate expression
// 	expr := node.Expression.Accept(v)

// 	if v.isHelperCall(node.Expression) || v.wasFuncCall(node.Expression) {
// 		// it is the responsibility of the helper/function to evaluate block
// 		result = expr
// 	} else {
// 		val := reflect.ValueOf(expr)

// 		truth, _ := isTrueValue(val)
// 		if truth {
// 			if node.Program != nil {
// 				switch val.Kind() {
// 				case reflect.Array, reflect.Slice:
// 					concat := ""

// 					// Array context
// 					for i := 0; i < val.Len(); i++ {
// 						// Computes new private data frame
// 						frame := v.dataFrame.newIterDataFrame(val.Len(), i, nil)

// 						// Evaluate program
// 						concat += v.evalProgram(node.Program, val.Index(i).Interface(), frame, i)
// 					}

// 					result = concat
// 				default:
// 					// NOT array
// 					result = v.evalProgram(node.Program, expr, nil, nil)
// 				}
// 			}
// 		} else if node.Inverse != nil {
// 			result, _ = node.Inverse.Accept(v).(string)
// 		}
// 	}

// 	v.popBlock()

// 	return result
// }

// // VisitPartial implements corresponding Visitor interface method
// func (v Evaluator) VisitPartial(node *ast.PartialStatement) interface{} {
// 	v.at(node)

// 	// partialName: helperName | sexpr
// 	name, ok := ast.HelperNameStr(node.Name)
// 	if !ok {
// 		if subExpr, ok := node.Name.(*ast.SubExpression); ok {
// 			name, _ = subExpr.Accept(v).(string)
// 		}
// 	}

// 	if name == "" {
// 		log.Fatalf("Unexpected partial name: %q", node.Name)
// 	}

// 	partial := v.findPartial(name)
// 	if partial == nil {
// 		log.Fatalf("Partial not found: %s", name)
// 	}

// 	return v.evalPartial(partial, node)
// }

// // VisitContent implements corresponding Visitor interface method
// func (v Evaluator) VisitContent(node *ast.ContentStatement) interface{} {
// 	v.at(node)

// 	// write content as is
// 	return node.Value
// }

// // VisitComment implements corresponding Visitor interface method
// func (v Evaluator) VisitComment(node *ast.CommentStatement) interface{} {
// 	v.at(node)

// 	// ignore comments
// 	return ""
// }

// // Expressions

// // VisitExpression implements corresponding Visitor interface method
// func (v Evaluator) VisitExpression(node *ast.Expression) interface{} {
// 	v.at(node)

// 	var result interface{}
// 	done := false

// 	v.pushExpr(node)

// 	// helper call
// 	if helperName := node.HelperName(); len(helperName) != 0 {

// 		if helper, ok := v.helpers.Get(helperName); ok {

// 		}

// 		if helper := v.findHelper(helperName); helper != zero {
// 			result = v.callHelper(helperName, helper, node)
// 			done = true
// 		}
// 	}

// 	if !done {
// 		// literal
// 		if literal, ok := node.LiteralStr(); ok {
// 			if val := v.evalField(v.curCtx(), literal, true); val.IsValid() {
// 				result = val.Interface()
// 				done = true
// 			}
// 		}
// 	}

// 	if !done {
// 		// field path
// 		if path := node.FieldPath(); path != nil {
// 			// @todo Find a cleaner way ! Don't break the pattern !
// 			// this is an exception to visitor pattern, because we need to pass the info
// 			// that this path is at root of current expression
// 			if val := v.evalPathExpression(path, true); val != nil {
// 				result = val
// 			}
// 		}
// 	}

// 	v.popExpr()

// 	return result
// }

// // VisitSubExpression implements corresponding Visitor interface method
// func (v Evaluator) VisitSubExpression(node *ast.SubExpression) interface{} {
// 	v.at(node)

// 	return node.Expression.Accept(v)
// }

// // VisitPath implements corresponding Visitor interface method
// func (v Evaluator) VisitPath(node *ast.PathExpression) interface{} {
// 	return v.evalPathExpression(node, false)
// }

// // Literals

// // VisitString implements corresponding Visitor interface method
// func (v Evaluator) VisitString(node *ast.StringLiteral) interface{} {
// 	v.at(node)

// 	return node.Value
// }

// // VisitBoolean implements corresponding Visitor interface method
// func (v Evaluator) VisitBoolean(node *ast.BooleanLiteral) interface{} {
// 	v.at(node)

// 	return node.Value
// }

// // VisitNumber implements corresponding Visitor interface method
// func (v Evaluator) VisitNumber(node *ast.NumberLiteral) interface{} {
// 	v.at(node)

// 	return node.Number()
// }

// // Miscellaneous

// // VisitHash implements corresponding Visitor interface method
// func (v Evaluator) VisitHash(node *ast.Hash) interface{} {
// 	v.at(node)

// 	result := make(map[string]interface{})

// 	for _, pair := range node.Pairs {
// 		if value := pair.Accept(v); value != nil {
// 			result[pair.Key] = value
// 		}
// 	}

// 	return result
// }

// // VisitHashPair implements corresponding Visitor interface method
// func (v Evaluator) VisitHashPair(node *ast.HashPair) interface{} {
// 	v.at(node)

// 	return node.Val.Accept(v)
// }
