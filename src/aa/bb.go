package main

import (
	"fmt"
	"time"
	"context"
	"math/rand"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	eatNum := chiHanBao(ctx) //主程序持续监听子goroutine的channel信使的通风报信
	for n := range eatNum {
		fmt.Println("n = ", n)
		if n >= 10 {
			cancel() //达到主程序的底线后，主程序发号施令，cancel()子goroutine
			break
		}
	}

	fmt.Println("calculating ...")
	time.Sleep(1 * time.Second)
}

func chiHanBao(ctx context.Context) <-chan int {
	c := make(chan int)
	// 个数
	n := 0
	// 时间
	t := 0
	go func() {
		for {
			//time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Printf("take %d seconds, eat %d hamburgers \n", t, n)
				return
			case c <- n:
				incr := rand.Intn(5)
				n += incr
				if n >= 10 {
					n = 10
				}
				t++
				fmt.Printf("I eat %d hamburgers\n", n)
			}
		}
	}()

	return c
}