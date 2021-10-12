/**
 * Created by GoLand
 * Time: 2021/10/12 5:06 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: palindrome_list.go
 * Desc: 234、是否是回文链表
 */
package leetcode

// 是否是回文链表
func isPalindromeList(l *listNode) bool {
	var stack []int
	for {
		stack = append(stack, l.Val)
		if l.Next == nil {
			break
		} else {
			l = l.Next
		}
	}

	// 双向指针
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		if stack[i] != stack[j] {
			return false
		}
	}

	return true
}
