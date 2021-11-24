package main

import "fmt"

func HasRichFriend() bool  {
	return true
}
func GetFriendCount() int{
	return 3
}

func main() {
	// color := "red"
	// if color == "red" {
	// 	fmt.Println("나는 빨간색")
	// } else {
	// 	fmt.Println("아무것도 아님")
	// }

	var age int = 22
	if age>=10 && age<=15{
		fmt.Println("청소년")
	}

	price:=35000

	if price >50000{
		if HasRichFriend(){
			fmt.Println("난 돈이 없어")
		} else {
			fmt.Println("뿜빠이")
		}
	} else if price >= 30000 && price <=50000{
		if GetFriendCount()>3 {
			fmt.Println("크음...")
		} else{
			fmt.Println("제발 나눠 내자")
		}
	} else{
		fmt.Println("오늘 내 지갑을 털어봐")
	}





}