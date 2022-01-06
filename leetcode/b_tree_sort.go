/**
 * Created by GoLand
 * Time: 2021/10/12 4:00 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: b_tree_sort.go
 * Desc: 二叉树的遍历
 */
package main

// 前序遍历：首先访问根结点然后遍历左子树，最后遍历右子树。在遍历左、右子树时，仍然先访问根结点，然后遍历左子树，最后遍历右子树
// 中序遍历：中序遍历首先遍历左子树，然后访问根结点，最后遍历右子树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144、二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	val := make([]int, 0)
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			val = append(val, root.Val)
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return val
}

// 94、二叉树中序遍历
func inorderTraversal(root *TreeNode) []int {
	val := make([]int, 0)
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		// 优先一路向左
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 左子节点已经没有元素,则向上移动;获取栈内最后一位
		root = stack[len(stack)-1]
		// 为了节点一步一步向上,保持站内最后一位为上一个节点
		stack = stack[:len(stack)-1]
		// 记录当前节点的值
		val = append(val, root.Val)
		// 节点转换为当前根节点的右节点(左根右的右)
		root = root.Right
	}
	return val
}
