package main

import "fmt"

//type:global int array
var arr_1 [3]int                  //arr_1 = [0,0,0]
var arr_2 = [5]int{1, 2, 3, 4, 5} //arr_2 = [1,2,3,4,5]

func main() {

	arr_3 := make([]int, 3) //size 3 int array arr_3[0,0,0]

	arr_3[1] = 2 //arr_3 = [0,2,0]

	fmt.Println(arr_1, arr_2, arr_3)
	fmt.Printf("arr_1 len:%d \n", len(arr_1))
	fmt.Printf("arr_2 len:%d \n", len(arr_2))

}
