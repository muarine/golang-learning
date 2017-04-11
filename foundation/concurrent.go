package main

import (
	"time"
	"fmt"
)

// goroutine	轻量级线程，go f(...interface)
func say(s string) {
	for i:=0;i<5 ;i++  {
		time.Sleep(100* time.Millisecond)
		fmt.Println(s)
	}
}

// channel 是有类型的管道，可使用操作符<-控制数据的流向
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	// goroutine
	go say("hello")
	say("world")
	// channel
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	x := <-c	// 从 c 中获取 a数组前半段的和
	go sum(a[len(a)/2:], c)
	y := <-c 	// 从 c 中获取 a数组后半段的和
	//x,y := <-c,<-c // 此时 c 分为两个数据，可以把c理解为一个栈，最后的线程将和赋值给了x
	fmt.Println(x, y, x+y)
}
