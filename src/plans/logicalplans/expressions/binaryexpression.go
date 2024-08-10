package plans

import (
	"github.com/apache/arrow/go/arrow"
	"github.com/whiletrues/swiftseek/src/datatypes"
	plans "github.com/whiletrues/swiftseek/src/plans/logicalplans"
)

type Operator string

const (
	Eq   Operator = "eq"
	Neq  Operator = "neq"
	Gt   Operator = "gt"
	Lt   Operator = "lt"
	GtEq Operator = "ge"
	LeEq Operator = "le"
)

type BinaryExpression interface {
	GetName() string
	GetOperator() Operator
	GetLeft() LogicalExpression
	GetRight() LogicalExpression
}

type BooleanBinaryExpression struct {
	name  string
	op    Operator
	left  LogicalExpression
	right LogicalExpression
}

func (expression *BooleanBinaryExpression) ToField(input plans.LogicalPlan) (*datatypes.Field, error) {
	return datatypes.CreateField(expression.name, arrow.BOOL), nil
}

func (expression *BooleanBinaryExpression) GetName() string {
	return expression.name
}

func (expression *BooleanBinaryExpression) GetOp() Operator {
	return expression.op
}

func (expression *BooleanBinaryExpression) GetLeft() LogicalExpression {
	return expression.left
}

func (expression *BooleanBinaryExpression) GetRight() LogicalExpression {
	return expression.right
}
