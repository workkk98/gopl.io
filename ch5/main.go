package main

import "fmt"

func foo() (foo int, bar error) {
	defer func() {
		recover()

		foo = 1
	}()

	panic('1')
}

func main() {
	foo, bar := foo()

	if bar != nil {
		return
	}

	fmt.Printf("foo: %d\n", foo)
}
