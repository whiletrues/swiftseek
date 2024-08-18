package logicalplan

import (
	"github.com/apache/arrow/go/arrow"
	"github.com/whiletrues/swiftseek/src/datatypes"
)

type MathBinaryExpression struct {
	op    Operator
	left  LogicalExpression
	right LogicalExpression
}

func CreateMathExpr(
	name string,
	operator Operator,
	left LogicalExpression,
	right LogicalExpression) *MathBinaryExpression {

	return &MathBinaryExpression{
		op:    operator,
		left:  left,
		right: right,
	}
}

func (expression *MathBinaryExpression) ToField(input LogicalPlan) (*datatypes.Field, error) {
	return datatypes.CreateField(string(expression.op), arrow.BOOL), nil
}

func (expression *MathBinaryExpression) GetOp() Operator {
	return expression.op
}

func (expression *MathBinaryExpression) GetLeft() LogicalExpression {
	return expression.left
}

func (expression *MathBinaryExpression) GetRight() LogicalExpression {
	return expression.right
}
