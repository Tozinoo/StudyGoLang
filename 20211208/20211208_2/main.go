package main

import "fmt"

func cap1() {
	f := make([]func(), 3) // 함수 리터럴 3개를 가진 슬라이스
	fmt.Println("캡1")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]() // 호출 되는 시점에 i 값은 3
	}
}
func cap2() {
	f := make([]func(), 3) // 함수 리터럴 3개를 가진 슬라이스
	fmt.Println("캡2")
	for i := 0; i < 3; i++ {
		v := i // v 변수에 i값 복사
		f[i] = func() {
			fmt.Println(v)
		}
	}
	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	cap1() // 3 3 3
	cap2() // 0 1 2
}
