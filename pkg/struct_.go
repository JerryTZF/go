/**
 * Created by GoLand
 * Time: 2021/10/3 12:28 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: struct_.go
 * Desc: 结构体相关操作
 */
package pkg

import "fmt"

// struct可以理解为轻量的(没有成员方法)类
type Person struct {
	name  string
	age   int
	hobby string
}

func (p *Person) speak() string {
	return "My name is " + p.name
}

func (p *Person) walk(where string) bool {
	if where == "China.Henan.he" {
		return true
	}
	return false
}

func CallMemberFunc() {
	// 类似new一个对象，然后调用其成员方法
	Leo := new(Person)
	Leo.age = 27
	Leo.name = "Leo"
	Leo.hobby = "speak & walk"
	var talk = Leo.speak()
	fmt.Println(talk)
	var walk = Leo.walk("China")
	fmt.Println(walk)
}

// 链表、二叉树等数据结构
// 用法可以参考LeetCode的二叉树和链表相关算法
type TreeNode struct {
	left  *TreeNode
	data  int
	right *TreeNode
}

func StructDemo() {
	// 三种定义初始化方式
	s1 := new(Person)
	s2 := &Person{}
	s3 := Person{}
	// s1和s2等价
	// 使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针
	fmt.Println(s1, s2, s3)

	s1.name = "Jerry"
	s1.hobby = "PingPang"
	s1.age = 27

	s3.name = "Tom"
	s3.hobby = "Basketball"
	s3.age = 25

	fmt.Println(s1, s3)

	// 初始化方式
	var s4 Person
	s4.name = "Iris"
	s4.age = 17
	s4.hobby = "TableBall"

	fmt.Println(s4)

	// 初始化方式
	var s5 = Person{
		"Kitty",
		28,
		"GunShooter",
	}

	fmt.Println(s5)

	// 匿名初始化方式
	s6 := &struct {
		A1 int
		A2 float64
		A3 string
		A4 struct{ A5 string }
	}{
		A1: 23,
		A2: 23.23,
		A3: "hello world",
		A4: struct {
			A5 string
		}{"nihao"},
	}

	fmt.Println(s6)
}

// 结构体中的匿名字段、嵌套结构体
type Inner struct {
	int1 int
	int
	Person
}

// 类型的成员方法(除接口、指针外其他任一类型都可以有成员方法)
// 1、 一个类型的成员方法是不允许重载的
// 2、类型的方法必须要在同一个包内
func CallMemberFunc_() {
	o1 := new(Inner)
	o1.int = 10
	o1.int1 = 5
	t1 := o1.B(5)
	//t2 := o1.B(10)
	fmt.Println(t1, o1)
	//fmt.Println(t2, o1)
}

func (i *Inner) A() int {
	return i.int + i.int1
}

func (i *Inner) B(n int) int {
	i.int1 = i.int + i.int1 + n
	return i.int
}

// 任意类型添加成员方法
type MyString string

func (s MyString) IsLargeString() bool {
	if len(s) > 10 {
		return true
	}
	return false
}
