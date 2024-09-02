package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	var wg sync.WaitGroup

	// タイムアウトを設定
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel() // タイムアウトした場合の後処理

	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		ch1 <- "A"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(800 * time.Millisecond)
		ch2 <- "B"
	}()

loop:
	for ch1 != nil || ch2 != nil {
		select {
		case <-ctx.Done(): // タイムアウトするとctx.Done()が返る
			fmt.Println("Timeout")
			break loop
		case v := <-ch1:
			fmt.Println(v)
			ch1 = nil
		case v := <-ch2:
			fmt.Println(v)
			ch2 = nil
		}
	}

	wg.Wait()
	fmt.Println("Done")
}
