package main

import "fmt"

//Global variable
//multiple declaration
var (
	firstName = "mehmet"
	lastName  = "ozturk"
	age       = 25
)

func main() {
	//short multiple declaration
	a, b, c, d := 1, 2, "go", true
	fmt.Println(a, b, c, d)
	//multiple declaration(not short declaration)
	var i, j int = 1, 2
	fmt.Println(i, j)
	//multiple assignment
	i, j = 2, 1
	fmt.Println(i, j)

}
