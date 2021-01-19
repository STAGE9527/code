package main

import "fmt"

// 方式1：交换两个变量的值
func swap1(a, b *int) {
	*a, *b = *b, *a
}

// 方式2：交换两个变量的值
func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	var i int = 10
	var p *int = &i // *int代表指针，&a代表获取a的地址
	*p = 20         // 给指针所指向的值赋值
	fmt.Println(i)  // 20

	a, b := 11, 22
	swap1(&a, &b)
	fmt.Println(a, b) // 22 11

	c, d := 33, 44
	c, d = swap2(c, d)
	fmt.Println(c, d) // 44 33

	var ii int
	ii = 10
	fmt.Printf("i的地址：%p\n", &ii)

	var pointer *int
	pointer = &ii
	fmt.Printf("pointer是一个指针变量：%v\n", pointer)
	fmt.Printf("pointer的地址：%p\n", &pointer)
	fmt.Printf("pointer指向的值：%v\n", *pointer)
}
