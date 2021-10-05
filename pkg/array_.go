/**
 * Created by GoLand
 * Time: 2021/9/26 12:00 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: array_.go
 * Desc:
 */
package pkg

import "fmt"

func ArrayDemo() {
	// 初始化长度为5的数组
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr_ := [5]string{0: "_", 3: "lk"}
	fmt.Println(arr1, arr_)

	// 初始化长度为5的数组(指针)
	arr2 := new([5]int)
	arr2[2] = arr1[2]
	calculate(arr2)
	//fmt.Println()
	fmt.Println(arr2)
}

// 参数为指针数组的函数
func calculate(a *[5]int) *[5]int {
	for k, v := range *a {
		a[k] = k + v
	}
	return a
}
