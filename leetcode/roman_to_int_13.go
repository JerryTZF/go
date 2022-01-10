/**
 * Created by GoLand
 * Time: 2022/1/10 2:09 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: roman_to_int_13.go
 * Desc: https://leetcode-cn.com/problems/roman-to-integer/
 */
package main

import "fmt"

func main() {
	fmt.Println(romanToInt("IV"))
}

// 罗马数字转整数
func romanToInt(s string) int {
	var n = 0
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for k := range s {
		if k+1 < len(s) && m[s[k]] < m[s[k+1]] {
			n -= m[s[k]]
		} else {
			n += m[s[k]]
		}
	}
	return n
}
