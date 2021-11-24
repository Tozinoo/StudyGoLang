package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func PrintScore(name string, math int, eng int, science int)(result string){
	total := math+eng+science
	avg:= total/3
	if avg>=90 {
		result = "A"
	} else if avg>=80 && avg<90 {
		result = "B"
	} else if avg>=70 && avg<80 {
		result = "C"
	} else if avg>=60 && avg<70 {
		result = "D"
	} else {
		result ="D"
	}
	 fmt.Printf("%s의 학점은 %s입니다.\n",name,result)
	 return
}

// 멀티반환 함수(여러개 리턴 값을 가지고 있음)
// 매개변수 타입이 같으면 아래와 생략가능
func Divide(a,b int)(result int, success bool){
	if b==0 {
		result =0 
		success = false
		return // 명시적으로 반환할 것을 지정하지 않음
	}
	result = a/b
	success = true
	return 
}

func main() {
	//c := Add(3, 6)
	//fmt.Println(c)
	
	PrintScore("홍길동", 80,90,100)
	PrintScore("배추도사", 100,90,100)
	PrintScore("은비까비", 50,90,60)

	
}