package main

import "fmt"

/*
 가변 인수를 받는 함수
 인수 타입 앞에 ... 해당 타입 인수를 여러개 받는 가변인수
 가변 인수는 함수 내부에서 해당 타입의 슬라이스로 처리
*/

// func sum(nums ...int) int {
// 	sum := 0

// 	fmt.Printf("타입 : %T\n", nums)
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	return sum
// }

// func Print(args ...interface{}) string { // 모든 타입을 받는 가변 인수
// 	for _, arg := range args { // 모든 인수를 돌면서
// 		switch b := arg.(type) { // 인수 타입에 따라서
// 		case bool:
// 			val := arg.(bool) // 변환
// 			// 출력 ...
// 		case int:
// 			val := arg.(int) // 변환
// 			// 출력 ...
// 		}
// 	}
// }

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b // 함수 리터럴(람다)을 이용해서 더하기 함수 정의 후 리턴
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b // 함수 리터럴(람다)을 이용해서 곱하기 함수 정의 후 리턴
		}
	} else {
		return nil
	}
}

func main() {

	// fmt.Println(sum(1, 2, 3, 4, 5)) // 타입 : [] int , 15
	// fmt.Println(sum(1, 2))          // 타입 : [] int , 3
	// fmt.Println(sum())              // 타입 : [] int , 0

	// f, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// defer fmt.Println("호출")
	// defer f.Close()

	/*
	 함수 리터럴
	 이름이 없는 함수 -  함수명을 쓰지 않고 함수 타입 변수 값으로 대입 되는 함수 값 (call by 익명함수 or 람다)
	 함수 명이 없기 때문에 직접 호출 불가
	 함수 타입 변수로만 호출 가능
	*/
	// fn := getOperator("*")
	// res := fn(3, 4) // 함수 타입 변수를 사용해서 함수 호출
	// fmt.Println(res)

	a := 0
	f := func() {
		a += 10 // 외부 변수는 자동으로 함수 내부로 저장
	}
	a++
	f() // 함수타입 변수
	fmt.Println(a)
}
