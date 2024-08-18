package logicalplan

import (
	"strings"

	"github.com/whiletrues/swiftseek/src/datatypes"
)

type AggregateOperator string

const (
	SUM   AggregateOperator = "sum"
	COUNT AggregateOperator = "count"
	MIN   AggregateOperator = "min"
	MAX   AggregateOperator = "max"
	AVG   AggregateOperator = "avg"
)

type AggregateExpression struct {
	name     string
	operator AggregateOperator
	expr     LogicalExpression
}

func CreateAggregateExpr(
	name string,
	expr LogicalExpression) *AggregateExpression {

	return &AggregateExpression{
		name: name,
		expr: expr,
	}
}

func (aggregateExpr *AggregateExpression) GetName() string {
	return aggregateExpr.name
}

func (aggregateExpr *AggregateExpression) GetOperator() AggregateOperator {
	return aggregateExpr.operator
}

func (aggregateExpr *AggregateExpression) ToField(input LogicalPlan) (datatypes.Field, error) {
	field, _ := aggregateExpr.expr.ToField(input)
	return datatypes.CreateField(aggregateExpr.name, field.GetArrowType()), nil
}

func (aggregateExpr *AggregateExpression) String() string {

	var sb strings.Builder

	sb.WriteString("(")
	sb.WriteString(string(aggregateExpr.name))
	sb.WriteString(" ")
	sb.WriteString(string(aggregateExpr.operator))
	if aggregateExpr.expr != nil {
		sb.WriteString(aggregateExpr.expr.String())
	}

	sb.WriteString(")")

	return sb.String()
}
