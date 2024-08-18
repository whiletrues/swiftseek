package logicalplan

import (
	"github.com/whiletrues/swiftseek/src/datatypes"
)

type LogicalExpression interface {
	ToField(input LogicalPlan) (datatypes.Field, error)
	String() string
}
