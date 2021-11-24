package main

import "fmt"

func HasRichFriend() bool {
	return true
}
func GetFriendCount() int {
	return 3
}

func GetMyAge() (int, bool) {
	return 100, true
}
func SwitchGetAge() int {
	return 30
}
func GetColor() ColorType {
	return Blue
}

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func ColorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Red"
	case Green:
		return "Red"
	case Yellow:
		return "Red"
	default:
		return "undefined"
	}
}

func main() {
	// color := "red"
	// if color == "red" {
	// 	fmt.Println("나는 빨간색")
	// } else {
	// 	fmt.Println("아무것도 아님")
	// }

	var age int = 22
	if age >= 10 && age <= 15 {
		fmt.Println("청소년")
	}

	price := 35000

	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("난 돈이 없어")
		} else {
			fmt.Println("뿜빠이")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendCount() > 3 {
			fmt.Println("크음...")
		} else {
			fmt.Println("제발 나눠 내자")
		}
	} else {
		fmt.Println("오늘 내 지갑을 털어봐")
	}

	// if 초기문; 조건문 {
	// 		blah blah
	// }
	// 검사에 사용할 변수를 초기화 할 때 사용

	if age, ok := GetMyAge(); ok && age < 20 {
		fmt.Println("젊음", age)
	} else if ok && age < 30 {
		fmt.Println("아직 괜춘", age)

	} else if ok {
		fmt.Println("굿굿굿")
	} else {
		fmt.Println("아무것도 아님")
	}
	fmt.Println("ㅇㅇㅇ", age)

	// switch
	num := 3
	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("0")
	}

	// 하나이상의 값도 구분할 수 있다. 쉼표로 구분...
	day := "sunday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("언제 금요일이 오나?")
	case "wednesday", "friday", "sunday":
		fmt.Println("퇴근시켜줘")
	}
	num1 := 18
	switch true { // 비교할 값이 true를 찾는다. -> case가 true인 녀석을 찾는다.
	case num1 < 10:
		fmt.Println("asdfasdf")
	case num1 < 30:
		fmt.Println("131313")

	case num1 >= 10 && num1 < 20:
		fmt.Println("하하하")
	}

	// if문처럼 초기문을 넣을 수 있다.
	switch age := SwitchGetAge(); age {
	case 100:
		fmt.Println("100살")
	case 30:
		fmt.Println("30살")
	}

	switch age := SwitchGetAge(); true {
	case age < 10:
		fmt.Println("10살 미만")

	case age < 20:
		fmt.Println("20살 미만")

	case age < 30:
		fmt.Println("30살 미만")
	case age < 40:
		fmt.Println("40살 미만")
	case age < 50:
		fmt.Println("하이")
	}

	//fallthroug : 다음 cas까지 실행한다. (끝까지 해주는거 아님!)
	abc := 3
	switch abc {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
		fallthrough
	case 5:
		fmt.Println("5")
	}

	// const 열거형에 따라 로직을 변경 할 때 쓰면 유용하다.

	fmt.Println("색깔", ColorToString(GetColor()))

}
