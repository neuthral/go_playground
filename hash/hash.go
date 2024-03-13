// "string", len(string) -->
//			(hash)1234 --> hash, len(hash)
//			--> newHash

package main

import (
	"fmt"
	"strings"
)

func customHash(data string) int64 {
	arr := []byte(data)
	length := int64(len(data))
	var hash int64 = 0
	for _, char := range arr {
		hash += int64(char) % length
		length -= 1
		fmt.Println(hash)
	}
	return hash
}

func bloom(data string) (byteArr []byte, charArr []string, intValu int) {
	byteArr = []byte(data)
	charArr = strings.Split(data, "")
	intValu = int(customHash(data))

	return
}

func main() {
	//inputData := "hello world"
	//customHashValue := customHash(inputData)
	//fmt.Println(customHashValue, "length ", int64(len(inputData)))

	fmt.Println(bloom("worldo"))

}
