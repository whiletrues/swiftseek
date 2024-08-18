package logicalplan

import (
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

func (aggregateExpr *AggregateExpression) ToField(input LogicalPlan) (*datatypes.Field, error) {
	field := aggregateExpr.expr.toField(input)
	return datatypes.CreateField(aggregateExpr.name, field.GetArrowType()), nil
}

func (aggregateExpr *AggregateExpression) String() string {
	return string(aggregateExpr.operator) + "(" + aggregateExpr.expr.String() + ")"
}
