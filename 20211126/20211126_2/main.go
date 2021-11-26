package main

import "fmt"

// type StringHeader struct{
// 	Data uintptr // 문자열의 데이터가 있는 메모리 주소를 나타내는 일종의 포인터
// 	Len int // int타입의 문자열의 길이
// }

func main() {
	// str := "안녕 난 홍길동이야"
	// str1 := str
	// fmt.Println(str)
	// fmt.Println()
	// fmt.Println(str1)
	// // str의 Data와 Len값만 str1에 복사한다.

	// StringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str))
	// StringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	// println(StringHeader1)
	// println(StringHeader2)
	var str2 string = "Hello World"
	var slice []byte = []byte(str2)

	slice[2] = 't'
	println(str2)
	fmt.Printf("%s\n", slice)
}
