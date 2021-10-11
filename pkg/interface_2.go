/**
 * Created by GoLand
 * Time: 2021/10/11 3:34 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: interface_2.go
 * Desc:
 */
package pkg

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

func CountInto(a Appender, s, e int) {
	for i := s; i < e; i++ {
		a.Append(i)
	}
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func Main() {
	var lis List
	if LongEnough(lis) {
		fmt.Printf("- lst is long enough\n")
	}

	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
}
