package plans

import (
	"errors"

	"github.com/whiletrues/swiftseek/src/datatypes"
	plans "github.com/whiletrues/swiftseek/src/plans/logicalplans"
)

type Column struct {
	name string
}

func (column *Column) ToField(input plans.LogicalPlan) (*datatypes.Field, error) {
	schema := input.GetSchema()

	fields := schema.GetFields()

	for _, field := range fields {
		if field.GetName() == column.name {
			return &field, nil
		}
	}

	return nil, errors.New("field not found")
}

func (column *Column) GetName() string {
	return column.name
}
