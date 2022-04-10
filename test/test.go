package test

import "fmt"

/**
委托模式：将自己的某个功能，交由受委托者完成，

必须写在同一个包内，以免包的循环引入
*/
type A struct {
	B
}

func (a A) Func1() {
	fmt.Println("A.func1")
}

type B struct {
	name string
}

func (b B) Func1() {
	fmt.Println("B.func1")
}

func (b B) Func2() {
	fmt.Println("B.func2")
}

func (b B) Func3() {
	fmt.Println("B.func3")
}

func (b B) FuncCom() {
	b.Func1()
	fmt.Println("B.funcCom")
	b.Func3()
}
