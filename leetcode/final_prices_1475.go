/**
 * Created by GoLand
 * Time: 2022/1/6 2:28 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: final_prices_1475.go
 * Desc: https://leetcode-cn.com/problems/final-prices-with-a-special-discount-in-a-shop/
 */
package main

import "fmt"

func main() {
	fmt.Println(finalPrices([]int{10, 1, 1, 6}))
}

func finalPrices(prices []int) (s []int) {
	l := len(prices)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if prices[i] >= prices[j] {
				s = append(s, prices[i]-prices[j])
				break
			}
		}
		if len(s) != i+1 {
			s = append(s, prices[i])
		}
	}
	return s
}
