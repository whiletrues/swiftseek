package logicalplan

type Operator string

const (
	// Comparison operators
	Eq   Operator = "eq"
	Neq  Operator = "neq"
	Gt   Operator = "gt"
	Lt   Operator = "lt"
	GtEq Operator = "ge"
	LeEq Operator = "le"

	// Logical operators
	And Operator = "and"
	Or  Operator = "or"

	// Math operators
	Add Operator = "add"
	Sub Operator = "sub"
	Mul Operator = "mul"
	Div Operator = "div"
	Mod Operator = "mod"
)

type BinaryExpression interface {
	GetOperator() Operator
	GetLeft() LogicalExpression
	GetRight() LogicalExpression
	String() string
}
