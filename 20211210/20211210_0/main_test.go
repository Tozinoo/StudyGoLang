package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("ㅇㅇㅇ %d", rst)
	}
}
