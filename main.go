package main

import (
	"fmt"
)

// 자료형 지정1

// func multiply(a int, b int) int {
// 	return a * b
// }

// func main() {
// 	fmt.Println((multiply(2, 2)))
// }

// 자료형 지정2

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// func main() {
// 	totalLenght, upperName := lenAndUpper("jaewoo")
// 	fmt.Println(totalLenght, upperName)
// }

// 자료형 지정3

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}
