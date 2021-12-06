package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello"
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str)) // 내부 구조체로 변경
	address := stringHeader.Data                                  // Data필드값을 변수로 저장
	str += "World"

	address1 := stringHeader.Data // Data필드값을 변수로 저장
	str += "Hi"
	address2 := stringHeader.Data // Data필드값을 변수로 저장

	fmt.Println(str)
	fmt.Printf("주소  : \t%x\n", address)
	fmt.Printf("주소1 : \t%x\n", address1)
	fmt.Printf("주소2 : \t%x\n", address2)

}
