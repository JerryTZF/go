/**
 * Created by GoLand
 * Time: 2021/9/29 11:03 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: func_.go
 * Desc: 函数的基本知识
 */
package pkg

import (
	"fmt"
	"log"
)

// 在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）

// 变长参数(s为一个[]int类型的slice)
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}

// defer操作
func DeferDemo() {
	defer func() {
		log.Panicln("defer triggered")
	}()
}

// 递归函数(recursive)
// 斐波那契数列
func Caller() {
	r := 0
	for i := 0; i < 10; i++ {
		r = recursive(i)
		fmt.Printf("第%d个数是:%d\n", i, r)
	}
}

func recursive(n int) (r int) {
	if n <= 1 {
		r = 1
	} else {
		r = recursive(n-1) + recursive(n-2)
	}
	return
}

// 函数作为参数
// 也可以参考strings.IndexFunc(s string, f func(rune) bool) int
type add func(a, b int) int

func callback(n int, fn add) int {
	return fn(n, 10)
}

func fn(n1, n2 int) int {
	return n1 + n2
}

func Caller_(n int) int {
	return callback(n, func(a, b int) int {
		return a - b
	})
	//return callback(n, fn)
}

// 将函数作为返回值
func Add2() func(n int) int {
	return func(n int) int {
		return n + 2
	}
}

func Adder(a int) func(n int) int {
	return func(n int) int {
		return a + n
	}
}

// ####格外注意####
// 在闭包中使用到的变量可以是在闭包函数体内声明的，也可以是在外部函数声明的
func Adder_() func(n int) int {
	var x int
	return func(n int) int {
		x += n
		return x
	}
}

func Caller__() {
	fn1 := Add2()
	fmt.Println(fn1(10))

	fn2 := Adder(10)
	fmt.Println(fn2(10))
	fmt.Println(fn2(10))

	// ####格外注意####
	// 在闭包中使用到的变量可以是在闭包函数体内声明的，也可以是在外部函数声明的
	fn3 := Adder_()
	fmt.Println(fn3(1)) // 1
	fmt.Println(fn3(1)) // 2
	fmt.Println(fn3(1)) // 3
}
