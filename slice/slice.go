package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2]

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100

	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice = append(slice, 500)

	fmt.Println("After append 500")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	/*
		slice1 := []int{1, 2, 3}
		slice2 := append(slice1, 4, 5)

		fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
		fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

		slice1[1] = 100

		fmt.Println("After change second element")
		fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
		fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

		slice1 = append(slice1, 500)

		fmt.Println("After append 500")
		fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
		fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	*/

	/*slice append 예시 1
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	*/

	/*array & slice
	  func changeArray(array2 [5]int) {
	  	array2[2] = 200
	  }

	  func changeSlice(slice2 []int) {
	  	slice2[2] = 200
	  }

	  func main() {
	  	array := [5]int{1, 2, 3, 4, 5}
	  	slice := []int{1, 2, 3, 4, 5}

	  	changeArray(array)
	  	changeSlice(slice)

	  	fmt.Println(array)
	  	fmt.Println(slice)
	*/

	/*
		var slice []int

		for i := 1; i <= 10; i++ {
			slice = append(slice, i)
		}

		slice = append(slice, 11, 12, 13, 14, 15)
		fmt.Println(slice)
	*/

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
