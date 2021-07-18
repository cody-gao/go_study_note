package main

import "fmt"

func main() {
	test2(10, 0)
}

func test2(x, y int) {
	defer fmt.Println("dispose...") //defer常用来释放资源、解除锁定，或执行一些清理操作 //可定义多个defer，按FILO顺序执行

	println(x / y)
}
