package datatypes

import "github.com/apache/arrow/go/arrow"

type Field struct {
	name      string
	arrowType arrow.Type
}

func CreateField(name string, arrowType arrow.Type) *Field {
	return &Field{name: name, arrowType: arrowType}
}

func (f *Field) GetName() string {
	return f.name
}

func (f *Field) GetArrowType() arrow.Type {
	return f.arrowType
}
