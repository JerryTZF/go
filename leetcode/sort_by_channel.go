/**
 * Created by GoLand
 * Time: 2021/10/25 3:47 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: sort_by_channel.go
 * Desc: 1114、按序打印
 */
package leetcode

import (
	"fmt"
	"time"
)

var ch1 = make(chan int)
var ch2 = make(chan int)

func first() {
	fmt.Println("1")
	ch1 <- 1
}
func second() {
	<-ch1
	fmt.Println("2")
	ch2 <- 1
}
func third() {
	<-ch2
	fmt.Println("3")
}
func main() {
	go first()
	go second()
	go third()
	time.Sleep(time.Second)
}
