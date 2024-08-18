package logicalplan

import (
	"github.com/whiletrues/swiftseek/src/datatypes"
)

type Projection struct {
	input       LogicalPlan
	expressions []LogicalExpression
}

func CreateProjection(input LogicalPlan, expressions []LogicalExpression) *Projection {
	return &Projection{
		input:       input,
		expressions: expressions,
	}
}

func (projection *Projection) GetSchema() *datatypes.Schema {

	projectionFields := make([]datatypes.Field, 0)

	for _, expression := range projection.expressions {
		field, _ := expression.ToField(projection.input)
		projectionFields = append(projectionFields, field)
	}

	return datatypes.CreateSchema(projectionFields)
}

func (projection *Projection) GetChildrens() []LogicalPlan {
	return []LogicalPlan{projection.input}
}

func (projection *Projection) String() string {

	result := "Projection :> ["
	for _, expression := range projection.expressions {
		result += expression.String() + ", "
	}

	return "Projection :> [ " + result + " ]"
}
