package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)

	for i := 0; i < r.Len(); i++ {
		r.Value = 'A' + i
		r = r.Next()
	}
	for k := 0; k < r.Len(); k++ {
		fmt.Printf("%c", r.Value)
		r = r.Next()
	}
}
