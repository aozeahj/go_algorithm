package test

import (
	"testing"
)

func TestA_sayHi(t *testing.T) {
	a := new(A)
	a.FuncCom()
}

func K() {
	a := 0
	for i := 0; i < 10; i++ {
		go func() {
			a++
			println(a)
		}()
	}
}
