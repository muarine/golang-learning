// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"testing"
	"fmt"
)

type Student struct {
	age 		int
	name 		string
	hasDelete 	bool
	hash		int16
}

func sum(array []int, result chan int){

	sum := 0

	for _,e := range array{
		sum += e
	}

	result <- sum

}

// goroutine并不是和线程一一对应，它是一种比线程更轻量级的实现，十几个goroutine可能在底层就是几个线程
func TestGoroutine(t *testing.T){
	array := []int{2, 3, 5, 6, 10, -5, 1, 0}

	result := make(chan int, 2)

	// 数组后半部分切片
	go sum(array[:len(array)/2], result)
	// 数组前半部分切片
	go sum(array[len(array)/2:], result)

	x,y := <-result,<-result

	fmt.Println(x,y, x+y)
}
