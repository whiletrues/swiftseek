package plans

import "github.com/whiletrues/swiftseek/src/datatypes"

type LogicalPlan interface {
	GetSchema() datatypes.Schema
	GetChildrens() []LogicalPlan
}
