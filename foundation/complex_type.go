package main

import "fmt"

func main() {

	// 指针 pointer

	i,j := 28,49
	p := &i			// & 符号会生成一个指向其作用对象的指针, p指针指向了i
	fmt.Println("p指针的内存地址值 ", p)		// 打印 p指针的内存地址

	*p = 21
	fmt.Println("p指针指向了i，对p赋值其实是对作用对象i赋值 ", i)
	fmt.Println("p指针指向的内存地址值 ", p)

	p = &j
	*p = *p / 7
	fmt.Println("p指针指向了j，对p操作实际上是对作用对象j操作 ", j)
	fmt.Println("此时p指针指向的内存地址值应该发生了变化 ", p)


	// 结构体
	type Student struct {
		name string
		age uint8
		sex string
	}
	student := Student{"zhangsan",18, "男"}
	student.name = "lisi"
	fmt.Println(student)

	// 结构体字段可以通过结构体指针来访问，指针间接的访问是透明的
	ss := &student
	ss.name = "wangwu"
	ss.age = 21
	ss.sex = "女"
	fmt.Println(ss)

}
