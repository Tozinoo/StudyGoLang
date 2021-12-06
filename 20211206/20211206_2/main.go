package main

import "fmt"

func main() {
	// type SliceHeader struct {
	// 	Data uintptr // 실제 배열을 가리키는 포인터
	// 	Len  int     // 요소 개수
	// 	Cap  int     // 실제 배열의 길이
	// }
	// var slice = make([]int, 3, 5)
	// slice = append(slice, 2)
	// slice = append(slice, 4)
	// slice = append(slice, 5)

	// slice1 := make([]int, 3, 5) // len : 3 cap :5 슬라이스 만듬
	// slice2 := append(slice1, 4, 5)
	// fmt.Println("슬라이스 1: ", slice1, len(slice1), cap(slice1))
	// fmt.Println("슬라이스 2: ", slice2, len(slice2), cap(slice2))
	// slice2 = append(slice2, 1)
	// fmt.Println("슬라이스 2: ", slice2, len(slice2), cap(slice2))

	// 슬라이싱 : 배열의 일부를 집어낸다
	// array[startIndex : endIndex]
	// 주의할 점 : endIndex -1까지

	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[1:2] //슬라이싱
	// fmt.Println("arr : ", arr)
	// fmt.Println("slice : ", slice, len(slice), cap(slice))

	// arr[1] = 100
	// fmt.Println("arr : ", arr)
	// fmt.Println("slice : ", slice, len(slice), cap(slice))

	// slice = append(slice, 200)
	// fmt.Println("arr : ", arr)
	// fmt.Println("slice : ", slice, len(slice), cap(slice))

	// slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := slice1[1:2]

	// cap 크기 조절하기(인덱스 3개로 )
	// slice[startIndex:endIndex:maxIndex]

	// slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := slice1[1:2:4]  // [2] cap : 3 ( maxIndex - startIndex)
	// fmt.Println(cap(slice2)) // 3

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1)) // 슬라이스1이랑 같은 길이의 슬라이스 만듬

	// 슬라이스 1의 모든 값을 복사
	copy(slice2, slice1)
	slice1[1] = 100

	fmt.Println(slice1)
	fmt.Println(slice2)

	// for i, v := range slice1 {
	// 	slice2[i] = v
	// }
}
