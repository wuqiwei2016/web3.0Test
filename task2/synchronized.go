package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//func main() {
//
//	var (
//		counter int
//		wg      sync.WaitGroup
//		mu      sync.Mutex // 互斥锁，用于保护计数器
//	)
//
//	// 启动10个协程
//	wg.Add(10)
//	for i := 0; i < 10; i++ {
//		go func() {
//			defer wg.Done()
//			// 每个协程执行1000次递增操作
//			for j := 0; j < 1000; j++ {
//				mu.Lock()   // 获取锁
//				counter++   // 安全地递增计数器
//				mu.Unlock() // 释放锁
//			}
//		}()
//	}
//
//	wg.Wait() // 等待所有协程完成
//	fmt.Printf("最终计数器值: %d\n", counter)
//}

func main() {
	var (
		counter int64 // 必须使用int64保证原子操作
		wg      sync.WaitGroup
	)

	// 启动10个协程
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 每个协程执行1000次原子递增操作
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1) // 原子递增
			}
		}()
	}

	wg.Wait() // 等待所有协程完成
	fmt.Printf("最终计数器值: %d\n", counter)
}
