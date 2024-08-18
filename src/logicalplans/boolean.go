package logicalplan

import (
	"strings"

	"github.com/apache/arrow/go/arrow"
	"github.com/whiletrues/swiftseek/src/datatypes"
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

func (expression *BooleanBinaryExpression) ToField(input LogicalPlan) (datatypes.Field, error) {
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

func (expression *BooleanBinaryExpression) String() string {
	var sb strings.Builder

	if expression.left != nil {
		sb.WriteString(expression.left.String())
	}

	sb.WriteString(" ")
	sb.WriteString(string(expression.op))

	if expression.right != nil {
		sb.WriteString(expression.right.String())
	}

	return sb.String()
}
