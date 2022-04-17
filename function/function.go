package main

import (
	"fmt"
)

/* 간단한 함수의 예시
func add(a int, b int) int {
	return a + b
}

func main() {
	c := add(3, 6)
	fmt.Println(c)
}
*/
/* 비효율적인 코드
func main() {
	math := 80
	eng := 74
	history := 95
	fmt.Println("김일등 님의 점수는", (math+eng+history)/3, "입니다.")

	math = 88
	eng = 92
	history = 53
	fmt.Println("송이등 님의 점수는", (math+eng+history)/3, "입니다.")
}
*/
/* 함수를 사용한 효율적인 코드
func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님의 평균 점수는", avg, "입니다.")
}

func main() {
	PrintAvgScore("김일등", 80, 74, 95)
	PrintAvgScore("송이등", 88, 92, 53)
}
*/
/* 멀티 반환 함수
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
*/
/* 변수명을 지정해서 반환
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
*/
/* 재귀 함수
func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func main() {
	printNo(3)
}

// 동작원리
printNo(3) : 3
	printNo(2) : 2
		printNo(1) : 1
			printNo(0) : 종료
		"After", 1
	"After", 2
"After", 3
*/

func Multiple(a, b int) int {
	return a * b
}

func main() {
	c := Multiple(3, 5)
	fmt.Println(c)
}
