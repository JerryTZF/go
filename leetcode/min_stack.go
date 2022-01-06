/**
 * Created by GoLand
 * Time: 2021/10/12 4:40 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: min_stack.go
 * Desc: 155、最小栈
 */
package main

import "math"

type MinStack struct {
	myStack, min []int
}

// 初始化
func Constructor155() MinStack {
	return MinStack{myStack: []int{}, min: []int{math.MaxInt64}}
}

//  将元素 x 推入栈中
func (m *MinStack) Push(i int) {
	m.myStack = append(m.myStack, i)
	m.min = append(m.min, min(m.min[len(m.min)-1], i))
}

// 删除栈顶的元素
func (m *MinStack) Pop() {
	m.myStack = m.myStack[:len(m.myStack)-1]
	m.min = m.min[:len(m.min)-1]
}

// 获取栈顶元素
func (m *MinStack) Top() int {
	return m.myStack[len(m.myStack)-1]
}

// 获取最小元素
func (m *MinStack) GetMin() int {
	return m.min[len(m.min)-1]
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
