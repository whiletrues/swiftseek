package main

import (
	"fmt"

	"github.com/apache/arrow/go/arrow"
	"github.com/whiletrues/swiftseek/src/datasources"
	"github.com/whiletrues/swiftseek/src/datatypes"
	logicalplan "github.com/whiletrues/swiftseek/src/logicalplans"
)

func main() {

	schema := datatypes.CreateSchema([]datatypes.Field{
		datatypes.CreateField("column1", arrow.INT32),
		datatypes.CreateField("column2", arrow.INT32),
	})

	datatypes.CreateLiteralValueVector(arrow.INT32, 1, 1)
	datatypes.CreateLiteralValueVector(arrow.INT32, 2, 1)

	records := []datatypes.RecordBatch{
		datatypes.CreateRecord([]datatypes.ColumnVector{
			datatypes.CreateLiteralValueVector(arrow.INT32, 1, 1),
			datatypes.CreateLiteralValueVector(arrow.INT32, 2, 1),
		}),
	}

	source := datasources.CreateInMemorySource(schema, &records)

	projection := logicalplan.CreateScan("./source", source, []string{"column1", "column2"})

	booleanexpr := logicalplan.CreateBooleanExpr("Eq", logicalplan.Eq, nil, nil)

	aggregateexpr := logicalplan.CreateAggregateExpr("SUM", booleanexpr)

	mathexpr := logicalplan.CreateMathExpr("Add", logicalplan.Add, booleanexpr, aggregateexpr)

	fmt.Println(mathexpr.String())

	columnVector := datatypes.CreateLiteralValueVector(arrow.INT32, 1, 1)

	if columnVector.GetType() == arrow.INT32 {
		fmt.Println("ok")
	}
}
