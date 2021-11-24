package main

import "fmt"

func HasRichFriend() bool  {
	return true
}
func GetFriendCount() int{
	return 3
}

func GetMyAge ()(int,bool){
	return 100, true
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


	
	// if 초기문; 조건문 {
	// 		blah blah
	// }
	// 검사에 사용할 변수를 초기화 할 때 사용

	if age,ok := GetMyAge(); ok && age<20 {
		fmt.Println("젊음",age)
		} else if ok&& age<30{
		fmt.Println("아직 괜춘",age)

	} else if ok{
		fmt.Println("굿굿굿")
	} else {
		fmt.Println("아무것도 아님")
	}
	fmt.Println("ㅇㅇㅇ",age)




}