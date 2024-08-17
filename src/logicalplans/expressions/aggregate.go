package logicalexpression

import (
	"github.com/whiletrues/swiftseek/src/datatypes"
	logicalplan "github.com/whiletrues/swiftseek/src/logicalplans"
)

type AggregateExpression struct {
	name string
	expr LogicalExpression
}

func CreateAggregateExpr(
	name string,
	expr LogicalExpression) *AggregateExpression {

	return &AggregateExpression{
		name: name,
		expr: expr,
	}
}
func (aggregateExpr *AggregateExpression) ToField(input logicalplan.LogicalPlan) (*datatypes.Field, error) {
	field := aggregateExpr.expr.toField(input)
	return datatypes.CreateField(aggregateExpr.name, field.GetArrowType()), nil
}
