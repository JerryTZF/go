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

// TODO 只读管道、只写管道
