// 포인터
package main

import (
	"fmt"
)

type Data struct {
	value int
	data  [200]int
}

func ChangeData(d Data) {
	d.value = 100
	d.data[100] = 200
	fmt.Printf("ChangeData안의 value 값 : %d", d.value)
	fmt.Println()
	fmt.Printf("ChangeData안의 data 값 : %d", d.data[100])
	fmt.Println()
}

func PointerChange(d *Data) { // 매개변수로 Data포인터로 받는다.
	d.value = 100
	d.data[100] = 200
}

func main() {

	var p *int // p는 int 타입 데이터의 메모리 주소를 가리키는 포인터 변수
	// 포인터는 메모리 주소를 값으로 가지기 때문에.. 숫자 이런거 안됨

	var a int = 100
	p = &a // 변수 a의 메모리 주소를 포인터 변수 p에 대입

	println("포인터 변수 p의 값(a의 메모리 주소) : ", p)
	println("p가 가리키는 메모리의 값 : ", *p)
	println("변경 이전 a 값 : ", a)

	*p = 300

	println("p가 가리키는 메모리의 값 : ", *p)
	println("변경 후 a 값 : ", a)

	// 포인터 변수 값 비교하기

	var x int = 20
	var y int = 10

	var p1 *int = &x // x의 메모리 공간
	var p2 *int = &x // x의 메모리 공간
	var p3 *int = &y // y의 메모리 공간

	fmt.Printf("p1==p2 : %v\n", p1 == p2)
	fmt.Printf("p2==p3 : %v\n", p2 == p3)

	// 포인터 변수는 초기화를 하지 않으면 default value는 nil
	// 0이긴 하지만 유효하지 않는 메모리 주소 값 -> 어떠한 공간도 가리키고 있지 않다.

	// var pointer *int

	// if pointer != nil {
	// 	// 포인터 변수가 nil이 아니라면 -> 유효한 메모리 주소를 가리킨다!
	// } else {
	// 	// 포인터 변수가 nil이라면 -> 유효한 메모리 주소를 가리키지 않는다.
	// }

	// 그래서 왜씀?
	// 대입이나 인수 전달 시 값을 복사하기 때문에 성능상에 이슈 발생.
	// 다른공간에 복사하기 때문에 변경사항이 적용 X

	var data Data

	// ChangeData(data)

	// fmt.Printf("value = %d \n", data.value)
	// fmt.Printf("data[100] = %d \n", data.data[100])

	PointerChange(&data) //data의 변수값이 아니라 data의 메모리 주소를 인수로 전달

	fmt.Printf("value = %d \n", data.value)
	fmt.Printf("data[100] = %d \n", data.data[100])
}
