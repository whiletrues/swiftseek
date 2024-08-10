package plans

import (
	"strconv"

	"github.com/apache/arrow/go/arrow"
	datatypes "github.com/whiletrues/swiftseek/src/datatypes"
	plans "github.com/whiletrues/swiftseek/src/plans/logicalplans"
)

/*** LiteralString  ***/

type LiteralString struct {
	value string
}

func (literal *LiteralString) ToField(input plans.LogicalPlan) (*datatypes.Field, error) {
	return datatypes.CreateField(literal.value, arrow.STRING), nil
}

/*** LiteralLong ***/

type LiteralLong struct {
	value int64
}

func (literal *LiteralLong) ToField(input plans.LogicalPlan) (*datatypes.Field, error) {

	name := strconv.FormatInt(literal.value, 10)

	return datatypes.CreateField(name, arrow.INT64), nil
}
