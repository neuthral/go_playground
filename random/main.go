package main

import (
	"fmt"
	"math/rand"
)

func minmax() {
	// initializing with the maximum int64
	var max int64 = 9223372036854775807
	var min int64 = max + 1 // -9223372036854775808
	// printing the value and data type
	fmt.Printf("Value is: %d and type is: %T\n", max, max)
	// printing out new value and type
	fmt.Printf("Value is: %d and type is: %T\n", min, min)
}

func rebuildArray(seed int, modulus int, ll int) []int {
	const MAX int64 = 9223372036854775807
	aa := make([]int, ll)
	rand.Seed(MAX)
	for i := range aa {
		aa[i] = rand.Intn(int(MAX)) % modulus
		fmt.Print(aa[i], " ")
	}
	return aa
}

func main() {
	data := []byte("hello")
	mod := len(data)

	fmt.Println(data)
	for i := 0; i < mod; i++ {
		new := int(data[i])
		fmt.Println(new % mod)
	}

}
