package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main()  {

	sum := 0

	for i := 0 ; i < 10; i++ {
		sum += i
	}

	fmt.Println("只有for循环一种遍历方式，包括循环初始化语句" , sum);

	sum = 1
	for ; sum < 1000;  {
		sum += sum
	}

	fmt.Println("for循环初始化语句和后置语句是可选的" , sum)

	sum = 1
	for sum < 1000  {
		sum += sum
	}
	fmt.Println("可以省略分号，for也就相当于 C语言中的while" , sum)

	// 死循环的写法
	//for{
	//	sum += sum
	//}

	// if 控制关键字
	fmt.Println(
		"在条件判断之前可执行一个初始化语句",
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	switch os := runtime.GOOS ; os {
	case "darwin":
		fmt.Println("OS X.")
		// fallthrough case分支不会自动终止，而是强制执行下一个case
		fallthrough
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}



	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}


	// defer 延迟函数执行，直到上层函数执行完毕，需要注意的是延迟函数调用的参数会立刻生成
	num := 0
	defer fmt.Println(" world ", num + 1)
	defer fmt.Print(",")
	fmt.Println("hello ")

	fmt.Println("print for loop")
	// 使用defer 会把函数压入一个栈中
	for i:= 0;i< 10 ; i++ {
		defer fmt.Println(i)
	}
}


func pow(x, n, lim float64) float64 {
	// if可在条件之前执行一个初始化语句，作用域只在if内可见
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}