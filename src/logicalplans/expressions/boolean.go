package logicalexpression

import (
	"github.com/apache/arrow/go/arrow"
	"github.com/whiletrues/swiftseek/src/datatypes"
	logicalplan "github.com/whiletrues/swiftseek/src/logicalplans"
)

type BooleanBinaryExpression struct {
	op    Operator
	left  LogicalExpression
	right LogicalExpression
}

func CreateBooleanExpr(
	name string,
	operator Operator,
	left LogicalExpression,
	right LogicalExpression) *BooleanBinaryExpression {

	return &BooleanBinaryExpression{
		op:    operator,
		left:  left,
		right: right,
	}
}

func (expression *BooleanBinaryExpression) ToField(input logicalplan.LogicalPlan) (*datatypes.Field, error) {
	return datatypes.CreateField(string(expression.op), arrow.BOOL), nil
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
