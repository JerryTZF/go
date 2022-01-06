/**
 * Created by GoLand
 * Time: 2022/1/6 3:38 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: reverse_int_7.go
 * Desc: https://leetcode-cn.com/problems/reverse-integer/
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
}

// 7 整数反转
func reverse(x int) (r int) {
	for x != 0 {
		if r < math.MinInt32/10 || r > math.MaxInt32/10 {
			return 0
		}
		l := x % 10
		x /= 10
		r = r*10 + l
	}
	return
}
