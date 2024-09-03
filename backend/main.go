package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
* read write mutexを使った場合,ロックをしていても、他のgoroutineが読み込みを行うことができる。
 */
func main() {
	var wg sync.WaitGroup
	var rwMu sync.RWMutex
	var c int

	wg.Add(4)
	go write(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	go read(&rwMu, &wg, &c)
	wg.Wait()
	fmt.Println("main end")

	// var wg sync.WaitGroup
	// var c int64

	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for j := 0; j < 10; j++ {
	// 			atomic.AddInt64(&c, 1)
	// 		}
	// 	}()
	// }

	// wg.Wait()
	// fmt.Println(c)
	// fmt.Println("main end")
}

// 通常のmutexでは、他のgoroutineがロックをしている間は待つが、read write mutexでは他のgoroutineが読み込みを行うことができる。
func read(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	defer mu.RUnlock() //読み込みロックを解除する

	time.Sleep(10 * time.Millisecond)
	mu.RLock()
	fmt.Println("read lock")
	fmt.Println(*c)
	time.Sleep(1 * time.Second)
	fmt.Println("read unlock")
}

/*
* ミューテックスロックを行い、cをインクリメントする
*/
func write(mu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	defer mu.Unlock() //書き込みロックを解除する

	mu.Lock()
	fmt.Println("write lock")
	*c++ // この時点でcは1になる
	time.Sleep(1 * time.Second)
	fmt.Println("write unlock")
}

/*
* ミューテックスを使ってみるが、i = 1が出力されることがあり、これは設計上の問題である。(go run -race main.go ででも検知できない)
 */
// func main() {
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	var i int
// 	wg.Add(2)
// 	go func() {
// 		// ミューテックスをロックしてからiをインクリメントする
// 		defer wg.Done()
// 		mu.Lock()
// 		defer mu.Unlock()
// 		// i++
// 		i = 1
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		defer mu.Unlock()
// 		// i++
// 		i = 2
// 	}()
// 	wg.Wait()
// 	fmt.Println(i)
// }
