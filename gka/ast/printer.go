package ast

func parenWrap(operator []rune, exprs ...Expr) []rune {
	runes := append([]rune{'('}, operator...)

	for _, expr := range exprs {
		runes = append(runes, expr.Print()...)
	}
	return append(runes, ')')
}

func parenWrap2(operator []rune, parts ...interface{}) []rune {
	runes := append([]rune{'('}, operator...)

	for _, part := range parts {
		switch part.(type) {
		case Expr:
			runes = append(runes, part.(Expr).Print()...)
		case Stmt:
			runes = append(runes, part.(Stmt).Print()...)
		case []rune:
			runes = append(runes, part.([]rune)...)
		}
	}
	return append(runes, ')')
}

// *** EXPRESSIONS *** //
func (expr *AssignExpr) Print() []rune {
	return parenWrap2([]rune{'='}, expr.Name.Lexeme, expr.Value)
}

func (expr *BinaryExpr) Print() []rune {
	return parenWrap2(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (expr *CallExpr) Print() []rune {
	return parenWrap2([]rune("call"), expr.Callee, expr.Arguments)
}

func (expr *GetExpr) Print() []rune {
	return parenWrap2([]rune("."), expr.Object, expr.Name.Lexeme)
}

func (expr *GroupingExpr) Print() []rune {
	return parenWrap2([]rune("group"), expr.Expression)
}

func (expr *LiteralExpr) Print() []rune {
	if expr.Value == "" {
		return []rune("nil")
	}

	return []rune(expr.Value)
}

func (expr *LogicalExpr) Print() []rune {
	return parenWrap(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (expr *SetExpr) Print() []rune {
	return parenWrap2([]rune{'='}, expr.Object, expr.Name.Lexeme, expr.Value)
}

func (expr *SuperExpr) Print() []rune {
	return parenWrap2([]rune("super"), expr.Method)
}

func (expr *ThisExpr) Print() []rune {
	return []rune("this")
}

func (expr *UnaryExpr) Print() []rune {
	return parenWrap(expr.Operator.Lexeme, expr.Right)
}

func (expr *VariableExpr) Print() []rune {
	return expr.Name.Lexeme
}

// *** STATEMENTS *** //
func (stmt *BlockStmt) Print() []rune {
	return parenWrap2([]rune("block"), stmt.Statements)
}

func (stmt *ClassStmt) Print() []rune {
	runes := append([]rune("(class "), stmt.Name.Lexeme...)
	if len(stmt.Superclass.Name.Lexeme) > 0 {
		runes = append(runes, []rune(" < ")...)
		runes = append(runes, stmt.Superclass.Print()...)
	}

	for _, method := range stmt.Methods {
		runes = append(runes, method.Print()...)
	}

	runes = append(runes, ')')

	return runes
}

func (stmt *ExpressionStmt) Print() []rune {
	return parenWrap([]rune(";"), stmt.Expression)
}

func (stmt *FunctionStmt) Print() []rune {
	runes := append([]rune("(fn "), stmt.Name.Lexeme...)
	runes = append(runes, '(')

	// Print function parameters
	for _, param := range stmt.Params {
		runes = append(runes, param.Lexeme...)
	}
	runes = append(runes, ')')

	// Print function body
	for _, statement := range stmt.Body {
		runes = append(runes, statement.Print()...)
	}
	runes = append(runes, ')')

	return runes
}

func (stmt *IfStmt) Print() []rune {
	if stmt.ElseBranch == nil {
		return parenWrap2([]rune("if"), stmt.Condition, stmt.ThenBranch)
	}
	return parenWrap2([]rune("if-else"), stmt.Condition, stmt.ThenBranch, stmt.ElseBranch)
}

func (stmt *PrintStmt) Print() []rune {
	return parenWrap([]rune("print"), stmt.Expression)
}

func (stmt *ReturnStmt) Print() []rune {
	if stmt.Value == nil {
		return []rune("(return)")
	}
	return parenWrap([]rune("return"), stmt.Value)

}

func (stmt *VarStmt) Print() []rune {
	if stmt.Initializer == nil {
		return parenWrap2([]rune("var"), stmt.Name.Lexeme)
	}

	return parenWrap2([]rune("var"), stmt.Name, '=', stmt.Initializer)
}

func (stmt *WhileStmt) Print() []rune {
	return parenWrap2([]rune("while"), stmt.Condition, stmt.Body)
}
