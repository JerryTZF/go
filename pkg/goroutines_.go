/**
 * Created by GoLand
 * Time: 2021/10/19 5:42 下午
 * Author: JerryTian<tzfforyou@163.com>
 * File: goroutines_.go
 * Desc:
 */
package pkg

// 使用管道实现一个互斥锁
type Empty interface{}
type semaphore chan Empty

// 写入管道
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
