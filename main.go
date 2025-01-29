package main

import (
	"crypto/sha256"
	"fmt"

	"gopl.io/ch2/popcount"
)

func arrayTest() {
	// a := [...]int{0, 1, 2}
	c1 := sha256.Sum256([]byte("x"))
	// fmt.Println("Hello, World!", popcount.PopCount(55), a[1])
	// fmt.Println(c1, len(c1))

	foo := [sha256.Size]int{}

	bar := make([]int, sha256.Size)

	bar[0] = 1
	fmt.Println("bar len", len(bar), cap(bar))

	for i, val := range c1 {
		fmt.Println("每一位的数值", val)
		foo[i] = popcount.PopCount(uint64(val))
	}

	// fmt.Println("sha256计算出来 [32]byte的位数", foo)
}

func main() {
	arrayTest()
}
