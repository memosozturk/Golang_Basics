package main

import (
	"fmt"
	"strings"
)

type String *string

func main() {
	var rstr String
	var pstr string

	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "mehmet ozturk"
	rstr = &pstr

	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "mehmet ozturk hataylı"

	fmt.Println(*rstr)
	fmt.Println(pstr)

	*rstr = "mehmet ozturk hataylı 2021"

	fmt.Println(*rstr)
	fmt.Println(pstr)

	replaceStr(pstr)

	fmt.Println(*rstr)
	fmt.Println(pstr)

	replaceStr2(rstr)

	fmt.Println(*rstr)
	fmt.Println(pstr)
}
func replaceStr(s string) {
	s = strings.Replace(s, "mehmet", "MEHMET", 1)
}
func replaceStr2(s String) {
	*s = strings.Replace(*s, "mehmet", "MEHMET", 1)
}
