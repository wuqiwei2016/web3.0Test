package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintOddNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
	}
}

func PrintEvenNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
	}
}

type Task func()

func runTask(tasks []Task) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	// 为每个任务启动一个协程
	for i, task := range tasks {
		go func(id int, t Task) {
			defer wg.Done()

			start := time.Now()
			t() // 执行任务
			duration := time.Since(start)

			fmt.Printf("任务%d 执行耗时: %v\n", id, duration)
		}(i+1, task) // i+1 让任务ID从1开始
	}
	wg.Wait() // 等待所有任务完成
}
