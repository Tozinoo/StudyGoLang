package main

import "fmt"

const APPLE int = 0
const BANANA int = 1
const GRAPGE int = 2

const PI = 3.14
const fPI float64 = 3.14

func PrintFruit(fruit int) {
	if fruit == APPLE {
		fmt.Println("사과는 아침에 먹어야 제맛")
	} else if fruit==BANANA {
		fmt.Println("나는 바나나")
	} else{
		fmt.Println("나는 포도")
	}
}
func main() {
	
	PrintFruit(APPLE)
	PrintFruit(BANANA)
	PrintFruit(GRAPGE)

	// iota : 1씩 증가하도록 정의할 때는 이녀석을 쓰는게 편리. 열거형
	const(
		RED int = iota
		BLACK int = iota
		YELLOW int = iota
	)
	const(
		TIGER int = iota +1
		RABBIT 
		DOG 
	)
	
	// var a int = PI * 100
	// var b int = fPI * 100
}