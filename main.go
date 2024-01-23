package main

import (
	"fmt"
	"strings"
	"time"
)

var out string = "out"

func main() {
	// fmt.Println("Hello World")
	// var name string = "kkh"
	// name := "kkh"
	// name = "lynn"
	// fmt.Println(name)

	// fmt.Println(multiply(2, 2))

	// totalLength, _ := lenAndUpper(name)
	// fmt.Println(totalLength)
	// repeatMe("1", "2", "3", "4")

	// total := superAdd(1, 2, 3, 4, 5, 6, 7)

	// fmt.Println(total)

	// go myCount("KKH1")
	// myCount("KKH2")

	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isPerson(person, c)
		// result := <-c
		if <-c {
			fmt.Println(person)
		}
	}
}

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	defer myCount("KKH3")
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

func myCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "count", i)
		time.Sleep(time.Second)
	}
}

func isPerson(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
