package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	// context.WithTimeoutで親のコンテキストを作成し、子のコンテキストを作成する
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	// errgroupを持つ子のコンテキストを作成
	eg, ctx := errgroup.WithContext(ctx) // 子goroutineでエラーが起きた場合、親のerrgroupにエラーが返ってくる

	s := []string{"task1", "task2", "task3", "task4"}
	// s := []string{"task1", "fake1", "task2", "fake2", "task3"}

	for _, v := range s {
		task := v
		eg.Go(func() error {
			return doTask(ctx, task)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("finished")
}

func doTask(ctx context.Context, task string) error {
	var t *time.Ticker

	// 引数taskを受け取り,caseで分岐
	switch task {
	case "task1":
		t = time.NewTicker(500 * time.Millisecond) // 500ms遅延させる
	case "task2":
		t = time.NewTicker(700 * time.Millisecond) // 700ms遅延させる
	default:
		t = time.NewTicker(1000 * time.Millisecond) // 200ms遅延させる
	}

	select {
	case <-ctx.Done():
		fmt.Printf("%v cancelled: %v\n", task, ctx.Err())
		return ctx.Err()
	case <-t.C:
		t.Stop()

		// if task == "fake1" || task == "fake2" {
		// 	return fmt.Errorf("%v failed", task) // 親のerrgroupにエラーを返す
		// }
		fmt.Printf("%v done\n", task)
	}

	return nil
}
