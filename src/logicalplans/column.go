package logicalplan

import (
	"errors"

	"github.com/whiletrues/swiftseek/src/datatypes"
)

type Column struct {
	name string
}

func CreateColumn(name string) *Column {
	return &Column{
		name: name,
	}
}

func (column *Column) ToField(input LogicalPlan) (*datatypes.Field, error) {
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

func (column *Column) String() string {
	return column.name
}
