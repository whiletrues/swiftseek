package logicalplan

import (
	"github.com/whiletrues/swiftseek/src/datasources"
	"github.com/whiletrues/swiftseek/src/datatypes"
)

type Scan struct {
	path       string
	source     *datasources.DataSource
	projection []string
}

func CreateScan(path string, source *datasources.DataSource, projection []string) *Scan {
	return &Scan{
		path:       path,
		source:     source,
		projection: projection,
	}
}

func (scan *Scan) GetSchema() *datatypes.Schema {
	source := *scan.source

	// Todo : Apply projection
	return source.GetSchema()
}

func (scan *Scan) GetChildrens() []LogicalPlan {
	return []LogicalPlan{}
}
