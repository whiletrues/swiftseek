package logicalexpression

import (
	"github.com/whiletrues/swiftseek/src/datatypes"
	logicalplan "github.com/whiletrues/swiftseek/src/logicalplans"
)

type LogicalExpression interface {
	toField(input logicalplan.LogicalPlan) datatypes.Field
}
