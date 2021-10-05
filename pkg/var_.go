/**
 * Created by GoLand
 * Time: 2021/9/7 2:31 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: var_.go
 * Desc: 变量相关
 */
package pkg

// 全局变量声明
var (
	a = R
	b = "hello world"
)

// 如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明
func shortEqual() int {
	a := 10
	a = 15 // 这里不能写成 a := 15
	return a
}

// 当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil