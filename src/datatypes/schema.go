package datatypes

type Schema struct {
	fields []Field
}

func (schema *Schema) GetFields() []Field {
	return schema.fields
}
