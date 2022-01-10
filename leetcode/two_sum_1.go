/**
 * Created by GoLand
 * Time: 2022/1/10 10:41 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: two_sum_1.go
 * Desc: https://leetcode-cn.com/problems/two-sum/
 */
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3,2,4}, 6))
}

// 两数之和 nums = [2,7,11,15], target = 9
func twoSum(nums []int, target int) (r []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		if i+1 == l {
			break
		}

		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				r = append(r, []int{i, j}...)
				return
			}
		}
	}
	return
}
