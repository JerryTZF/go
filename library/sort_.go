/**
 * Created by GoLand
 * Time: 2021/9/27 10:34 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: sort_.go
 * Desc: 排序的标准库
 */
package library

import (
	"fmt"
	"sort"
)

// 常用类型内置排序方法
func BasicSort() {
	// 整型切片
	s1 := []int{2, 5, 3, 7, 1, 2}
	sort.Ints(s1)                // 升序排序
	sort.Sort(sort.IntSlice(s1)) // 升序排序(同上)
	fmt.Println(s1)
	sort.Sort(sort.Reverse(sort.IntSlice(s1))) // 降序排序
	fmt.Println(s1)

	// 字符串切片
	s2 := []string{"aJ", "cB", "bA", "eE", "dQ"}
	sort.Strings(s2)                // 升序排序
	sort.Sort(sort.StringSlice(s2)) // 升序排序(同上)
	fmt.Println(s2)
	sort.Sort(sort.Reverse(sort.StringSlice(s2))) // 降序排序
	fmt.Println(s2)

	// 浮点型切片
	s3 := []float64{1.12, 1.11, 1.13, 1.15, 1.14}
	sort.Float64s(s3)                // 升序排序
	sort.Sort(sort.Float64Slice(s3)) // 升序排序(同上)
	fmt.Println(s3)
	sort.Sort(sort.Reverse(sort.Float64Slice(s3))) // 降序排序
	fmt.Println(s3)

	// 字符型切片
	s4 := []byte{'a', 'A', 'c', '1'}
	sort.Slice(s4, func(i, j int) bool { return s4[i] < s4[j] }) // 升序排序
	fmt.Println(s4)
	sort.Slice(s4, func(i, j int) bool { return s4[i] > s4[j] }) // 降序排序
}

// 结构体排序
func StructSort() {
	family := []struct {
		name string
		age  int
	}{
		{"Jerry", 13},
		{"Loe", 15},
		{"Iris", 11},
	}

	// 年龄升序排序
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].age < family[j].age
	})
	fmt.Println(family)

	// 名字降序
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].name > family[j].name
	})
	fmt.Println(family)
}

// 实现接口的排序(相对比较自定义)
type person struct {
	name string
	age  int
}

type sortByAge []person

func (s sortByAge) Len() int { return len(s) }

func (s sortByAge) Less(i, j int) bool { return s[i].age < s[j].age }

func (s sortByAge) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func StructSortByInterface() {
	family := []person{{"Alice", 11}, {"Jerry", 27}, {"Blue", 12}}
	sort.Sort(sortByAge(family))
	fmt.Println(family)
}
