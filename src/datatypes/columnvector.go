package datatypes

import (
	"github.com/apache/arrow/go/arrow"
)

type ColumnVector interface {
	GetType() arrow.DataType
	GetValue(index int) any
	Size() int
}

type LiteralValueVector struct {
	arrowType arrow.DataType
	value     any
	size      int
}

func (literalValueVector *LiteralValueVector) GetType() arrow.DataType {
	return literalValueVector.arrowType
}

func (literalValueVector *LiteralValueVector) GetValue(index int) any {
	return literalValueVector.value
}

func (literalValueVector *LiteralValueVector) Size() int {
	return literalValueVector.size
}
