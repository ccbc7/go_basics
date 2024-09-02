package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()

	ch1 <- 10
	close(ch1)

	v, ok := <-ch1 // okにはch1のクローズ状態が入る
	fmt.Printf("v: %v, ok: %v\n", v, ok)
	wg.Wait()

	// バッファ付きチャネルの場合, クローズがされても、読み込まれていないデータがあれば、okはtrueになる
	ch2 := make(chan int, 2)
	ch2 <- 20
	ch2 <- 30
	close(ch2)
	v, ok = <-ch2
	fmt.Printf("v: %v, ok: %v\n", v, ok) // 20, true
	v, ok = <-ch2
	fmt.Printf("v: %v, ok: %v\n", v, ok) // 30, true
	v, ok = <-ch2
	fmt.Printf("v: %v, ok: %v\n", v, ok) // 0, false

	ch3 := generateCountStream()
	for v := range ch3 {
		fmt.Println(v)
	}

	// データの値を持たない通知専用のチャネル
	nCH := make(chan struct{})
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %v start\n", i)
			<-nCH
			fmt.Println(i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(nCH)
	fmt.Println("finish---")


	wg.Wait()
	fmt.Println("finish")
}

// カプセル化されたチャネルを生成する
// 返り値として、読みとり専用のチャネルを返す
func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- i
		}
	}()
	return ch
}
