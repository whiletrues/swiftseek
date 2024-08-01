package datatypes

import (
	"github.com/apache/arrow/go/arrow"
)

type ColumnVector interface {
	GetType() arrow.DataType
	GetValue(index int) any
	Size() int
}
