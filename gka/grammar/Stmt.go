// This file is auto generated by /tools/AstGenerator.go

package grammar

import . "gka.com/front-end"

type Stmt interface{}

type Block struct {
	Statements []Stmt
}
type Class struct {
	Name       Token
	Superclass Variable
	Methods    []Function
}
type Expression struct {
	Expression Expr
}
type Function struct {
	Name   Token
	Params []Token
	Body   []Stmt
}
type If struct {
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}
type Print struct {
	Expression Expr
}
type Return struct {
	Keyword Token
	Value   Expr
}
type Var struct {
	Name        Token
	Initializer Expr
}
type While struct {
	Condition Expr
	Body      Stmt
}
