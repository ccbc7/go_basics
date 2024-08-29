package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ui1 uint16 // 16bitの符号なし整数型 この時点で初期値は0
	fmt.Printf("memory address of ui1: %p\n", &ui1)

	var ui2 uint16
	fmt.Printf("memory address of ui2: %p\n", &ui2)

	var p1 *uint16 // ポインタ変数
	fmt.Printf("value of p1: %v\n", p1) // ポインタ変数の初期値はnil

	/*
	* ポインタ変数の容量は必ず8byte
	* なぜなら、64bitのアーキテクチャでは、メモリアドレスの最大値が64bitだから
	* &はアドレス演算子, &をつけることで、変数のメモリアドレスを取得できる
	* *はポインタ演算子, *をつけることで、ポインタ変数の値を参照できる
	*/
	p1 = &ui1 // ui1のメモリアドレスをp1に代入
	fmt.Printf("value of p1: %v\n", p1) // p1のメモリアドレス
	fmt.Printf("size of p1: %d\n", unsafe.Sizeof(p1)) // ポインタ変数の容量
	fmt.Printf("memory address of p1: %p\n", &p1) // p1のメモリアドレス
	fmt.Printf(("value of uil(dereference): %v\n"), *p1) // ポインタ変数の値を参照 この時点では0

	*p1 = 1
	fmt.Printf(("value of uil(dereference): %v\n"), *p1) // ポインタ変数の値を参照 この時点では1

	var pp1 **uint16 = &p1 // *p1のメモリアドレスをpp1に代入

	fmt.Printf("value of pp1: %v\n", pp1) // pp1はポインタなので、p1のメモリアドレスが表示される
	fmt.Printf("memory address of pp1: %p\n", &pp1) // pp1のメモリアドレス

	fmt.Printf("value of pp1(dereference): %v\n", *pp1) // p1のメモリアドレスが表示される
	fmt.Printf("value of pp1(dereference)(dereference): %v\n", **pp1) //結果的にはui1の値が表示される

	/*
	* shadowing
	* if文の中で宣言した変数名と、if文の外で宣言した変数名が同じ場合、その変数は別のメモリ領域に割り当てられる
	* そのため、if文の中で宣言した変数はif文の外で宣言した変数に影響を与えない
	*/
	ok, result := true, "A"
	if ok {
		result := "B"
		fmt.Println(result) // B
	} else {
		result := "C"
		fmt.Println(result) // C
	}
	fmt.Println(result) // A

	/*
	* 同じアドレスと出力を出したい場合は、:=ではなく、=を使う
	*/
	ok, result = true, "A"
	if ok {
		result = "B"
		fmt.Println(result) // B
	} else {
		result = "C"
		fmt.Println(result) // C
	}
	fmt.Println(result) // B
}
