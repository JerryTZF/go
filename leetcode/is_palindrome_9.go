/**
 * Created by GoLand
 * Time: 2022/1/10 11:54 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: is_palindrome_9.go
 * Desc: https://leetcode-cn.com/problems/palindrome-number/
 */
package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(123321))
}

// 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	r := 0
	x_ := x
	for x != 0 {
		l := x % 10

		x /= 10
		r = r*10 + l
	}

	return r == x_
}