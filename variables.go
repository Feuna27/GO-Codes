package main

import (
	"fmt"
)

func main() {
	var student1 string = "john" //type is string
	var student2 = "jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Printf("student1: %s\n ", student1)
	fmt.Printf("student2: %s\n", student2)
	fmt.Printf("x: %d\n", x)
	fmt.Printf("x is of type %T\n", x)
}
