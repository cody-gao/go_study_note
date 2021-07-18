package main

import (
	"fmt"
)

type user struct { //结构体类型
	name string
	age  int
	sex  string
}

type manager struct {
	user  //匿名嵌入其他类型
	title string
}

func main() {
	var m manager

	m.name = "cody" //直接访问匿名字段的成员
	m.age = 18
	m.sex = "男"

	fmt.Println(m)
}
