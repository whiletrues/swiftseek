package datatypes

type RecordBatch struct {
	schema *Schema
	fields []ColumnVector
}

func CreateRecord(fields []ColumnVector) RecordBatch {
	return RecordBatch{
		fields: fields,
	}
}

func (record *RecordBatch) RowCount() int {
	return 0
}

func (record *RecordBatch) ColumnCount() int {
	return 0
}

func (record *RecordBatch) GetField(index int) ColumnVector {
	return record.fields[index]
}
