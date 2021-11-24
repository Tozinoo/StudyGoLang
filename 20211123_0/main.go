package main

import (
	"fmt"
)

func main() {
	
	

	var a int =10
	var b int = 20
	var f float64 = 123123.456

	fmt.Println(1)
	fmt.Printf("a : %d b: %d f: %f\n",a,b,f)

	// var input int32
	// var input1 int32
	
 	// fmt.Scan(&input,&input1)

	// 비트연산 AND, OR, XOR, 비트클리어
	// & , | , ^ , &^
	// 시프트연산 <<(왼쪽 시프트) , >>(오른쪽 시프트)
	// 0000 0001  <<1
	// 0000 0010

	var fa float64 =0.1
	var fb float64 =0.2
	var fc float64 =0.3

	fmt.Printf("%f+%f == %f : %v\n", fa,fb,fc,fa+fb==fc)
	fmt.Println(fa+fb)

	
	
}