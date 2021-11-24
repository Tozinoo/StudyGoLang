package main

/*
	패키지 선언: 이 코드가 어떤 패키지에 속하나?
	go언어의 모든 코드는 반드시 패키지 선언으로 시작
	package main: main패키지에 속한 코드다! 라는 것을 컴파일러에게 알려줌
	go언어는 패키지로 시작하고 package main은 프로그램 시작점이 있는 패키지다!
	main 패키지 필수로 필요
*/

/*
	클래스 X 메서드를 가지는 구조체를 지원
	상속 X
	메서드 O
	인터페이스 O
	익명함수 O
	가비지 컬렉터 O
	포인터 O
	제네릭 프로그래밍 X (일반화)
	네임스페이스 X 패키지 단위로 구분
*/


func main() {
	// fmt.Print("나는 프린트 함수")
	// fmt.Println("Hello world");
	// fmt.Print("나는 프린트 함수")
	// fmt.Printf("나는 프린트함수\n")
	// // var : 변수를 선언하는 키워드
	// // int : 정수형 타입
	// // 변수는 항상 동일하게 

	// var a int = 10
	// var b int = 20
	// var c float32 = 20.5
	// var plus1 int = a+b
	// // var plus2 int = a+c
	// var msg string = "Hello"

	// fmt.Println(a)
	// fmt.Println(msg)
	// fmt.Println(plus1)
	// fmt.Println(c)

	// // 변수 속성은 4가지가 있다.
	// // 이름 값 타입 주소

	// // 숫자 타입
	// // uint(unsigned int) 부호가 없는 정수 ex) uint8 1바이트 -> 0~255,
	// // 32비트 환경에서 int는 int32
	// // 64비트 환경에서 int는 int64
	// // byte = uint8의 별칭
	// // float
	// // complex
	// // rune = int32의 별칭 UTF-8문자 하나를 나타낼 때 사용
	
	// // bool
	// // string

	// // Sizeof
	// var int_size int
	// var int_size8 int8
	// var int_size16 int16
	// var int_size32 int32
	// var int_size64 int64
	// var uint_size64 uint64

	// fmt.Println("int형 사이즈 : ",unsafe.Sizeof(int_size),"바이트")
	// fmt.Println("int64형 사이즈 : ",unsafe.Sizeof(int_size8),"바이트")
	// fmt.Println("int64형 사이즈 : ",unsafe.Sizeof(int_size16),"바이트")
	// fmt.Println("int64형 사이즈 : ",unsafe.Sizeof(int_size32),"바이트")
	// fmt.Println("int64형 사이즈 : ",unsafe.Sizeof(int_size64),"바이트")
	// fmt.Println("int64형 사이즈 : ",unsafe.Sizeof(uint_size64),"바이트")

	// var num = 5 // 우변에 있는 값을 추론
	// var num1 int8 = 5
	// fmt.Println("변수 num의 사이즈", unsafe.Sizeof(num)) // 메모리 낭비
	// fmt.Println("변수 num1의 사이즈", unsafe.Sizeof(num1))

	// num2 :=10 // 선언대입문 :   :=를 써서 var키워드, 타입을 생략
	// fmt.Println("var키워드와 타입 생략 방법", num2)

	// if num1==5 {
	// 	fmt.Println("asdf")
	// }
	

	// 타입변환(형변환)
	// num3 :=3
	// var fnum float64 = 3.5

	// var c int = int(fnum) // 타입이 다르기 때문에 대입도 되지 않는다.

	// res := num3 * fnum // 타입이 다르기 때문에 연산이 되지 않는다.
	
	// var e int64 = 7

	// res1 := num3  * e

	// 초기값을 넣지 않으면 아래와 같이 디폴트값이 들어감
	// 그 외 ..nil
	
}

// 고언어는 정적 동작
// 타입을 엄격하게 검사
// 가비지 컬렉터 메모리 회수 지원 