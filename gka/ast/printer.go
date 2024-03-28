package ast

import "fmt"

type printer struct {
}

// *** EXPRESSIONS *** //
func (expr *AssignExpr) Print() string {
	return fmt.Sprintf("(Assign %v %v)", expr.Name.Lexeme, expr.Value.Print())
}

func (expr *BinaryExpr) Print() string {
	return fmt.Sprintf("(Binary %v %v %v)", expr.Operator.Lexeme, expr.Left.Print(), expr.Right.Print())
}

func (expr *CallExpr) Print() string {
	args := ""
	for _, arg := range expr.Arguments {
		args += " " + arg.Print()
	}
	return fmt.Sprintf("(Call %v%s)", expr.Callee.Print(), args)
}

func (expr *GetExpr) Print() string {
	return ""
}

func (expr *GroupingExpr) Print() string {
	return ""
}

func (expr *LiteralExpr) Print() string {
	return ""
}

func (expr *LogicalExpr) Print() string {
	return ""
}

func (expr *SetExpr) Print() string {
	return ""
}

func (expr *SuperExpr) Print() string {
	return ""
}

func (expr *ThisExpr) Print() string {
	return ""
}

func (expr *UnaryExpr) Print() string {
	return ""
}

func (expr *VariableExpr) Print() string {
	return ""
}

// *** STATEMENTS *** //
func (stmt *BlockStmt) Print() string {
	return ""
}

func (stmt *ClassStmt) Print() string {
	return ""
}

func (stmt *ExpressionStmt) Print() string {
	return ""
}

func (stmt *FunctionStmt) Print() string {
	return ""
}

func (stmt *IfStmt) Print() string {
	return ""
}

func (stmt *PrintStmt) Print() string {
	return ""
}

func (stmt *ReturnStmt) Print() string {
	return ""
}

func (stmt *VarStmt) Print() string {
	return ""
}

func (stmt *WhileStmt) Print() string {
	return ""
}
