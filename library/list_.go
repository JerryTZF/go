/**
 * Created by GoLand
 * Time: 2021/10/13 9:51 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: list_.go
 * Desc: 链表
 */
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l1 := list.New()

	// 链表末尾插入
	e1 := l1.PushBack("tail")
	// 链表头部插入
	e2 := l1.PushFront("head")
	// 在指定元素后插入
	l1.InsertAfter("realTail", e1)
	// 在指定元素前面插入
	l1.InsertBefore("realHead", e2)

	listSort(l1, 1)

	l2 := list.New()

	// 在头部插入另一个队列
	l2.PushFrontList(l1)
	// 在队尾插入另一个队列
	l2.PushBackList(l1)

	listSort(l2, 1)

	// 删除元素
	e3 := l1.Remove(e2)
	fmt.Println(e3)

	l3 := list.New()

	el1 := l3.PushBack(1)
	el2 := l3.PushBack(2)

	// 将元素1移到元素2前面
	l3.MoveBefore(el2, el1)

	// 将元素移动到头部
	l3.MoveToFront(el2)
	// 将元素移动到尾部
	l3.MoveToBack(el1)
	// 获取列表长度
	l3.Len()
}

// 逆序遍历链表
func listSort(l *list.List, f int) {
	if f > 0 {
		for i := l.Front(); i != nil; i = i.Next() {
			fmt.Println("l's Element=", i.Value)
		}
	} else {
		for i := l.Back(); i != nil; i = i.Prev() {
			fmt.Println("l's Element=", i.Value)
		}
	}
	fmt.Println("==============")
}
