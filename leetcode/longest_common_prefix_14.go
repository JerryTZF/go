/**
 * Created by GoLand
 * Time: 2022/1/6 3:44 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: longest_common_prefix_14.go
 * Desc:
 */
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flight", "flow"}))
}

// 14 最长公共前缀
func longestCommonPrefix(s []string) string {
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
