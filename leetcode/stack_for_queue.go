/**
 * Created by GoLand
 * Time: 2021/10/12 3:26 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: stack_for_queue.go
 * Desc: 232、用栈实现队列
 */
package main

type MyQueue struct {
	inQueue, outQueue []int
}

// 初始化
func Constructor232() MyQueue {
	return MyQueue{}
}

// 将元素压入队尾
func (q *MyQueue) Push(x int) {
	q.inQueue = append(q.inQueue, x)
}

// 输入栈 压入 输出栈
func (q *MyQueue) transfer() {
	for len(q.inQueue) > 0 {
		q.outQueue = append(q.outQueue, q.inQueue[len(q.inQueue)-1])
		q.inQueue = q.inQueue[:len(q.inQueue)-1]
	}
}

// 从队列开头移除并返回元素
func (q *MyQueue) Pop() int {
	if len(q.outQueue) == 0 {
		q.transfer()
	}
	x := q.outQueue[len(q.outQueue)-1]
	q.outQueue = q.outQueue[:len(q.outQueue)-1]
	return x
}

// 返回队列开头的元素
func (q *MyQueue) Peek() int {
	if len(q.outQueue) == 0 {
		q.transfer()
	}
	return q.outQueue[len(q.outQueue)-1]
}

// 如果队列为空，返回 true ；否则，返回 false
func (q *MyQueue) Empty() bool {
	return len(q.outQueue) == 0 && len(q.inQueue) == 0
}
