package main

import "fmt"

func func1() {
LABEL:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL // LABEL相当于标志，程序跳到此处执行
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

func func2() { // 相当于for循环
	i := 0
LABEL:
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto LABEL
}

func main() {
	func1()
	func2()
}
