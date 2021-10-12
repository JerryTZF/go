/**
 * Created by GoLand
 * Time: 2021/10/9 2:01 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: interface_.go
 * Desc:
 */
package pkg

import "fmt"

// 定义接口
type Shaper interface {
	Area() float64
}

// 实现接口的类型
type Square struct {
	side float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (s *Square) Who() string {
	return "正方形面积"
}

type Rectangle struct {
	length, weight float64
}

func (r *Rectangle) Area() float64 {
	return r.weight * r.length
}

func (r *Rectangle) Who() string {
	return "长方形面积"
}

func InterfaceDemo() {
	c := &Rectangle{weight: 10, length: 20}
	z := &Square{side: 5}

	s := []Shaper{c, z}
	// "多态"
	for _, v := range s {
		fmt.Println(v.Area())
	}
}

// 一个接口被多个类型实现
type ValueAble interface {
	getValue() float64
}

type A struct {
	Count int
	Unit  float64
}

type B struct {
	Price float64
}

func (i *A) getValue() float64 {
	return i.Unit * float64(i.Count)
}

func (i *B) getValue() float64 {
	return i.Price
}

// 只要是实现了 ValueAble 的类型都可以作为参数，类似Python的鸭子类型
func showValue(a ValueAble) {
	fmt.Printf("value of a is: %.2f", a.getValue())
}

func ShowValue() {
	// 实现接口的类型变量可以赋值给对应的接口类型
	var o ValueAble = &A{Count: 10, Unit: 12.21}
	showValue(o)
}

// 一个类型实现多个接口
type AllImplement struct {
	Side float64
}

// 实现两个接口
func (i *AllImplement) Area() float64 {
	return i.Side * i.Side
}

func (i *AllImplement) getValue() float64 {
	return i.Side
}

// 接口嵌套
type I1 interface {
	F1() string
}

type I2 interface {
	F2() string
}

type I3 interface {
	I1
	I2
	F3() string
}

// 类型断言
// 因为在 Golang 中，`接口`变量的动态类型是变化的，有时我们需要知道一个接口 变量 的动态 类型 究竟是什么，
// 这就需要使用类型断言，类型断言就是对接口变量的类型进行检查
func IsInstance() {
	var o Shaper = &Square{side: 5}
	var a Shaper = &Rectangle{length: 10, weight: 23.22}
	if t, ok := o.(*Square); ok {
		fmt.Printf("接口变量`o`的类型是：%T\n", t)
	} else {
		fmt.Printf("接口变量`o`的类型不是*Square")
	}

	if u, ok := a.(*Square); ok {
		fmt.Printf("接口变量`u`的类型是：%T\n", u)
	} else {
		fmt.Printf("接口变量`u`的类型不是*Square")
	}

	// switch-type 判断类型
	switch a.(type) {
	case *Square:
		fmt.Println("a.type is *Square")
	case *Rectangle:
		fmt.Println("a.type is *Rectangle")
	default:
		fmt.Println("a.type is unknown")
	}
}
