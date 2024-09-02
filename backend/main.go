package main

import(
	// "sync"
	"fmt"
	// "time"
	"runtime"
)

func main() {
	ch1 := make(chan int)
	// このチャネルには受信がないので、gotuitineリークが発生する
	go func () {
		fmt.Println(<-ch1)
	}()
	ch1 <- 10 //受信操作があるまでgorutineは処理されない
	fmt.Printf("num of working gorutines: %d\n", runtime.NumGoroutine())

	// バッファ付きチャネル,バッファとは、チャネルにデータを格納する領域のこと
	ch2 := make(chan int, 1)
	ch2 <- 2
	ch2 <- 3 // バッファが1なので、この時点でブロックされる
	fmt.Println(<-ch2)
}





/*
* 基礎的なチャネルを作成
*/
// func main() {
// 	// チャネルを作成
// 	ch := make(chan int)
// 	// syncで待ち合わせる
// 	var wg sync.WaitGroup
// 	// WaitGroupに1を追加
// 	wg.Add(1)

// 	go func() {
// 		// 関数が終了したらDoneを呼び出す
// 		defer wg.Done()
// 		// チャネルに10を書き込み
// 		ch <- 10
// 		// 0.5秒待つ
// 		time.Sleep(500 * time.Millisecond)
// 	}()

// 	// チャネルから受信
// 	fmt.Println(<-ch)

// 	// wg.Wait()がないと、goroutineが終了する前にmain関数が終了してしまうので設置
// 	wg.Wait()
// }
