package logicalplan

import "github.com/whiletrues/swiftseek/src/datatypes"

type Selection struct {
	input      LogicalPlan
	expression LogicalExpression
}

func CreateSelection(input LogicalPlan, expression LogicalExpression) *Selection {
	return &Selection{
		input:      input,
		expression: expression,
	}
}

func (selection *Selection) GetSchema() *datatypes.Schema {
	schema := selection.input.GetSchema()

	return &schema
}

func (selection *Selection) GetChildrens() []LogicalPlan {
	return []LogicalPlan{selection.input}
}

func (selection *Selection) String() string {
	return "Filter :> " + selection.expression.String()
}
