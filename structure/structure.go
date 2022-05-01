package main

import "fmt"

/* 연습문제 1
type Product sturct {
	Name string
	Price int
	ReviewScore float64
}
*/

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}
	fmt.Printf("유저: %s, ID: %s, 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s, ID: %s, 나이: %d, VIP 레벨: %d, VIP 가격: %d만원\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price,
	)
}

// type House struct {
// 	Address string
// 	Size    int
// 	Price   float64
// 	Type    string
// }

/* 일반 structure
func main() {
	// var house House
	// house.Address = "서울시 강동구 ..."
	// house.Size = 28
	// house.Price = 9.8
	// house.Type = "아파트"

	var house House = House{
		"서울시 강동구 ...",
		28,
		9.8,
		"아파트",
	}

	fmt.Println("주소:", house.Address)
	fmt.Println("크기:", house.Size)
	fmt.Println("가격:", house.Price)
	fmt.Println("타입:", house.Type)
}
*/
