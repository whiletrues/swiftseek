package datatypes

type Record struct {
	schema Schema
	fields []ColumnVector
}

func (record *Record) RowCount() int {
	return 0
}

func (record *Record) ColumnCount() int {
	return 0
}

func (record *Record) GetField(index int) ColumnVector {
	return record.fields[index]
}
