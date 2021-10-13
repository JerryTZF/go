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
	"golang-study/pkg"
)

func main() {
	//library.StrIndex()
	if r, e := pkg.ErrorReturned(0); e == nil {
		fmt.Println(r)
	} else {
		fmt.Println(r, e)
	}
	a := pkg.DeferError_(10, 0)
	fmt.Println(a)
	//library.ListDemo()
	leetcode.CalPoints([]string{"5", "2", "C", "D", "+"})
}
