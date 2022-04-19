package main

import "fmt"

/*
var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}
*/
/* 중접 if
func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}
*/
/* if 초기문; 조건문
func getMyAge() (int, bool) {
	return 33, true
}
*/

func main() {
	/* if 초기문; 조건문
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	// fmt.Println("Your age is", age) - if 초기문 종료 시 변수 할당 X
	*/

	/* 중첩 if
	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("앗 신발끈이 풀렸네")
		} else {
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendsCount() > 3 {
			fmt.Println("어이쿠 신발끈이..")
		} else {
			fmt.Println("사람도 얼마 없는데 나눠 내자")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다.")
	}
	*/
	/* 쇼트서킷
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}
	if true || IncreaseAndReturn() < 5 {
		fmt.Println("2 증가")
	}
	fmt.Println("cnt:", cnt)
	*/
	/* 함수 예시 1
	light := "red"
	light = "green"

	if light == "green" {
		fmt.Println("길을 건넌다.")
	} else {
		fmt.Println("대기한다.")
	}
	*/
	/* 함수 예시 2
	temp := 33

	if temp > 28 {
		fmt.Println("에어컨을 켠다.")
	} else if temp <= 3 {
		fmt.Println("히터를 켠다.")
	} else {
		fmt.Println("대기한다.")
	}
	*/
	/*
		var age int = 22

		if age >= 10 && age <= 15 {
			fmt.Println("You are young")
		} else if age > 30 || age < 20 {
			fmt.Println("You are not 20s")
		} else {
			fmt.Println("Best age of your life")
		}
	*/
	temp := 20
	rain := 90

	if temp >= 25 {
		if rain >= 80 {
			fmt.Println("덥고 비가 옵니다.")
		} else if rain >= 20 {
			fmt.Println("덥고 습합니다.")
		} else {
			fmt.Println("야외 활동하기 좋습니다.")
		}
	} else if temp < 10 || rain >= 80 {
		fmt.Println("야외 활동하기 좋지 않습니다.")
	} else {
		fmt.Println("좋은 날씨입니다.")
	}
}
