package main

import (
	"fmt"
	"unsafe"
)

type Test struct {
	A int8
	C int8
	E int8
	B int
	D int
}
type Student struct {
	Age   int
	Score float64
}

func PrintStudent(st Student) {
	fmt.Printf("나이 : %d, 점수 : %.2f", st.Age, st.Score)
}

func main() {

	var s = Student{15, 20.5555}
	PrintStudent(s)

	s1 := s          //s 구조체 모든 멤버(필드)들이 s1 복사
	PrintStudent(s1) // 함수 호출할 때도 구조체가 복사

	fmt.Println(unsafe.Sizeof(s))

	var c = Test{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(c))

}
