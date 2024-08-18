package datasources

import "github.com/whiletrues/swiftseek/src/datatypes"

type InMemorySource struct {
	schema  *datatypes.Schema
	records *[]datatypes.RecordBatch
}

func CreateInMemorySource(schema *datatypes.Schema, records *[]datatypes.RecordBatch) *InMemorySource {
	return &InMemorySource{
		schema:  schema,
		records: records,
	}
}

func (source *InMemorySource) GetSchema() *datatypes.Schema {
	return source.schema
}

func (source *InMemorySource) Scan(projection []string) []datatypes.RecordBatch {
	return *source.records
}
