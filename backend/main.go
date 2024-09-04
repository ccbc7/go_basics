package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	cores := runtime.NumCPU()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nums := []int {1, 2, 3, 4, 5, 6, 7, 8}

	// ouyChsには各goroutineの出力チャネルが格納される
	outChs := make([]<-chan string, cores)


	inData := generator(ctx, nums...)
	for i := 0; i < cores; i++ {
		outChs[i] = fanOut(ctx, inData, i+1)
	}

	var i int
	flag := true
	for v := range fanIn(ctx, outChs...) {
		if i == 3 {
			cancel()
			flag = false
		}
		if flag {
			fmt.Println(v)
		}
		i++
	}
	fmt.Println("finish")
}

func generator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()
	return out
}

/*
* fanOut: 1つのチャネルから複数のチャネルにデータを分配する
* 引数1: コンテクスト
* 引数2: 入力チャネル
* 引数3: ID
* 戻り値: 出力チャネル
*/
func fanOut(ctx context.Context, in <-chan int, id int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		// 重い処理をシミュレート
		heavyWork := func(i int, id int) string {
			time.Sleep(200 * time.Millisecond)
			return fmt.Sprintf("result:%v (id:%v)", i*i, id)
		}
		for v := range in {
			select {
			case <-ctx.Done():
				return
			// 重い処理を行い、結果を出力チャネルに送信
			case out <- heavyWork(v, id):
			}
		}
	}()
	return out
}

/*
* fanIn: 複数のチャネルからデータを受け取り、1つのチャネルに集約する
*/
func fanIn(ctx context.Context, chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	multiplex := func(ch <-chan string) {
		defer wg.Done()
		for text := range ch {
			select {
			case <-ctx.Done():
				return
			case out <- text:
			}
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go multiplex(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
