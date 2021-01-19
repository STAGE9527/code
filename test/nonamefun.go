package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	//Func就是一个全局匿名函数
	Func = func(a, b int) int {
		return a * b
	}
)

func addsum(a, b int) int {
	result := func(a1, b1 int) int {
		return a1 + b1
	}(a, b) // 定义时就调用
	return result

	//result := func(a1, b1 int) int {
	//	return a1 + b1
	//}
	//return result(a, b)
}

func testdefer(a, b int) {
	defer fmt.Println(a)
	defer fmt.Println(b)
	a++
	b++
	tmp := a + b
	fmt.Println(tmp)
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	fmt.Println(addsum(10, 10)) // 20
	res := Func(1, 3)
	fmt.Println(res)

	testdefer(1, 5)

	readFileContent("./test.txt")
}

func readFileContent(filename string) {
	if filename == "" {
		fmt.Println("请输入文件")
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	scaner := bufio.NewScanner(file)
	for scaner.Scan() {
		content := scaner.Text()
		if len(content) == 0 {
			continue
		}
		fmt.Println(content)
	}
}
