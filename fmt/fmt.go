package main

import (
	// "bufio"
	"fmt"
	// "os"
)

// fmt 기본 프린트 형태
/*
func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f:", f)
	fmt.Printf("a: %d b: %d f: %f\n", a, b, f)
}
*/

// 쵝소 출력 너비 지정
/*
func main() {
	a := 123
	b := 456
	c := 123456789

	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%05d, %05d\n", a, b)
	fmt.Printf("%-5d, %-05d\n", a, b)

	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-05d\n", c, c)
}
*/

// 실수 자릿수
/*
func main() {
	a := 324.13455
	c := 3.14

	fmt.Printf("%08.2f\n", a)
	fmt.Printf("%08.2g\n", a)
	fmt.Printf("%8.5g\n", a)
	fmt.Printf("%f\n", c)
}
*/

// 특수 문자
/*
func main() {
	str := "Hello\tGO\t\tWorld\n\"Go\"is Awesome!\n"

	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)
}
*/

// Scan(), Scanf(), Scanln()
/*
func main() {
	var a int
	var b int

	// n, err := fmt.Scan(&a, &b) // Scan()
	// n, err := fmt.Scanf("%d %d\n", &a, &b) // Scanf()
	n, err := fmt.Scanln(&a, &b) // Scanln()
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
*/

// Scan() 실패 시 남은 표준 입력 스트림 제거
/*
func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
*/

// 연습문제
func main() {
	var a int = 123
	var b int = 4567
	var f float64 = 3.14159269

	fmt.Printf("%6d\n", a)
	fmt.Printf("%06d\n", b)
	fmt.Printf("%6.2f\n", f)
}
