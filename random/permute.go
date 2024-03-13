package main

import (
	"fmt"
	"os"
)

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}

	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// permute all combinations of input string

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Generate all permutation of PATTERN \n")
		fmt.Println("Usage:", os.Args[0], "PATTERN", "| example {abcde}")
		return
	}
	str := os.Args[1]
	Perm([]rune(str), func(a []rune) {
		fmt.Println(string(a))
	})
}
