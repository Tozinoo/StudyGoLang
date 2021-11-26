package main

import "fmt"

func main() {
	// 문자열 순회
	// 1. 인덱스를 사용해 직접 접근하는 방법
	var str string = "Hello 경일"
	// for i := 0; i < len(str); i++ { // 문자열 크기만큼 돌림
	// 	fmt.Printf("타입 : %T, 값 : %d, 문자 : %c\n", str[i], str[i], str[i])
	// }

	// 위 녀석은 한글 깨짐. 타입이 uint8이기 때문.. 얘는 1바이트이므로 ( 한글 3바이트 )
	// 한글깨짐 방지를 위해 []타입 OR range를 이용한다.

	// 깨짐현상을 막기 위한 []rune으로 타입 변환

	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입 : %T, 값 : %d, 문자 : %c\n", arr[i], arr[i], arr[i])
	}

	// for _, v := range v {

	// }
}
