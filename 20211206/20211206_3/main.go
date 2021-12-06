package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student //[]Student의 별칭 Students

func main() {
	// 삭제
	// slice := []int{1, 2, 3, 4, 5, 6}
	// index := 2

	// for i := index + 1; i < len(slice); i++ {
	// 	slice[i-1] = slice[i]
	// }
	// slice = slice[:len(slice)-1]
	// fmt.Println(slice)

	// 삽입
	// slice1 := []int{1, 2, 3, 4, 5, 6}
	// slice1 = append(slice1, 0)
	// index1 := 4

	// // 맨 뒤부터 추가할혀는 위치까지 값을 하나씩 옮긴다
	// for i := len(slice1) - 2; i >= index1; i-- {
	// 	slice1[i+1] = slice1[i]
	// }

	// slice1[index1] = 299
	// fmt.Println(slice1)

	s := []int{200, 45, 677, 123, 6}
	sort.Ints(s)
	fmt.Println(s)
}
