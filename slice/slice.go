package main

import "fmt"

func main() {
	var slice []int

	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
	}

	slice = append(slice, 11, 12, 13, 14, 15)
	fmt.Println(slice)

	/* append 1
	var slice = []int{1, 2, 3}
	slice2 := append(slice, 4)

	fmt.Println(slice)
	fmt.Println(slice2)
	*/

	/* slice 기초
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10
	fmt.Println(slice)
	*/
}
