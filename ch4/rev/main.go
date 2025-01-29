// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	// 原地修改
	// data := []string{"1", "", "hello"}
	// fmt.Println("noneempty", noneempty(data))
	// fmt.Println("data", data)

	// 原地修改 去重
	data2 := []string{"1", "hello", "hello", "2"}
	fmt.Println("noneduplicate", noneduplicate(data2))
	fmt.Println("data", data2)

	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	// 这里还是用了slice
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	// 	//!+slice
	// 	s := []int{0, 1, 2, 3, 4, 5}
	// 	// Rotate s left by two positions.
	// 	reverse(s[:2])
	// 	reverse(s[2:])
	// 	reverse(s)
	// 	fmt.Println(s) // "[2 3 4 5 0 1]"
	// 	//!-slice

	//	// Interactive test of reverse.
	//	input := bufio.NewScanner(os.Stdin)
	//
	// outer:
	//
	//	for input.Scan() {
	//		var ints []int
	//		for _, s := range strings.Fields(input.Text()) {
	//			x, err := strconv.ParseInt(s, 10, 64)
	//			if err != nil {
	//				fmt.Fprintln(os.Stderr, err)
	//				continue outer
	//			}
	//			ints = append(ints, int(x))
	//		}
	//		reverse(ints)
	//		fmt.Printf("%v\n", ints)
	//	}
	//
	// NOTE: ignoring potential errors from input.Err()
}

// !+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Printf("%v\n", s)
}

func noneempty(strings []string) []string {
	i := 0
	for _, val := range strings {
		if val != "" {
			strings[i] = val
			i++
		}
	}

	return strings[:i]
}

func noneduplicate(strings []string) []string {
	i := 0
	for _, val := range strings {
		hasDuplicate := false
		for _, singleString := range strings[:i] {
			if singleString == val {
				hasDuplicate = true
				break
			}
		}

		if !hasDuplicate {
			strings[i] = val
			i++
		}
	}

	return strings[:i]
}

//!-rev
