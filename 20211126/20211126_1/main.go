package main

func main() {
	// 문자열 순회
	// 1. 인덱스를 사용해 직접 접근하는 방법
	// var str string = "Hello 경일"
	// for i := 0; i < len(str); i++ { // 문자열 크기만큼 돌림
	// 	fmt.Printf("타입 : %T, 값 : %d, 문자 : %c\n", str[i], str[i], str[i])
	// }

	// 위 녀석은 한글 깨짐. 타입이 uint8이기 때문.. 얘는 1바이트이므로 ( 한글 3바이트 )
	// 한글깨짐 방지를 위해 []타입 OR range를 이용한다.

	// 깨짐현상을 막기 위한 []rune으로 타입 변환

	// arr := []rune(str)
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("타입 : %T, 값 : %d, 문자 : %c\n", arr[i], arr[i], arr[i])
	// 	// 변환하는 과정에서 별도의 배열을 할당 -> 불필요한 메모리를 사용한다.
	// }

	// for _, v := range str {
	// 	fmt.Printf("타입 : %T, 값 : %d, 문자 : %c\n", v, v, v)
	// }

	// 문자열 연산
	var str1 string = "경일"
	var str2 string = "게임"

	var str3 = str1 + str2 // 문자열 끼리 + 는 문자열 잇기
	println(str3)
	str1 += str2
	println(str1)

	// 문자열 비교
	str4 := "hello"
	str5 := "helloo"
	str6 := "hello"

	println(str4 == str5)
	println(str5 == str6)
	println(str4 == str6)
	println(str4 != str6)
	println(str4 != str5)
	println()

	// <,>,<=,>= 문자열 대소비교

	// 문자열 대소비교 길이와 상관없이 첫글자부터 비교.
	str7 := "AAA"
	str8 := "AAB"
	println(str7 < str8)
}
