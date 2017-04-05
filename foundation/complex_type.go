package main

import (
	"fmt"
	"strings"
	"math"
	"sort"
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
	fmt.Println(s1)  // [1 2]	从下标0开始计数,直到end位，不含end位
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
	fmt.Println("funSlice1 操作后", slice1)	// 数组内为共用引用所以slice1中的2变成了2222
	funSlice2(&slice1)
	fmt.Println("funSlice2 操作后", slice1)	// &指向了slice1的指针，跨函数操作也操作slice1本身

	// slice 由函数 make 创建。这会分配一个全是零值的数组并且返回一个 slice 指向这个数组
	a := make([]int, 5)
	printSlice("a", a)	// [0 0 0 0 0]
	b := make([]int, 0, 5)
	printSlice("b", b)	// 指定size和cap容量
	c := b[:2]
	printSlice("c", c)	// [0 0]
	d := c[2:5]
	printSlice("d", d)	// [0 0 0]


	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	lll := make([]int, 5)
	lll = append(lll, 1,2,3,4,5,6,7)
	fmt.Printf("%s len:%d cap:%d %v\n", "lll",len(lll), cap(lll), lll)


	// 函数值 函数也可作为参数传递
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))
	fmt.Println(complexFunc(hypot))
	fmt.Println(complexFunc(math.Pow))

	// 闭包，作为一个函数值，引用了函数之外的变量，这个函数可以对引用的变量进行访问和赋值,换句话说，这个函数被“绑定”在这个变量上。
	var add,neg = adder(),adder()
	for i:=0;i<10 ;i++  {
		fmt.Println(i,add(i),neg(-2*i))
	}


	// 闭包实现fibonacci数列
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	// map集合
	map1 := make(map[string]string)

	// 创建
	map1["title"] = "张三"
	map1["count"] = "1"
	// 读取
	fmt.Println(map1["title"])
	// 更新
	map1["title"] = "李斯"
	fmt.Println(map1["title"])
	// 删除
	delete(map1, "title")

	// 遍历
	for key,value := range map1{
		fmt.Printf("%s %s \n", key, value)
	}

	// 验证map是否有序
	progameLanguages := map[string]int{
		"unix": 	0,
		"python": 	1,
		"go": 		2,
		"javascript": 	3,
		"testing": 	4,
		"philosophy": 	5,
		"startups": 	6,
		"productivity": 7,
		"hn": 		8,
		"reddit": 	9,
		"C++": 		10,
	}

	// 从go 1版本开始，map的遍历顺序是随机的
	for key,value := range progameLanguages  {
		fmt.Printf("%s %d\n", key, value)
	}

	var keys []string

	for k := range progameLanguages {
		keys = append(keys, k)
	}
	// 可以利用sort根据ASCII排序，再根据键读取值的特性
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", progameLanguages[k])
	}

}

func adder() func(int) int{
	var sum = 0
	return func(x int) int {
		sum += x
		fmt.Printf("sum:%d x:%d\n", sum, x)
		return sum
	}
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

func printSlice(s string, slice []int)  {
	fmt.Printf("%s len:%d cap:%d slice:%v\n", s, len(slice), cap(slice), slice)
}

func complexFunc(f func(float1,float2 float64) float64) float64 {
	return f(3,4)
}

var slice = make([]int,11)

func fibonacci() func() int {
	var i = 0
	return func() int{
		switch i {
		case 0:
			slice[0]=0
			i++
			return 0
		case 1:
			slice[1]=1
			i++
			return 1
		case 2:
			slice[2] = slice[i-1] + slice[i-2]
			i++
			return slice[2]
		default:
			var j = i
			slice[j] = slice[i-1] + slice[i-2]
			i++
			return slice[j]
		}
	}
}
