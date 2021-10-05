/**
 * Created by GoLand
 * Time: 2021/9/25 4:31 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: control_flow_.go
 * Desc:
 */
package pkg

import "fmt"

// goto 模拟循环
func GotoLoop(n int) {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i == n {
		return
	}
	goto HERE
}

// 普通switch
func SwitchDemo(n int) {
	switch n {
	case 10, 20, 30:
		fmt.Printf("n in [10,20,30] and n=%d\n", n)
		fallthrough
	case 40:
		fmt.Println("n=40")
	default:
		fmt.Println("n!=10,20,30,40")
	}
}

// switch 替代多个if-else结构
func SwitchReplaceIf(m int) {
	switch {
	case m > 0:
		fmt.Println("m>0")
	case m == 0:
		fmt.Println("m=0")
	case m < 0:
		fmt.Println("m<0")
	}
}

// 带有初始化的判断
func SwitchWithExpress(s [2]int) {
	switch a, b := s[0], s[1]; {
	case a < b:
		fmt.Println("s[1]>s[0]")
	case a == b:
		fmt.Println("s[1]==s[0]")
	case a > b:
		fmt.Println("s[1]<s[0]")
	}
}

// 普通for循环
func NormalForLoop(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

// 双指针for循环
func MultipleCounter() {
	for i, j := 10, 0; i >= 5 && j <= 5; i, j = i-1, j+1 {
		fmt.Printf("i=%d;j=%d\n", i, j)
	}
}

// 双指针实现切片反转
// 在函数调用时，像切片(slice)、字典(map)、接口(interface)、通道(channel),这样的引用类型都是默认使用引用传递(即使没有显式的指出指针)
func ReverseSlice(s []string) []string {
	for head, tail := 0, len(s)-1; head < tail; head, tail = head+1, tail-1 {
		s[head], s[tail] = s[tail], s[head]
	}
	return s
}

// 类似do-while循环
func DoWhile(n int) {
	for n > 10 {
		n--
		fmt.Println(n)
	}
}

// 类似while-do循环
func WhileDo() {
	var m int
	for ; ; m++ {
		if m > 10 {
			break
		}
	}
}

// 死循环
func ForeverLoop() {
	var i int
	for {
		i++
	}
}

// 遍历
func Foreach() {
	// 可迭代对象为字符串时
	for index, char := range "hello world" {
		fmt.Printf("index=%d;char=%c\n", index, char)
	}

	for index, item := range []string{"你好", "世界"} {
		fmt.Printf("index=%d;char=%s\n", index, item)
	}
}
