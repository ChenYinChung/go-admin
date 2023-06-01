package main

import (
	"fmt"
	"go-admin/pkg/util"
)

func main() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[2:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
	fmt.Printf("HmacSha256轉16字符串: %s\n", util.HmacSha256ToHex("secret", "apple"))
	fmt.Printf("HmacSha256轉base字符串: %s\n", util.HmacSha256ToBase64("secret", "apple"))

}
