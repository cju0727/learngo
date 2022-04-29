package main

import "fmt"

func main() {
	// var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(t[i])
	// }

	// for i, v := range t {
	// 	fmt.Println(i, v)
	// }
	/*
		a := [5]int{1, 2, 3, 4, 5}
		b := [5]int{500, 400, 300, 200, 100}

		for i, v := range a {
			fmt.Printf("a[%d] = %d\n", i, v)
		}

		fmt.Println()

		for i, v := range b {
			fmt.Printf("b[%d] = %d\n", i, v)
		}
	*/
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}
	for _, arr := range a { //[0] : 12345, [1] : 56789
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	/*
		nums := [...]int{10, 20, 30, 40, 50}

		nums[2] = 300

		for i := 0; i < len(nums); i++ {
			fmt.Println(nums[i])
		}
	*/
}
