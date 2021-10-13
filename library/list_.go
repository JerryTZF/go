/**
 * Created by GoLand
 * Time: 2021/10/13 9:51 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: list_.go
 * Desc: 链表
 */
package library

import (
	"container/list"
	"fmt"
)

func ListDemo() {
	l1 := list.New()

	// 链表末尾插入
	e1 := l1.PushBack("99")
	// 链表首位插入
	e2 := l1.PushFront(1)
	// 在指定元素后插入
	l1.InsertAfter(2, e2)
	// 在指定元素前插入
	l1.InsertBefore("98", e1)

	// 正序遍历链表
	for i := l1.Front(); i != nil; i = i.Next() {
		fmt.Println("L1's Element = ", i.Value)
	}

	l2 := list.New()
	// 在头部插入另一个队列
	l2.PushFrontList(l1)
	// 在队尾插入另一个队列
	l2.PushBackList(l1)

	// 逆序遍历链表
	for i := l2.Back(); i != nil; i = i.Prev() {
		fmt.Println("L2's Element = ", i.Value)
	}

	// 删除元素
	e3 := l1.Remove(e2)
	fmt.Println(e3)

	// 将元素1移到元素2前面
	l3 := list.New()
	el1 := l3.PushBack(1)
	el2 := l3.PushBack(2)
	l3.MoveBefore(el2, el1)

	// 将元素移动到头部
	l3.MoveToBack(el2)
	// 将元素移动到尾部
	l3.MoveToBack(el1)
	// 获取列表长度
	l3.Len()
}
