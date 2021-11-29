package main

import "fmt"

func main() {
	//regular declaration
	var var_1 string
	var_1 = "variable declaration"
	//Short declaration
	var_2 := 11
	fmt.Println(var_1, var_2)
	//multiple declaration
	mehmet, ahmet, ali, emre, efe := 1, 2, 3, 4, 5
	fmt.Println(mehmet, ahmet, ali, emre, efe)
}
