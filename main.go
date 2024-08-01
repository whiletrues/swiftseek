package main

import (
	"fmt"

	"github.com/apache/arrow/go/arrow"
)

func main() {
	fmt.Print("hello world")

	//pool := memory.NewGoAllocator()
	arrow.StructOf([]arrow.Field{
		{Name: "a", Type: arrow.PrimitiveTypes.Int32},
	}...)
}
