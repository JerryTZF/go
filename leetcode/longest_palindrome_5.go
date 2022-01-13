/**
 * Created by GoLand
 * Time: 2022/1/13 1:56 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: longest_palindrome_5.go
 * Desc: https://leetcode-cn.com/problems/longest-palindromic-substring/
 */
package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("aca"))
}

// 5、最长回文子串
func longestPalindrome(s string) string {
	l, rl, r := len(s), 0, string(s[0])
	if l == 1 {
		return s
	}
	for i := 2; i <= l; i++ {
		for j := 0; j <= l-i; j++ {
			ok, rl_ := func(s string) (bool, int) {
				// 双指针判断是否回文
				for h, t := 0, len(s)-1; h < t; h, t = h+1, t-1 {
					if s[h] != s[t] {
						return false, 0
					}
				}
				return true, len(s)
			}(s[j : j+i])
			if ok && rl_ > rl {
				rl, r = rl_, s[j:i+j]
			}
		}
	}
	return r
}