package main

import (
	"fmt"
	"sync"
)

//	func main() {
//		// 创建一个无缓冲通道
//		ch := make(chan int)
//
//		// 使用WaitGroup等待协程完成
//		var wg sync.WaitGroup
//		wg.Add(2)
//
//		// 生产者协程：生成1到10的整数并发送到通道
//		go func() {
//			defer wg.Done()
//			for i := 1; i <= 10; i++ {
//				ch <- i
//			}
//			close(ch) // 发送完成后关闭通道
//		}()
//
//		// 消费者协程：从通道接收并打印整数
//		go func() {
//			defer wg.Done()
//			for num := range ch {
//				fmt.Printf("接收到: %d\n", num)
//			}
//		}()
//
//		// 等待两个协程都完成
//		wg.Wait()
//		fmt.Println("程序结束")
//	}
func main() {
	// 创建一个缓冲大小为10的通道
	ch := make(chan int, 10) // 缓冲通道，可以存储10个元素

	// 使用WaitGroup等待协程完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 生产者协程：生成1到100的整数并发送到通道
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Printf("发送: %d\n", i) // 打印发送的数字
		}
		close(ch) // 发送完成后关闭通道
	}()

	// 消费者协程：从通道接收并打印整数
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("接收到: %d\n", num)
		}
	}()

	// 等待两个协程都完成
	wg.Wait()
	fmt.Println("程序结束")
}
