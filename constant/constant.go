package main

import "fmt"

/*
const Pig int = 0
const Cow int = 1
const Chicken int = 2
*/
const PI = 3.14

// const Gravity = 9.80665

// const FloatPI float64 = 3.14

/*
func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}
*/

func main() {
	/*
		const PI1 float64 = 3.141592653589793238
		var PI2 float64 = 3.141592653589793238

		PI2 = 4

		fmt.Printf("%f\n", PI1)
		fmt.Printf("%f\n", PI2)
	*/
	/*
		PrintAnimal(Pig)
		PrintAnimal(Cow)
		PrintAnimal(Chicken)
	*/
	var a int = PI * 100
	// var b int = FloatPI * 100 // error 발생

	fmt.Println(a)
	// fmt.Println(b)

}
