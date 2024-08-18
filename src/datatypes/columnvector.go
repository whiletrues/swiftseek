package datatypes

import (
	"github.com/apache/arrow/go/arrow"
)

type ColumnVector interface {
	GetType() arrow.Type
	GetValue(index int) any
	Size() int
}
