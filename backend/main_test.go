package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	// main関数に対してテストを実行
	main()
}
