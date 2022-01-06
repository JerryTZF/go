/**
 * Created by GoLand
 * Time: 2021/10/2 9:00 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: merge_two_lists.go
 * Desc:
 */
package main

// 21 合并两个有序链表
type listNode struct {
	Val  int
	Next *listNode
}

func MergeTwoLists(l1 *listNode, l2 *listNode) *listNode {
	guard := new(listNode)
	head := guard
	for {
		if l1 == nil {
			head.Next = l2
			break
		}
		if l2 == nil {
			head.Next = l1
			break
		}

		if l1.Val < l2.Val {
			head.Next = &listNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			head.Next = &listNode{Val: l2.Val}
			l2 = l2.Next
		}

		head.Next = head
	}

	return guard.Next
}
