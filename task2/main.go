package main

import (
	"fmt"
	"sync"
	"time"
)

func addNum(num *int) {
	*num += 10
}

func main() {
	// 指针操作
	val := 10
	fmt.Println("原始的值", val)
	addNum(&val)
	fmt.Println("修改的值", val)

	// Goroutine使用
	var wg sync.WaitGroup
	wg.Add(2)
	go PrintOddNumbers(&wg)
	go PrintEvenNumbers(&wg)
	wg.Wait()

	// 定义一组示例任务
	tasks := []Task{
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("任务3完成")
		},
	}

	fmt.Println("开始执行任务...")
	runTask(tasks) // 调用在goroutine.go中定义的函数
	fmt.Println("所有任务执行完毕!")
}
