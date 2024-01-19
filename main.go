package main

import (
	"fmt"
	"strings"
)

var out string = "out"

func main() {
	// fmt.Println("Hello World")
	// var name string = "kkh"
	name := "kkh"
	name = "lynn"
	fmt.Println(name)

	fmt.Println(multiply(2, 2))

	totalLength, _ := lenAndUpper(name)
	fmt.Println(totalLength)
	repeatMe("1", "2", "3", "4")

	total := superAdd(1, 2, 3, 4, 5, 6, 7)

	fmt.Println(total)
}

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	defer fmt.Println("I'm done")
	fmt.Println("Run")
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}
