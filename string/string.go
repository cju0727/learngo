package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2
	fmt.Println((str1))
	/*
		str := "Hello 월드!"
		for _, v := range str {
			fmt.Printf(" 타입:%T 값:%d 문자:%c\n", v, v, v)
		}
	*/
	/* []rune을 이용한 문자열 순회
	str := "Hello 월드!"
	arr := []rune(str)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}
	*/
	/* for문을 이용한 문자열 순회 (한글 깨짐)
	str := "Hello 월드!"

	for i := 0; i < len(str); i++ {
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	}
	*/
	/* str과 []rune 타입 변환
	str := "Hello 월드"
	runes := []rune(str)

	fmt.Printf("len(str) = %d\n", len(str))
	fmt.Printf("len(runes) = %d\n", len(runes))
	*/
	/* []rune 타입
	str := "Hello world"

	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(str)
	fmt.Println(string(runes))
	*/

	/* 한글과 영어의 바이트 차이
	str1 := "가나다라마"
	str2 := "abcde"
	fmt.Printf("len(str1) = %d\n", len(str1))
	fmt.Printf("len(str2) = %d\n", len(str2))
	*/

	/* rune 타입
	var char rune = '한'

	fmt.Printf("%T\n", char)
	fmt.Println(char)
	fmt.Printf("%c\n", char)
	*/

	// 	poet1 := "죽는 날까지 하늘을 우러러\n한 점 부끄럼이 없기를,\n잎새에 이는 바람에도\n나는 괴로워했다.\n"

	// 	poet2 := `죽는 날까지 하늘을 우러러
	// 한 점 부끄럼이 없기를,
	// 잎새에 이는 바람에도
	// 나는 괴로워했다.`
	// 	fmt.Println(poet1)
	// 	fmt.Println(poet2)

	// str1 := "Hello\t'world'\n"

	// str2 := `Go is "awesome"!\nGo is simple and\t'powerful'`
	// fmt.Println(str1)
	// fmt.Println(str2)

}
