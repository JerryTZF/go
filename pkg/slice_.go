/**
 * Created by GoLand
 * Time: 2021/9/26 11:51 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: slice_.go
 * Desc: 切片相关
 */
package pkg

import "fmt"

// 切片初始化方式
func InitSlice() {
	s1 := make([]int, 5, 10)
	fmt.Printf("切片长度为%d,切片容量为%d,切片为:%v\n", len(s1), cap(s1), s1)

	s2 := []int{2, 4, 6, 8, 10}
	fmt.Printf("切片长度为%d,切片容量为%d,切片为:%v\n", len(s2), cap(s2), s2)

	s3 := []int{1: 4, 3: 8, 4: 10}
	fmt.Printf("切片长度为%d,切片容量为%d,切片为:%v\n", len(s3), cap(s3), s3)
}

// 切片取值
func SplitSlice() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := arr[2:5]
	fmt.Println("s1:", s1)
	s2 := arr[:5]
	fmt.Println("s2:", s2)
	s3 := arr[2:]
	fmt.Println("s3:", s3)
	s4 := arr[1:2:5] // 容量为5-1=4
	fmt.Printf("s4:%v,cap:%v,len:%v", s4, cap(s4), len(s4))
}

// 二维切片
func TwoDimensionalSlice() {
	s := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}}
	for k, v := range s {
		for _, j := range v {
			fmt.Printf("k:%d v:%d\n", k, j)
		}
	}
}

// 数组和切片关系
func Relationship() {
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[2:4]                // [3,4]
	fmt.Println("s1容量:", cap(s1)) // 3
	s1[0] = 999
	// 未超过切片的容量,固底层数组依旧为原数组,会保持引用.固改变切片会修改原数组
	// [1 2 999 4 5]
	fmt.Println("arr:", arr)

	// 动态追加s1切片,使其容量依旧小于原数组
	s1 = append(s1, []int{6, 7}...)
	// s1:[999 4 6 7];arr:[1 2 999 4 5]
	fmt.Printf("s1:%v;arr:%v\n", s1, arr)

	// 动态追加s1切片,使其容量大于原数组
	s1 = append(s1, []int{8}...)
	fmt.Printf("s1:%v;arr:%v\n", s1, arr)

	// 此时修改切片,已经不会对原数组进行修改了
	s1[0] = 888
	fmt.Printf("s1:%v;arr:%v\n", s1, arr)
}

// 判断两个切片是否相等
func IsSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for index, item := range b {
		if a[index] != item {
			return false
		}
	}

	return true
}

type customSlice struct {
	name string
	age  int
}

func CustomSlice() {
	s := []customSlice{
		{name: "小明", age: 12},
		{name: "小郑", age: 22},
	}
	fmt.Println(s)
}

// 切片反转参考 pkg.control_flow_.go

// 切片的append操作
func CustomAppend() {
	s1 := []string{"a", "b", "c", "d"}
	s2 := make([]string, len(s1))
	// 复制切片
	copy(s2, s1)
	fmt.Println(s2)

	// 向空切片追加(复制)
	s3 := append(make([]string, 0), s1...)
	fmt.Println(s3)

	// 删除切片最后一个元素
	s4 := append(s1[:len(s1)-1], s1[len(s1):]...)
	fmt.Println(s4)

	// 向切片的指定位置插入切片
	index := 2
	s5 := append(s1[:index], append([]string{"aa", "bb"}, s1[index:]...)...)
	fmt.Println(s5)

	// 删除指定位置的元素
	index_ := 1
	s6 := append(s1[:index_], s1[index_+1:]...)
	fmt.Println(s6)
}
