package main

import (
	"math"
	"fmt"
	"time"
)

// 1 结构体类型定义(指针接收)方法
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 2 內建接口

type Test struct {
	when time.Time
	why string
}

func (t Test) String() string{
	return fmt.Sprintf("when:%s,why:%s\n", t.when.String(), t.why)
}

func (t *Test) Error() string {
	return fmt.Sprintf("%v,%s",t.when, t.why)
}

// error类型
func run() error {

	return &Test{time.Now(), "didn't work"}
}
func main() {
	// 结构体类型上定义方法
	// Scale和Abs方法都是用指针接收
	v := &Vertex{3,5}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

	// 內建接口 string error
	t := Test{time.Now(),"为什么呢"}
	fmt.Printf("%s", t)

	if err := run(); err != nil {
		fmt.Println(err)
	}

	//
}
