package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func testContextWithTimeout() {
	// 创建一个带有超时时间的上下文对象
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// 启动一个 goroutine，模拟一个耗时的操作
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// 在超时或取消时结束 goroutine 的执行
				fmt.Println("goroutine canceled")
				return
			default:
				// 模拟一个耗时的操作
				time.Sleep(100 * time.Millisecond)
				fmt.Println("working...")
			}
		}
	}(timeoutCtx)

	// 等待一段时间，然后取消上下文对象
	time.Sleep(time.Second * 2)
	cancel()

	// 等待一段时间，以便观察输出结果
	time.Sleep(500 * time.Second)
}

func testContextWithValue() {
	// 定义一个全局的 context.Context 对象
	var ctx = context.Background()
	// 创建 10 个 goroutine，每个 goroutine 都会设置一个值
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 在 context 中设置一个值
			ctx = context.WithValue(ctx, "id", fmt.Sprintf("goroutine-%d", i))
			// 执行一些操作，可以访问 ctx 中的值
			// 从 context 中获取值
			id := ctx.Value("id").(string)
			fmt.Println(id)
		}(i)
	}
	wg.Wait()
}

func main() {
	// testContextWithTimeout()
	testContextWithValue()
}
