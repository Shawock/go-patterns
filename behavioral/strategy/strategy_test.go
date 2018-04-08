package strategy

import "fmt"

func ExampleAddition_Apply() {
	optn := Operation{Addition{}}
	result := optn.Operate(10, 20)
	fmt.Println(result)
	// Output:
	// execute Addition
	// 30
}

func ExampleMultiplication_Apply() {
	optn := Operation{Multiplication{}}
	result := optn.Operate(10, 20)
	fmt.Println(result)
	// Output:
	// execute Multiplication
	// 200
}