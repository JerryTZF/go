/**
 * Created by GoLand
 * Time: 2021/9/7 2:43 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: leetcode_basic.go
 * Desc:
 */
package leetcode

import (
	"math"
	"sort"
	"strings"
)

// 7 整数反转
func Reverse(x int) (r int) {
	for x != 0 {
		if r < math.MinInt32/10 || r > math.MaxInt32/10 {
			return 0
		}
		l := x % 10
		x /= 10
		r = r*10 + l
	}
	return r
}

// 14 最长公共前缀
func LongestCommonPrefix(s []string) string {
	sort.Strings(s)
	a, z := s[0], s[len(s)-1]
	for i := len(a); i >= 0; i-- {
		t := a[0:i]
		if strings.Index(z, t) == 0 {
			return t
		}
	}

	return ""
}

// 13、罗马数字转整数
func RomanToInt(s string) int {
	var n = 0
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for k := range s {
		if m[s[k]] < m[s[k+1]] && k+1 < len(s) {
			n -= m[s[k]]
		} else {
			n += m[s[k]]
		}
	}
	return n
}

// 9、回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	r := 0
	for x > r {
		r = r*10 + x%10
		x /= 10
	}
	return x == r || x == r/10
}

// 1、两数之和
func twoSum(n []int, t int) []int {
	for k, v := range n {
		for i := k + 1; i < len(n); i++ {
			if v+n[i] == t {
				return []int{k, i}
			}
		}
	}

	return nil
}

// 1002、查找共用字符
func CommonChars(s []string) (r []string) {
	var a, n string
	for k := range s {
		if k+1 == len(s) {
			break
		}

		if k == 0 {
			a = s[k]
		} else {
			a = n
		}

		n = func(s1, s2 string) (n string) {
			for _, v := range s1 {
				if index := strings.Index(s2, string(v)); index != -1 {
					s2 = string(append([]byte(s2)[:index], []byte(s2)[index+1:]...))
					n += string(v)
				}
			}
			return
		}(a, s[k+1])
	}

	if n == "" {
		return
	} else {
		for _, v := range n {
			r = append(r, string(v))
		}
		return
	}
}

// 1047. 删除字符串中的所有相邻重复项
func RemoveDuplicates(s string) string {
	stack := make([]byte, 0)
	for k := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[k] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[k])
		}
	}
	return string(stack)
}

// 20、有效的括号
func isValid(s string) bool {
	stack := make([]byte, 0)
	hashTable := map[byte]byte{'{': '}', '[': ']', '(': ')'}
	for k := range s {
		// 入栈
		if len(stack) == 0 || s[k] == '{' || s[k] == '(' || s[k] == '[' {
			stack = append(stack, s[k])
		}

		// 出栈
		if s[k] == '}' || s[k] == ']' || s[k] == ')' {
			if s[k] == hashTable[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

