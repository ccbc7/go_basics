package calculator

// 大文字で始まる関数や変数は他のパッケージからも参照できる

var  offset float64 = 1
var Offset float64 = 1

func Sum(a float64, b float64) float64 {
	return a + b + offset
}
