package main

import (
	"fmt"
)

const epsilon = 0.00001

/* 작은 오차 무시하기 1
func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}
*/
/*
func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}
*/
/* 작은 오차 무시하기 2
func main() {
		var a float64 = 0.1
		var b float64 = 0.2
		var c float64 = 0.3

		fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
		fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

		a = 0.0000000000004
		b = 0.0000000000002
		c = 0.0000000000007

		fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))
*/
/*
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d)) // c가 작은 경우 -1, 큰 경우 1, 같을 경우 0을 반환
}
*/
func main() {
	/*
		var a int = 10
		var b int = 20

		a, b = b, a
		fmt.Println(a, b)
	*/
	fmt.Println(3*4^7<<2+3*5 == 7)
}
