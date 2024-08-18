package main

import (
	"fmt"

	"github.com/apache/arrow/go/arrow"
	"github.com/whiletrues/swiftseek/src/datatypes"
	logicalplan "github.com/whiletrues/swiftseek/src/logicalplans"
)

func main() {

	logicalplan.CreateBooleanExpr("", logicalplan.Eq, nil, nil)

	columnVector := datatypes.CreateLiteralValueVector(arrow.INT32, 1, 1)

	if columnVector.GetType() == arrow.INT32 {
		fmt.Println("ok")
	}
}
