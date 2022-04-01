package main

import (
	"fmt"
	"strings"
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

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

// func main() {
// 	repeatMe("nico", "lynn", "dal", "marl", "flynn")
// }

// Naked return, defer
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLenght, up := lenAndUpper("nico")
	fmt.Println(totalLenght, up)
}
