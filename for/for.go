package main

import (
	"fmt"
	// "time"
)

func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}

func main() {
	a := 1
	b := 0

	for ; a <= 9; a++ {
		var found bool
		if b, found = find45(a); found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	//  초기문;  조건문;  후처리
	// for i := 0; i < 10; i++ {
	// 	fmt.Print(i, ", ")
	// }

	// 예시 1
	// i := 1
	// for {
	// 	time.Sleep(time.Second)
	// 	fmt.Println(i)
	// 	i++
	// }

	/* 예시 2
	   stdin := bufio.NewReader(os.Stdin)
	   for {
	   	fmt.Println("입력하세요.")
	   	var number int
	   	_, err := fmt.Scanln(&number)
	   	if err != nil {
	   		fmt.Println("숫자를 입력하세요.")

	   		stdin.ReadString('\n')
	   		continue
	   	}
	   	fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
	   	if number%2 == 0 {
	   		break
	   	}
	   }
	   fmt.Println("for문이 종료되었습니다.")
	*/
	/* 이중 for문 예시 1
	   for i := 0; i < 3; i++ {
	   	for j := 0; j < 5; j++ {
	   		fmt.Print("*")
	   	}
	   	fmt.Println()
	   }
	*/
	/* 이중 for문 예시 2
	   dan := 2
	   b := 1
	   for {
	   	for {
	   		fmt.Printf("%d + %d = %d\n", dan, b, dan*b)
	   		b++
	   		if b == 10 {
	   			break
	   		}
	   	}
	   	b = 1
	   	dan++
	   	if dan == 10 {
	   		break
	   	}
	   }
	   fmt.Println("for문이 종료됐습니다.")
	*/
}

/* 연습문제 3
func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d * %d = %d\n", i, i, i*i)
	}
}
*/
/* 연습문제 4
func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5 - i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
*/
