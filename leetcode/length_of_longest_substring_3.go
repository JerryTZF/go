/**
 * Created by GoLand
 * Time: 2022/1/10 5:13 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: length_of_longest_substring_3.go
 * Desc: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

// 最长子串
func lengthOfLongestSubstring(s string) int {
	r, l, st := 0, len(s), []byte{}
	for i := 0; i < l; i++ {
		st = append(st, s[i])
		for j := i + 1; j < l; j++ {
			// 两两不等 && j元素不在栈内
			if s[i] != s[j] && s[j] != s[j-1] && strings.IndexByte(string(st), s[j]) == -1 {
				st = append(st, s[j])
			} else {
				break
			}
		}
		if r < len(st) {
			r = len(st)
		}

		st = make([]byte, 0)
	}
	return r
}
