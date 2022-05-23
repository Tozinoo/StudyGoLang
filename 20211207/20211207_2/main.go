package main

import "fmt"

// 뒤에 er 붙이는게 관례
type Tester interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

// 빈 인터페이스

func PrintValue(v interface{}) { // 빈 인터페이스를 인수로 받는 함수
	switch t := v.(type) {
	case int:
		fmt.Printf(" int %d \n", int(t))
	case float64:
		fmt.Printf(" float %f \n", float64(t))
	case string:
		fmt.Printf(" string %s \n", string(t))
	default:
		fmt.Printf(" 위에 없는 다른 타입 : %T:%v\n", t, t)
	}
}

type People struct {
	Age int
}

type Mover interface {
	Moving()
}

func (s Student) String() string { // Student 구조체의 String() 메서드
	return fmt.Sprintf("하이 %s, %d살", s.Name, s.Age)
}

func main() {
	// interface : 구현을 포함하지 않는 메서드의 집합
	// 구체화된 타입 아닌 인터페이스만 가지고 메서드를 호출 -> 프로그램 요구사항 변경 시 유연하게 대처 가능
	// Go -> 인터페이스 구현 여부를 그 타입이 인터페이스에 해당하는 메서드를 가지고 있는지로 판단하는 덕 타이핑을 지원한다.

	// type 인터페이스이름 interface
	// 인터페이스도 타입 중 하나이다.
	// 인터페이스 변수 선언 가능. 변수의 값으로도 사용할 수 있다.
	// type DuckInterface interface {
	// 	// 메서드의 집함
	// 	// 주의 사항 : 메서드는 반드시 메서드 명이 있어야 함.
	// 	// 매개변수와 리턴이 다르더라도 이름이 같은 메서드는 안됨
	// 	// 인터페이스에서는 메서드 구현을 포함하지 않는다.
	// 	walk()
	// 	Move(distance int) int
	// }

	// type Test interface {
	// 	String() string
	// 	// String(a int) string
	// 	// _() string
	// }

	st := Student{"이순신", 10} //Student 타입
	var t Tester             //Test 타입

	t = st // t 값으로 st 대입

	fmt.Printf("%s\n", t.String())

	// 덕타이핑 : 인터페이스를 포함하고 있는지 여부를 확인할 때
	// 타입 선언 -> 인터페이스 구현 여부를 명시적으로 나타낼 필요 없이
	// 				인터페이스에 정의한 메서드 포함 여부만으로 결정하는 방식

	// 인터페이스의 또다른 기능
	// 1. 인터페이스 기본 값
	// 2. 빈 인터페이스
	// 3. 포함된 인터페이스

	type Reader interface {
		Read() (n int, err error)
		Close() error
	}

	type Writer interface {
		Write() (n int, err error)
		Close() error
	}

	type ReadWriter interface {
		Reader
		Writer
	}

	// 빈 인터페이스
	// 메서드를 들고 있지 않은 빈 녀석
	// 이녀석은 어떤 값이든 받을 수 있는 함수, 메서드, 변수 값을 만들 때 사용

	PrintValue(10)
	PrintValue(3.14)
	PrintValue("빈 인터페이스")
	PrintValue(People{100})

	// var m Mover
	//m.Moving()
}
