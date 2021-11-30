package main

import "fmt"

//constant declaration
const const_1 = "deger 1"

//multiple constant declaration
const (
	const_2 = "deger2"
	const_3 = "deger3"
	const_4 = "deger4"
)

//iota "asc" declaration
const (
	const_5 = iota
	const_6
	const_7
)

func main() {
	fmt.Println(const_1, const_2, const_3, const_4, const_5, const_6, const_7)
}
