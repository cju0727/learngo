package main

import "fmt"

// type Data struct {
// 	value int
// 	data  [200]int
// }

// func ChangeData(arg *Data) {
// 	arg.value = 999
// 	arg.data[100] = 999
// }

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	userPointer := NewUser("AAA", 23)

	fmt.Println(userPointer)

	// var data Data

	// ChangeData(&data)

	// fmt.Printf("value = %d\n", data.value)
	// fmt.Printf("data[100] = %d\n", data.data[100])

	// var a int = 500
	// var p *int // int 포인터 변수 p 선언

	// p = &a // a의 메모리 주소를 p에 대입(복사)

	// fmt.Printf("p의 값 : %p\n", p)
	// fmt.Printf("p가 가리키는 메모리의 값 : %d\n", *p)
	// *p = 100
	// fmt.Printf("a의 값 : %d\n", a)

	/*
		var a int = 10
		var b int = 20

		var p1 *int = &a
		var p2 *int = &a
		var p3 *int = &b

		fmt.Printf("p1 == p2 : %v\n", p1 == p2)
		fmt.Printf("p2 == p3 : %v\n", p2 == p3)
	*/
}
