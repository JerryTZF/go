/**
 * Created by GoLand
 * Time: 2021/9/7 2:43 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: leetcode_basic.go
 * Desc:
 */
package main

import (
	"math"
	"strconv"
	"strings"
)

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

// 496、下一个更大的元素
func nextGreaterElement(nums1, nums2 []int) (res []int) {
	var t, p bool
	for i := 0; i < len(nums1); i++ {
		t, p = false, false
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				t = true
			}
			if t && nums2[j] > nums1[i] {
				res = append(res, nums2[j])
				p = true
				break
			}
		}

		if p == false {
			res = append(res, -1)
		}
	}
	return res
}

// 682、棒球比赛
func CalPoints(ops []string) (r int) {
	stack := make([]int, 0, len(ops))
	for i := range ops {
		switch ops[i] {
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		default:
			p, _ := strconv.Atoi(ops[i])
			stack = append(stack, p)
		}
	}
	for _, v := range stack {
		r += v
	}
	return r
}

// 844、比较含退格的字符串
func BackspaceCompare(s string, t string) bool {
	f := func(str string) []byte {
		bytes := make([]byte, 0)
		for i := range str {
			if str[i] == '#' && len(bytes) != 0 {
				bytes = bytes[:len(bytes)-1]
			}
			if str[i] != '#' {
				bytes = append(bytes, str[i])
			}
		}
		return bytes
	}
	return string(f(s)) == string(f(t))
}

// 1021、删除最外层括号
func RemoveOuterParentheses(s string) (str string) {
	stack, index := make([]byte, 0), 0
	for k := range s {
		stack = append(stack, s[k])
		if s[k] == '(' {
			index++
		}
		if s[k] == ')' {
			index--
		}
		if index == 0 {
			str += string(stack[1 : len(stack)-1])
			// 这里每次都会清空切片,会造成性能损失
			stack = []byte{}
		}
	}
	return
}

// 1021、删除最外层括号
// 力友方案
func RemoveOuterParentheses_(s string) string {
	var build strings.Builder
	var cnt int
	for _, l := range s {
		if l == '(' {
			if cnt > 0 {
				build.WriteRune(l)
			}
			cnt++
		} else {
			if cnt > 1 {
				build.WriteRune(l)
			}
			cnt--
		}
	}
	return build.String()
}

// 1441、用栈操作构建数组
func BuildArray(target []int, n int) []string {
	stack := make([]string, 0)
	for i := 1; i <= n; i++ {
		flag := false
		if len(target) > 0 {
			for k, v := range target {
				if v == i {
					// 命中则推入一个"Push
					stack = append(stack, "Push")
					// 移除命中的该元素,可以减少后续迭代次数
					target = append(target[:k], target[k+1:]...)
					// 标记该元素存在
					flag = true
					// 进入下一次迭代
					break
				}
			}
			if flag == false {
				stack = append(stack, []string{"Push", "Pop"}...)
			}
		} else {
			return stack
		}
	}
	return stack
}

// 3、无重复字符的最长子串(硬解)
func LengthOfLongestSubstring(s string) int {
	l, stack, length := 0, "", len(s)
	for i := 0; i < length; i++ {
		stack += string(s[i])
		for j := i + 1; j < length; j++ {
			if s[i] != s[j] && s[j] != s[j-1] && strings.IndexByte(stack, s[j]) == -1 {
				stack += string(s[j])
			} else {
				break
			}
		}
		if len(stack) > l {
			l = len(stack)
		}
		stack = ""
	}
	return l
}

// 3、无重复字符的最长子串(官解)
func LengthOfLongestSubstring_(s string) int {
	m, n, rk, ans := map[byte]int{}, len(s), -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = func(a, b int) int {
			if a < b {
				return b
			}
			return a
		}(ans, rk-i+1)
	}
	return ans
}

// 26、删除有序数组中的重复项(不允许出现拷贝,需要对原数组进行操作)
// 自解
func RemoveArrayDuplicates(nums []int) int {
	return len(recursive(nums))
}

// 递归会消耗较多的资源
func recursive(n []int) []int {
	flag := false
	for k := range n {
		if k == 0 {
			continue
		}
		if n[k] == n[k-1] {
			n = append(n[:k], n[k+1:]...)
			flag = true
			break
		}
	}
	if flag {
		return recursive(n)
	}
	return n
}

// 26、删除有序数组中的重复项(不允许出现拷贝,需要对原数组进行操作)
// 官解
func RemoveArrayDuplicates_(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 5、最长回文子串
func LongestPalindrome(s string) string {
	length, max, str := len(s), 0, string(s[0])
	if length == 1 {
		return s
	}
	for l := 2; l <= length; l++ {
		for i := 0; i <= length-l; i++ {
			ok, le := func(s string) (bool, int) {
				for head, tail := 0, len(s)-1; head < tail; head, tail = head+1, tail-1 {
					if s[head] != s[tail] {
						return false, 0
					}
				}
				return true, len(s)
			}(s[i : l+i])
			if ok && le > max {
				max, str = le, s[i:l+i]
			}
		}
	}
	return str
}

// 8. 字符串转换整数 (atoi)(没理解题意)
func MyAtoi(str string) int {
	return convert(clean(str))
}

func clean(s string) (sign int, abs string) {
	// 先去除首尾空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}
	// 判断第一个字符
	switch s[0] {
	// 有效的
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	// 有效的，正号
	case '+':
		sign, abs = 1, s[1:]
	// 有效的，负号
	case '-':
		sign, abs = -1, s[1:]
	// 无效的，当空字符处理，并且直接返回
	default:
		abs = ""
		return
	}
	for i, b := range abs {
		// 遍历第一波处理过的字符，如果直到第i个位置有效，那就取s[:i]，从头到这个有效的字符，剩下的就不管了，也就是break掉
		// 比如 s=123abc，那么就取123，也就是s[:3]
		if b < '0' || '9' < b {
			abs = abs[:i]
			// 一定要break，因为后面的就没用了
			break
		}
	}
	return
}

// 接收的输入是已经处理过的纯数字
func convert(sign int, absStr string) int {
	absNum := 0
	for _, b := range absStr {
		// b - '0' ==> 得到这个字符类型的数字的真实数值的绝对值
		absNum = absNum*10 + int(b-'0')
		// 检查溢出
		switch {
		case sign == 1 && absNum > math.MaxInt32:
			return math.MaxInt32
		// 这里和正数不一样的是，必须和负号相乘，也就是变成负数，否则永远走不到里面
		case sign == -1 && absNum*sign < math.MinInt32:
			return math.MinInt32
		}
	}
	return sign * absNum
}
