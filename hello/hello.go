package main

import (
	"fmt"
	"math/cmplx"
)

func add2(x int) (z, v int) {
	z = x * 2
	v = x + 2
	return
}

// var 可定义在package和method下
var java, thrift bool
var java2, thrift2 = "java2", "thrift2"

var (
	ToBe   bool       = true
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("hello world \n")
	// fmt.Printf(stringutil.Reverse("\n !oG olleH"))
	print()
	println("aAAA")
	fmt.Println(add2(1))

	var num int
	var java3, thrift3, python = 1, true, "no!"
	var ss string
	fmt.Println("var定义ss变量不赋值，运行时自动初始化空字符串 ss == ''", ss == "")
	fmt.Println("num java thrift ：", num, java, thrift)
	fmt.Println("var 可定义在package内或者function内， java2，thrift2", java2, thrift2)
	fmt.Println("可初始化多个变量，直接指定变量值后不必指定数据类型, java3, thrift3,python", java3, thrift3, python)

	// := 必须在函数体内,简洁赋值语句在明确类型的地方，可以用于替代 var 定义
	k := 3
	fmt.Println(":= 简洁复制语句-在明确类型的地方 k =", k)

	/**
	基本数据类型
	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // uint8 的别名
	rune // int32 的别名
	float32 float64
	complex64 complex128
	*/

	const format = "%T(%v)\n"
	fmt.Printf(format, ToBe, ToBe)
	fmt.Printf(format, MaxInt, MaxInt)
	fmt.Printf(format, z, z)


}
