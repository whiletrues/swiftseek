package datatypes

import "github.com/apache/arrow/go/arrow"

type LiteralValueVector struct {
	arrowType arrow.Type
	value     any
	size      int
}

func CreateLiteralValueVector(
	arrowType arrow.Type,
	value any,
	size int) *LiteralValueVector {

	return &LiteralValueVector{
		arrowType: arrowType,
		value:     value,
		size:      size,
	}
}

func (literalValueVector *LiteralValueVector) GetType() arrow.Type {
	return literalValueVector.arrowType
}

func (literalValueVector *LiteralValueVector) GetValue(index int) any {
	return literalValueVector.value
}

func (literalValueVector *LiteralValueVector) Size() int {
	return literalValueVector.size
}
