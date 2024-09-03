package main

import (
	"context"
	"fmt"

	// "sync"
	"time"
)

func main() {
	// 40ミリ秒以内にデッドラインが切れるように設定
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(40*time.Millisecond))
	defer cancel()

	// サブタスクを実行し、結果を受け取る
	ch := subTask(ctx)

	v, ok := <-ch
	if ok {
		fmt.Println(v)
	}
	fmt.Println("finish")
}

func subTask(ctx context.Context) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		deadline, ok := ctx.Deadline()
		if ok {
			// もし30ミリ秒以内にデッドラインが切れる場合は、処理を中断する
			if deadline.Sub(time.Now().Add(30*time.Millisecond)) < 0 {
				fmt.Println("impossible meet to deadline")
				return
			}
		}
		time.Sleep(30 * time.Millisecond)
		ch <- "hello"
	}()
	return ch
}

// func main() {
// 	var wg sync.WaitGroup
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	wg.Add(1)

// 	// criticalTask
// 	go func() {
// 		defer wg.Done()
// 		v, err := criticalTask(ctx)
// 		if err != nil {
// 			fmt.Printf("criticalTask canceled due to %v\n", err)
// 			cancel() // ここで親のcontextをキャンセルする
// 			return
// 		}
// 		fmt.Println("success", v)
// 	}()

// 	wg.Add(1)

// 	// normalTask
// 	go func() {
// 		defer wg.Done()
// 		v, err := normalTask(ctx)
// 		if err != nil {
// 			fmt.Printf("normalTask canceled due to %v\n", err)
// 			return
// 		}
// 		fmt.Println("success", v)
// 	}()
// 	wg.Wait()
// }

// // 引数にcontextを受け取ることで、親の
// func criticalTask(ctx context.Context) (string, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 1200*time.Millisecond)
// 	defer cancel() // このcancelはあくまでこの関数内でのみ有効なもので、親のcontextには影響を与えない
// 	t := time.NewTicker(1000 * time.Millisecond)
// 	select {
// 	case <-ctx.Done(): // ctx.Done()が起きるということは、親のcontextでタイムアウトが発生したとき
// 		return "", ctx.Err()
// 	case <-t.C: // t.Cが起きるということは、NewTickerで指定した時間感覚が経過したとき
// 		t.Stop()
// 	}
// 	return "criticalTask", nil
// }

// func normalTask(ctx context.Context) (string, error) {
// 	t := time.NewTicker(3000 * time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-t.C:
// 		t.Stop()
// 	}
// 	return "normalTask", nil
// }

/*
* context.WithTimeout()で指定した時間が経過すると、親のcontextでタイムアウトが発生する
 */
// func main() {
// 	var wg sync.WaitGroup
// 	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
// 	defer cancel()

// 	wg.Add(3)
// 	go subTask(ctx, &wg, "a")
// 	go subTask(ctx, &wg, "b")
// 	go subTask(ctx, &wg, "c")
// 	wg.Wait()
// }

// func subTask(ctx context.Context, wg *sync.WaitGroup, id string) {
// 	defer wg.Done()
// 	t := time.NewTicker(500 * time.Millisecond) //　指定した時間感覚でチャネルへ書き込み信号を送る

// 	select {
// 	// ctx.Done()が起きるということは、親のcontextでタイムアウトが発生したとき
// 	case <-ctx.Done():
// 		fmt.Println(ctx.Err())
// 		return
// 	// t.Cが起きるということは、NewTickerで指定した時間感覚が経過したとき
// 	case <-t.C: // Ticker構造体のCフィールドには、チャネルが格納されているのでここに保存されている時間を取り出す
// 		t.Stop()
// 		fmt.Println(id)
// 	}
// }
