package main

import "fmt"

//基本类型学习

// 分别声明
var name1 string
var age1 string

// 批量声明
var (
	name2 string
	age2  int
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("hello wold", arr)
}
