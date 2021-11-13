/**
 * Created by GoLand
 * Time: 2021/10/3 12:28 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: struct_.go
 * Desc: 结构体相关操作
 */
package pkg

import (
	"fmt"
	"strconv"
)

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

// 如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。
// 还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法
func (p *Person) String() string {
	return "(" + p.name + "_" + p.hobby + "_" + strconv.Itoa(p.age) + ")"
}

func StringParse() {
	p := &Person{name: "Toms", age: 18, hobby: "Iris"}
	fmt.Println(p)
}

// 结构体、集合、高阶函数
type Any interface{}

type Car struct {
	Model        string
	Manufacturer string
	BuilderYear  int
}

type Cars []*Car

// 对切片中的每辆车做处理:f(c)
func (cs Cars) Process(f1 func(car *Car)) {
	for _, c := range cs {
		f1(c)
	}
}

func (cs Cars) FindAll(f2 func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f2(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}

	return appender, sortedCars
}

func Run() {
	// make some cars
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	// query
	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuilderYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}

// ************
// 演示多继承
// ************
type Employee struct {
	Title  string
	Salary int
}

type Son struct {
	Mother string
	Father string
}

type Jerry struct {
	Ability string
	Hobby   string
}

type Me struct {
	*Employee
	*Jerry
	Son
}

func (j *Jerry) ChangeHobby(h string) {
	j.Hobby = h
}

func DemoForExtend() {
	e := &Employee{Title: "技术组长", Salary: 15000}
	j := &Jerry{Ability: "PHP", Hobby: "乒乓球"}
	s := Son{Mother: "王", Father: "田"}

	me := &Me{e, j, s}

	// 一定注意:这里不要并发修改hobby
	me.ChangeHobby("网球")
	j.ChangeHobby("不是网球")
	fmt.Println(me.Hobby)
}
