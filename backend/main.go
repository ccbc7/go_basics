package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"time"
	"sync"
)


/*
* 逐次処理と並行処理の比較
*/
func main() {
	// traceパッケージを使ってプログラムのトレースを取得する
	f, err := os.Create("trace.out")

	// エラー処理
	if err != nil {
		log.Fatalln("Error:", err)
	}

	// deferでtraceファイルを閉じる
	defer func () {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()

	// Startメソッドでtraceを開始する
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	// deferでtraceを停止する
	defer trace.Stop()

	// contextパッケージを使ってキャンセル可能なコンテキストを作成する
	ctx, t := trace.NewTask(context.Background(), "main")

	// deferでタスクを終了する
	defer t.End()

	// runtimeパッケージを使ってCPUの数を取得する
	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

	// 逐次処理
	// task(ctx, "task1")
	// task(ctx, "task2")
	// task(ctx, "task3")

	// 並行処理
	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "task1")
	go cTask(ctx, &wg, "task2")
	go cTask(ctx, &wg, "task3")
	wg.Wait()

	/*
	* gorutineの立ち上げは高速で行われるが、起動には少し時間がかかる
	* 実行の順番は保証されない 321 とか 231 とか
	*/

	s := []int{4, 5, 6}
	for _, i := range s {
		wg.Add(1)
		// Goルーチンが実行される時点でiの値はループの最後の値（この場合は6）になっている可能性がある
		// なので引数を指定してあげることで,処理を逐次的に行うことができる
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()







	println("main function finish")
}

// 逐次処理用の関数
func task(ctx context.Context, name string) {
	// deferでタスクを開始し、.End()で終了する
	defer trace.StartRegion(ctx, name).End()
	// time.Sleepで1秒待つ
	time.Sleep(time.Second)
	// nameを出力する
	fmt.Println(name)
}

// 並行処理用の関数
func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	// deferでタスクを開始し、.End()で終了する
	defer trace.StartRegion(ctx, name).End()
	// wait groupのカウントが１減る
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}










/*
* 1. goroutineの完了を待つ
 */
// func main() {
// 	// sync.WaitGroupは、複数のGoルーチンの完了を待つための同期プリミティブ
// 	var wg sync.WaitGroup

// 	// Addメソッドで待つgoroutineの数を増やす
// 	wg.Add(1)

// 	go func() {
// 		// deferで関数が終了する直前にwg.Done()を呼び出す, wg.Done()でgoroutineのカウントが減る
// 		defer wg.Done()
// 		fmt.Println("goroutine invoked")
// 	}()
// 	// wg.Wait()を呼び出すことでgoroutineの完了を待つ
// 	wg.Wait()

// 	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
// 	fmt.Println("main fun finish")
// }
