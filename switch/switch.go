package main

import "fmt"

/*
func getMyAge() int {
	return 22
}
*/
/* const 열거값과 switch
type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undifined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}
*/
/* switch 연습 문제 3번
func GetDirection(angle float64) Direction {
	switch angle {
	case angle >= 315, angle >= 0 && angle < 45:
		return North
	case angle >= 45 && angle < 135:
		return East
	case angle >= 135 && angle < 225:
		return South
	case angle >= 225 && angle < 315:
		return West
	default:
		return None
	}
}
*/

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
	/* case에서 break
	a := 3

	switch a {
	case 1:
		fmt.Println("a==1")
		break
	case 2:
		fmt.Println("a==2")
		break
	case 3:
		fmt.Println("a==3")
	case 4:
		fmt.Println("a==4")
	default:
		fmt.Println("a > 4")
	}
	*/

	// fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))

	/* switch 초기문 사용
	switch age := getMyAge(); true {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
	*/
	/* 스위치 예시 1
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
	*/
	/* 스위치 예시 2
	day := 3

	switch day {
	case 1:
		fmt.Println("첫째 날입니다.")
		fmt.Println("오늘은 팀 미팅이 있습니다.")
	case 2:
		fmt.Println("둘째 날입니다.")
		fmt.Println("오늘은 면접이 있습니다.")
	default:
		fmt.Println("프로젝트를 진행하세요.")
	}
	*/
	/* 한 번에 여러 값 비교
	day := "thursday"

	switch day {
	case "monday", "thesday":
		fmt.Println("월, 화요일은 수업 가는 날입니다.")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금요일은 실습 가는 날입니다.")
	}
	*/
	/* 조건문 비교
	temp := 10

	switch true {
	case temp < 10, temp > 30:
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다.")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요.")
	case temp >= 15 && temp < 25:
		fmt.Println("야외 활동하기 좋은 날씨입니다.")
	default:
		fmt.Println("따뜻합니다.")
	}
	*/

}
