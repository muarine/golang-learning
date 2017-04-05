package main

import (
	"fmt"
	"strings"
)

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

	// 二维数组，
	game := [][]string{
		[] string{"_","_","_"},
		[] string{"_","_","_"},
		[] string{"_","_","_"},
	}

	game[0][0] = "X"
	game[1][2] = "X"
	game[2][2] = "X"

	for i:= 0; i<len(game); i++ {
		fmt.Printf("%s\n", strings.Join(game[i]," "))
	}

	// slice  没有固定长度、内部由array指针、len长度、cap容量构成
	// append函数在现有长度下增长是翻倍的<1024, >1024后以1.25倍数增长
	// 选择slice元素,[start,end-1]

	s := []int{1,2,3,4,5,6}
	s1 := s[0:2]
	fmt.Println(s1)  // [1 2]
	s2 := s[4:]
	fmt.Println(s2)  // [5 6]
	s3 := s[:4]
	fmt.Println(s3)  // [1 2 3 4]

	s = append(s , 7)

	fmt.Println(s)


	// slice 考点
	slice1 := []int{1,2,3,4,5}
	fmt.Println("初始化的slice数组", slice1)
	funSlice1(slice1)
	fmt.Println("funSlice1 操作后", slice1)
	funSlice2(&slice1)
	fmt.Println("funSlice2 操作后", slice1)

}

func funSlice1(slice []int) {
	slice[1] = 2222
	slice = append(slice, 6666)
	fmt.Printf("funSlice1-->data:\t%#v\n", slice)
	fmt.Printf("funSlice1-->len:\t%#v\n", len(slice))
	fmt.Printf("funSlice1-->cap:\t%#v\n", cap(slice))
}

func funSlice2(slice2 *[]int){
	*slice2 = append(*slice2, 66666)
}