// This file is auto generated by /tools/AstGenerator.go

package grammar

import . "gka.com/front-end"

type Expr interface{}

type Assign struct {
	Name  Token
	Value Expr
}
type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}
type Call struct {
	Callee    Expr
	Paren     Token
	Arguments []Expr
}
type Get struct {
	Object Expr
	Name   Token
}
type Grouping struct {
	Expression Expr
}
type Literal struct {
	Value string
}
type Logical struct {
	Left     Expr
	Operator Token
	Right    Expr
}
type Set struct {
	Object Expr
	Name   Token
	Value  Expr
}
type Super struct {
	Keyword Token
	Method  Token
}
type This struct {
	Keyword Token
}
type Unary struct {
	Operator Token
	Right    Expr
}
type Variable struct {
	Name Token
}