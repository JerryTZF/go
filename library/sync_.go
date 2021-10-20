/**
 * Created by GoLand
 * Time: 2021/10/20 11:48 上午
 * Author: JerryTian<tzfforyou@163.com>
 * File: sync_.go
 * Desc: sync包的学习和整理
 */
package library

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
)

// 互斥锁的使用
type BankA struct {
	sync.Mutex // 可以理解为实现了上锁和解锁的一个官方结构体
	balance    map[string]float64
}

func (b *BankA) In(account string, value float64) {
	// 上锁
	b.Lock()
	// 析构解锁
	defer b.Unlock()

	if _, ok := b.balance[account]; !ok {
		b.balance[account] = 0.00
	}

	b.balance[account] += value
}

func (b *BankA) Out(account string, value float64) error {
	// 上锁
	b.Lock()
	// 析构解锁
	defer b.Unlock()

	v, ok := b.balance[account]
	if !ok || v < value {
		return fmt.Errorf("%s not enough balance or is null", account)
	}
	b.balance[account] -= value
	return nil
}

// 文件读写锁
// 当有一个协程在读的时候，所有写的协程必须等到所有读的协程结束才可以获得锁进行写操作。
// 当有一个协程在读的时候，所有读的协程不受影响都可以进行读操作。
// 当有一个协程在写的时候，所有读、写的协程必须等到写的协程结束才可以获得锁进行读、写操作。

type BankB struct {
	sync.RWMutex
	balance map[string]float64
}

func (b *BankB) In(account string, value float64) {
	b.Lock()
	defer b.Unlock()
	if _, ok := b.balance[account]; !ok {
		b.balance[account] = 0.00
	}

	b.balance[account] += value
}

func (b *BankB) Out(account string, value float64) error {
	// 上锁
	b.Lock()
	// 析构解锁
	defer b.Unlock()

	v, ok := b.balance[account]
	if !ok || v < value {
		return fmt.Errorf("%s not enough balance or is null", account)
	}
	b.balance[account] -= value
	return nil
}

func (b *BankB) Query(account string) float64 {
	b.RLock()
	defer b.RUnlock()
	if _, ok := b.balance[account]; !ok {
		return 0.0
	}

	return b.balance[account]
}

// waitGroup
// 阻塞携程组全部完成
func WaitGroup() {
	wg := new(sync.WaitGroup)
	t1 := time.Now().UnixNano() / 1e6
	// 创建1000个协程
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(no int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println(no)
		}(i)
	}

	wg.Wait()
	fmt.Printf("耗时: %d 毫秒", time.Now().UnixNano()/1e6-t1)
}

// 最多可执行10个协程
func WaitGroup_() {
	wg := New(10)
	t1 := time.Now().UnixNano() / 1e6
	for i := 0; i < 100; i++ {
		wg.Add()
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println(n)
		}(i)
	}

	wg.Wait()
	fmt.Printf("耗时: %d 毫秒", time.Now().UnixNano()/1e6-t1)
}

// 保证只执行一次
func SyncOnce() {
	once := new(sync.Once)
	for i := 0; i < 1000; i++ {
		go func(n int) {
			once.Do(func() {
				time.Sleep(time.Second)
				fmt.Println(n)
			})
		}(i)
	}

	time.Sleep(5 * time.Second)
}

// TODO 临时对象池暂时不用关注 Orz

// 线程安全的map相关操作
func SafeMap() {
	m := sync.Map{}

	// 写
	for i := 0; i < 5; i++ {
		go func(n int) {
			m.Store("sync", n)
		}(i)
	}

	// 读
	if v, ok := m.Load("sync"); ok {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}

	// 存在指定的 key 则读取，否则写入
	if v, ok := m.LoadOrStore("SYNC", "faker"); ok {
		fmt.Println("这是写入操作，值为：", v)
	} else {
		fmt.Println("这是读取操作，值为：", v)
	}

	// 删
	m.Delete("SYNC")

	// 遍历操作
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("key:%v;value:%v\n", key, value)
		//if key == "sync" {
		//	return true
		//}
		//return false
		return true
	})
}

// 封装一个`可执行最大协程数`的包
// 详情参见：https://github.com/remeh/sizedwaitgroup
type SizedWaitGroup struct {
	Size    int
	current chan struct{}
	wg      sync.WaitGroup
}

func New(limit int) SizedWaitGroup {
	size := math.MaxInt32
	if limit > 0 {
		size = limit
	}
	return SizedWaitGroup{Size: size, current: make(chan struct{}, size), wg: sync.WaitGroup{}}
}

func (s *SizedWaitGroup) Add() {
	_ = s.AddWithContext(context.Background())
}

func (s *SizedWaitGroup) AddWithContext(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s.current <- struct{}{}:
		break
	}
	s.wg.Add(1)
	return nil
}

func (s *SizedWaitGroup) Done() {
	<-s.current
	s.wg.Done()
}

func (s *SizedWaitGroup) Wait() {
	s.wg.Wait()
}
