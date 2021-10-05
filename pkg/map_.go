/**
 * Created by GoLand
 * Time: 2021/9/26 12:00 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: map_.go
 * Desc:
 */
package pkg

import (
	"fmt"
	"sort"
	"strconv"
)

// map的常规操作
func NormalMapAction() {
	m1 := make(map[string]string)
	for k, v := range []string{"aa", "bb", "cc"} {
		m1[v] = v + strconv.Itoa(k)
	}

	fmt.Println(m1)

	// 不存在的key对应的值的判断
	if v, ok := m1["cc"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("下标hh不存在")
	}

	// 删除指定键值对
	delete(m1, "aa")
	fmt.Println(m1)

	// 复合类型：Slice的每一个元素为Map
	// Map的循环出来的顺序并不是声明好的顺序
	s1 := []map[int]string{{1: "one"}, {2: "two"}, {3: "three"}}
	for k := range s1 {
		fmt.Println(s1[k][k+1])
	}

	// 复合类型：Map的value里面为一个Slice
	m2 := map[int][]string{1: {"a", "b"}, 2: {"c", "d"}}
	for k := range m2 {
		fmt.Println(m2[k][1])
	}

}

// Map的排序
func SortMap() {
	m := map[string]int{"z": 1, "y": 2, "m": 3, "a": 4}

	// 对key进行升序排列
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("k:%v,v:%v\n", k, m[k])
	}
}

