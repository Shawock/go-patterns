package strategy

import "fmt"

// 定义一个接口，这个接口能够处理两个整数，返回一个整数，如何处理由实现者决定
type Operator interface {
	Apply(left, right int) int
}

// 实现者1：相加
type Addition struct {
}

func (Addition) Apply(left, right int) int {
	fmt.Println("execute Addition")
	return left + right
}

// 实现者2：相乘
type Multiplication struct {
}

func (Multiplication) Apply(left, right int) int {
	fmt.Println("execute Multiplication")
	return left * right
}

// 需要一个统一的对象接收接口，这样可以动态改变接口而不用改方法签名
type Operation struct {
	optr Operator
}

func (optn Operation) Operate(left, right int) int {
	return optn.optr.Apply(left, right)
}
