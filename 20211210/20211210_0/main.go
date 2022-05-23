package main

import "fmt"

/*
테스트 코드를 작성하는 규칙
1. _test.go로 끝나야함
2. func TestAdd(t*testing.T)
3. testing package import

*/

func square(x int) int {
	return x * x
}

func main() {
	fmt.Printf("결과 : %d\n", square(9))

}
