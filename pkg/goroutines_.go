/**
 * Created by GoLand
 * Time: 2021/10/19 5:42 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: goroutines_.go
 * Desc:
 */
package pkg

import (
	"fmt"
	"sort"
	"time"
)

// 并发
// 指在同一时刻只能有一条指令执行，但多个进程指令被快速地轮换执行，使得在宏观上具有多个进程同时执行的效果
// 并行
// 并行是指在同一时刻，有多条指令在多个处理器上同时执行

// 协程相关
func GoroutineDemo() {
	c := make(chan string)
	go func() {
		c <- "No.1"
	}()

	fmt.Println(<-c)
	close(c)
}

func GoChan() {
	// 协程间通信
	c1 := make(chan int)
	// 并发接收数据
	go func(c chan int) {
		for {
			data := <-c
			if data == -1 {
				fmt.Println("Receive Over")
				break
			}
			fmt.Println("Receive Data =", data)
		}
		c <- -1
	}(c1)

	// 投递数据
	for i := 1; i <= 3; i++ {
		// 将数据通过channel投送给numPrint
		c1 <- i
	}

	time.Sleep(5 * time.Second)

	// 通知 numPrint 结束循环
	fmt.Println("Notify send over")
	c1 <- -1

	<-c1
}

// 只读管道、只写管道
func producer(sender chan<- int) {
	for i := 0; i < 3; i++ {
		sender <- i
		// 休眠时会阻塞管道读取数据
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("Send", i)
	}
	// <-sender (会报错)
	close(sender)
}

func consumer(receive <-chan int) {
	for n := range receive {
		fmt.Println("Receive =", n)
		// 休眠时会阻塞管道写入数据
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func MQDemo() {
	// 创建无缓冲管道(协程间通信是同步的)
	//ch := make(chan int)
	// 创建有缓冲管道(协程间通信是异步的)
	ch_ := make(chan int, 3)
	go producer(ch_)
	consumer(ch_)
}

// 管道超时处理
func ChanTimeout() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("received num = ", num)
			case <-time.After(time.Second * 3):
				fmt.Println("TimeOut")
				quit <- true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(4 * time.Second)
	}
	<-quit
	fmt.Println("Over")
}

// 例子：将一个长切片分为N块，交由N个协程完成排序，然后通过管道返回，然后合并为一个排好序的切片
func SortByGoroutines() {
	s := [][]int{{1, 3, 2, 5, 4, 6, 8, 7}, {9, 13, 10, 12, 11, 16, 15}, {21, 24, 23, 26, 25, 28, 27}}
	ch := make(chan []int)
	var rs []int
	for _, v := range s {
		go func(sk []int) {
			sort.Sort(sort.IntSlice(sk))
			ch <- sk
		}(v)
	}
	for i := 0; i < len(s); i++ {
		rs = append(rs, <-ch...)
	}
	fmt.Println(rs)
}
