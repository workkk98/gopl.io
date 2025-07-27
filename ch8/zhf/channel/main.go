package main

import "fmt"

type Foo struct {
	A int
	B int
}

func producer(ch chan Foo) {
	for i := 0; i < 6; i++ {
		// Foo字面量
		ch <- Foo{
			A: i,
			B: 2 * i,
		}
	}

	close(ch)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	ch := make(chan Foo) // 无缓冲通道

	go producer(ch)

	for {
		x, ok := <-ch

		if !ok {
			break
		}
		fmt.Printf("\r x^2= %d\n", x.A+x.B)
	}

}
