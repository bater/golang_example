package main

import (
	"fmt"
	tool "test"
)

func main() {
	i := 2
	fmt.Printf("Is %d even? %v\n", i, tool.Even(i))
}
