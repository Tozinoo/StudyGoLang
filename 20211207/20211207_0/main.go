package main

import "fmt"

type myNum int

func (a myNum) add(b int) int {
	return int(a) + b
}

func main() {
	// method : 함수의 일종임. 고에서는 class X
	// 구조체 밖에 메서드 지정.
	// 리시버 : 어느 구조체에 속하냐? 표시를 해줘야 함
	// 리시버 사용 -> 메서드가 속하는 타입을 알려주는 녀석

	// 메서드 선언하기
	rect1 := &Rect{width: 10, height: 10}
	rect2 := &Rect{width: 20, height: 20}
	fmt.Println(rect1.cal())
	fmt.Println(rect2.cal())

	var a myNum = 10
	fmt.Println(a.add(20))
	var b int = 20
	fmt.Println(myNum(b).add(100))
}

func (r *Rect) cal() int {
	return r.width * r.height
}

type Rect struct {
	width  int
	height int
}
