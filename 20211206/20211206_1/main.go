package main

import "fmt"

func main() {
	// 슬라이스 : 동적 배열 -> 자동으로 배열 크기를 증가시켜준다.
	// 길이가 요소 개수에 따라 자동으로 증가해 관리가 편하다
	// 슬라이싱 기능을 사용해서 배열의 일부를 나타내는 슬라이스를 만들 수 있다.

	// var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 슬라이스를 초기화 하지 않으면 길이가 0인 슬라이스가 만들어진다.
	// var slice []int
	// if len(slice) == 0 {
	// 	fmt.Println("슬라이스가 비어 있다", slice)
	// }

	// //var slice1 = []int{1, 2, 3}
	// var slice2 = []int{1, 5: 2, 10: 3}
	// fmt.Println(slice2)

	// // var arr = [...]int{1, 2, 3} // 배열
	// // var slice = []int{1, 2, 3}  // 슬라이스!!!!!!!!!!! 헷갈리지 마시오

	// // make()를 이용한 초기화
	// // var slice3 = make([]int, 3)
	// // fmt.Println(slice3)

	// // 슬라이스 접근 :  배열과 같다
	// var slice4 = make([]int, 3)
	// slice4[1] = 5

	// // 슬라이스 순회 -> 동적으로 길이가 늘어나는것 말고, 배열과 동일함.
	// var slice5 = make([]int, 5)

	// for i := 0; i < len(slice5); i++ {
	// 	slice5[i] = i + 1
	// }
	// fmt.Println(slice5)

	// var slice6 = []int{1, 2, 3}
	// slice7 := append(slice6, 4)
	// fmt.Println(slice7)

	// append : 첫번째 인수로 들어온 슬라이스의 값을 변경하는 게 아니라 요소가 추가된 새로운 슬라이스를 반환
	// 기존 슬라이스에 요소를 추가하고 싶다? append()결과를 기존 슬라이스에 대입해서 변경해야 된다

	var slice8 []int
	for i := 0; i <= 10; i++ {
		slice8 = append(slice8, i)
	}
	var slice9 []int = slice8
	fmt.Println(slice9)

}
