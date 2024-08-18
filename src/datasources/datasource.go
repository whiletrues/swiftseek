package datasources

import (
	"github.com/whiletrues/swiftseek/src/datatypes"
)

type DataSource interface {
	GetSchema() *datatypes.Schema
	Scan(projection []string) []datatypes.RecordBatch // Todo : Use Iter
}
