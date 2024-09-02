package main

import (
	"fmt"
	"sync"
	"time"
)

const bufSize = 3

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, bufSize)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < bufSize; i++ {
			time.Sleep(1000 * time.Millisecond)
			ch <- "hello"
		}
	}()

	// gorutineとは独立して、こちらのfor文が実行される
	for i := 0; i < 3; i++ {
		select {
		case m := <-ch: // チャネルからメッセージを受信した場合
			fmt.Println(m) // 受信したメッセージを表示
		default: // メッセージが受信されなかった場合
			fmt.Println("no message received") // メッセージが受信されなかったことを表示
		}
		time.Sleep(1500 * time.Millisecond) // 1.5秒間スリープ
	}

	wg.Wait() // ゴルーチンの完了を待つ
	fmt.Println("Done")
}
