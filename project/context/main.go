package main

import (
	"context"
	"fmt"
	"time"
)

var stream chan string

func Consumer(ctx context.Context, stream string) error {
	fmt.Println(stream)
	return nil
}

func Product(ctx context.Context) error {
	value := ctx.Value("key")
	if v, ok := value.(string); ok {
		stream <- v
	}
	return nil
}

func run(ctx context.Context) {
	for {
		select {
		case val := <-stream:
			Consumer(ctx, val)
		}
	}
}
func sleep2Sec(ctx context.Context) {
	defer func(t time.Time) {
		fmt.Println("sleep2Sec run", time.Now().Sub(t))
	}(time.Now())
	t := time.NewTimer(time.Second * 10)
	select {
	case <-ctx.Done(): // 不监听 取消后对程序没有影响，会继续运行
		fmt.Println("ctx canceled")
	case <-t.C:
		fmt.Println("timer run")
	}
}
func ctxCancel() {
	ctx := context.Background()
	ctx1 := context.WithValue(ctx, 1, 2)
	ctx2, cancelFunc := context.WithCancel(ctx1)
	ctx3 := context.WithValue(ctx2, 2, 3)
	ctx4, cancelFunc2 := context.WithCancel(ctx3)
	defer cancelFunc()
	defer cancelFunc2()
	fmt.Println(<-ctx4.Done())
	fmt.Println(<-ctx3.Done())
	fmt.Println(<-ctx2.Done())

}

func main() {
	// ctx := context.Background()
	// stream = make(chan string, 10)
	// go run(ctx)
	// for i := 0; i < 100; i++ {
	// 	ctx1 := context.WithValue(ctx, "key", strconv.Itoa(i))
	// 	Product(ctx1)
	// }
	// ctx2, cancelFunc := context.WithTimeout(ctx, time.Second*2)
	// ctx3, cancelFunc := context.WithTimeout(ctx2, time.Second*10) // 设置 ctx3 10 秒后取消，ctx2 2s后取消，其所有孩子context 都会取消，估ctx3也会两秒后取消
	// defer cancelFunc()
	// go sleep2Sec(ctx3)
	// time.Sleep((10 * time.Second))
	ctxCancel()

}
