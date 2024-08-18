package logicalplan

import (
	"github.com/whiletrues/swiftseek/src/datatypes"
)

type LogicalExpression interface {
	toField(input LogicalPlan) datatypes.Field
	String() string
}
