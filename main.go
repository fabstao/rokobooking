package main

import "fmt"

type Test struct {
	Name string
}
type Yo struct {
	Test
	apellido string
}

func main() {

	yo := Yo{
		Test{"yo"},
		"el",
	}
	fmt.Println(yo.Name, yo.apellido)
}
