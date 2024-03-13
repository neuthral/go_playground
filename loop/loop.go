package main

import (
	"fmt"
	"math/rand"
)

func init() {
	//var seed int64
	seed := int64(999)
	rand.Seed(seed)
}
func main() {
	pos := 0

	for {
		pos += 1
		fmt.Print()
		hash := rand.ExpFloat64()
		fmt.Print(hash)
		break
	}
}
