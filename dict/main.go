package main

import (
	"fmt"

	dict "github.com/go-study/dict/dictionary"
)

func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello"
	definition, err := dictionary.Search("first")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
