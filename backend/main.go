package main

import "fmt"

type Os int // int型を基にした独自の型Osを定義

const (
	Mac Os = iota + 1 //iotaは連番を生成する
	Windows
	Linux
)

func main() {
	var i int
	fmt.Println(i) // 0

	var f int = 2
	fmt.Println(f) // 2

	g := 3         // 型を省略
	fmt.Println(g) // 3

	h := uint16(2)                                       // 型を指定
	fmt.Printf("h %v %T\n", h, h)                        //%vは値を、%Tは型を表示, 第一引数が%vに、第二引数が%Tに対応
	fmt.Printf("g: %[1]v %[1]T, h: %[2]v %[2]T\n", g, h) //%[1]は第一引数、%[2]は第二引数

	j := 1.23456
	k := "hello"
	l := true
	fmt.Printf("j: %[1]v %[1]T, k: %[2]v %[2]T, l: %[3]v %[3]T\n", j, k, l)

	pi, title := 3.14159, "円周率"
	fmt.Printf("pi: %T, title: %T\n", pi, title) // pi: float64, title: string

	x := 1
	y := 2.22
	z := float64(x) + y // xをfloat64に型変換
	fmt.Println(z)      // 3.22

	// 定数を出力
	fmt.Printf("Mac: %v, Windows: %v, Linux: %v\n", Mac, Windows, Linux) // Mac: 1, Windows: 2, Linux: 3

	i = 2 //変数iに再代入
	fmt.Printf("i: %v\n", i) // i: 2

	i += 1
	fmt.Printf("i: %v\n", i) // i: 3
}
