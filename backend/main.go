/*
* セレクト文を使って複数のチャネルの値を連続して受け取る方法
 */

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const bufSize = 5

func main() {
	ch1 := make(chan int, bufSize)
	ch2 := make(chan int, bufSize)
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Millisecond)
	defer cancel()

	wg.Add(3)
	go countProducer(&wg, ch1, bufSize, 50)
	go countProducer(&wg, ch2, bufSize, 500)
	go countConsumer(ctx, &wg, ch1, ch2)
	wg.Wait()
	fmt.Println("main: done")
}

/*
* 第1引数: 変数をポインタ型で受け取る
* 第2引数: チャネル書き込み用で受け取る
* 第3引数: バッファサイズ
* 第4引数: 遅延時間
 */
func countProducer(wg *sync.WaitGroup, ch chan<- int, size int, sleep int) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < size; i++ {
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		ch <- i
	}
}

/*
* 第1引数: コンテキスト
* 第2引数: 変数をポインタ型で受け取る
* 第3引数: チャネル読み込み用で受け取る
* 第4引数: チャネル読み込み用で受け取る
 */
func countConsumer(ctx context.Context, wg *sync.WaitGroup, ch1 <-chan int, ch2 <-chan int) {
	defer wg.Done()
loop:
	for ch1 != nil || ch2 != nil {
		// 1つのselect文で複数のチャネルを受信することができる
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break loop
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Printf("ch1: %v\n", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Printf("ch2: %v\n", v)
		}
	}
}
