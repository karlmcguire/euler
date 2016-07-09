package main

import (
	"fmt"
)

func pali(n int) bool {
	s := []byte(fmt.Sprintf("%d", n))
	l := len(s)

	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}

	return true
}

func main() {
	p := 0

	for x := 100; x < 1000; x++ {
		for y := 100; y < 1000; y++ {
			t := x * y
			if t > p && pali(t) {
				p = t
			}
		}
	}

	println(p)
}
