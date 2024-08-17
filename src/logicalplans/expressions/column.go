package logicalexpression

import (
	"errors"

	"github.com/whiletrues/swiftseek/src/datatypes"
	logicalplan "github.com/whiletrues/swiftseek/src/logicalplans"
)

type Column struct {
	name string
}

func (column *Column) ToField(input logicalplan.LogicalPlan) (*datatypes.Field, error) {
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