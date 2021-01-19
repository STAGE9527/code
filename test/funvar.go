package main

import "fmt"

// opFunc为自定义的类型名字，这里它是一个函数，接收两个值，返回一个值
type opFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

// op为变量名字，op_func为自己定义的类型
func operator(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func operator2(op opFunc, a, b int) int {
	return op(a, b)
}

func main() {
	xx := add
	result := operator(xx, 10, 10)
	fmt.Println(result) // 20

	result2 := operator2(xx, 10, 10)
	fmt.Println(result2) // 20
}
