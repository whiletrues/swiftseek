package datatypes

type Schema struct {
	fields []Field
}

func CreateSchema(fields []Field) *Schema {
	return &Schema{
		fields: fields,
	}
}

func (schema *Schema) GetFields() []Field {
	return schema.fields
}
