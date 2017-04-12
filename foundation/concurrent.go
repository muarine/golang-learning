package main

import (
	"time"
	"fmt"
	"sync"
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

// 针对 channel 的操作range和close
func fibonacci2(n int, c chan int)  {
	x,y := 0,1
	for i:=0;i<n ;i++  {
		c <- x
		x,y = y,x+y
	}
	// 发送方关闭channel
	close(c)
}

// 互斥锁
// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	// 将UnLock操作压入栈顶，直到return前才会执行UnLock
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	// goroutine
	go say("hello")
	say("world")
	// channel 有类型区分
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	//x := <-c	// 从 c 中获取 a数组前半段的和
	go sum(a[len(a)/2:], c)
	//y := <-c 	// 从 c 中获取 a数组后半段的和
	x,y := <-c,<-c // 此时 c 分为两个数据，可以把c理解为一个栈，最后的线程将和赋值给了x
	fmt.Println(x, y, x+y)
	// channel 带缓冲
	ch := make(chan int , 2)
	ch <- 2
	ch <- 1
	fmt.Println(<-ch);
	fmt.Println(<-ch);

	// channel 操作close和range
	cc := make(chan int , 10)
	go fibonacci2(cap(cc), cc)
	// range会不断从channel里接收值
	for i := range cc {
		fmt.Println(i)
	}
	// 判定channel是否已关闭，关闭操作由生产者处理
	v,ok := <-cc
	fmt.Println(v)
	fmt.Println(ok)

	// 互斥锁，sync.Mutex Lock UnLock
	ccc := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go ccc.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(ccc.Value("somekey"))

}
