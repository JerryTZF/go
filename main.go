/**
 * Created by GoLand
 * Time: 2021/9/7 1:58 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: main.go
 * Desc:
 */
package main

import (
	"fmt"
	"golang-study/leetcode"
)

func main() {
	//library.ReflectDemo()
	//n := leetcode.RemoveArrayDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 4})
	//fmt.Println(n)
	//library.SafeMap()
	//ss := "babqw"
	//fmt.Println(ss[1:3])
	s := leetcode.LongestPalindrome("abcda")
	fmt.Println(s)
	//is, l := leetcode.HuiWenTest("a")
	//fmt.Println(is, l)
	//leetcode.CalPoints([]string{"5", "2", "C", "D", "+"})
	//t := leetcode.BuildArray([]int{1, 2}, 4)
	//fmt.Println(t)
}
