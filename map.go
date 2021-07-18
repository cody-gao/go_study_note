package main

import "fmt"

func main() {
	//创建字典型对象
	m := make(map[string]int)

	//向字典填充数据
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	//使用ok-idiom 获取值，可知道key/value是否存在
	x, ok := m["c"] //所谓 ok-idiom 模式，是指在多返回值中用一个名为ok的布尔值来标示操作是否成功。因为很多操作默认返回0值，所以需额外说明
	if false == ok {
		fmt.Println(x, ok)
	}

	for _, v := range m {
		fmt.Println("之前的值为", v)
	}

	delete(m, "b")
	fmt.Println("删除元素", "b")

	fmt.Println("字典的长度为", len(m))

	for _, v2 := range m {
		fmt.Println("之后的值为", v2)
	}

}
