package datasources

import (
	"github.com/whiletrues/swiftseek/src/datatypes"
)

type DataSource interface {
	Schema() datatypes.Schema
	Scan() []datatypes.Record // Todo : Use Iter
}
