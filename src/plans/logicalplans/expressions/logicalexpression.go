package plans

import plans "github.com/whiletrues/swiftseek/src/plans/logicalplans"

type LogicalExpression interface {
	toField(input plans.LogicalPlan)
}
