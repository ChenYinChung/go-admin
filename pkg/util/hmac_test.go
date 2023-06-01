package util

import (
	"fmt"
	"testing"
)

func TestHmac(t *testing.T) {
	// secret 是加密要使用的key
	// apple 是要加密的内容
	fmt.Printf("HmacSha256轉16字符串: %s\n", HmacSha256ToHex("secret", "apple"))
	fmt.Printf("HmacSha256轉base字符串: %s\n", HmacSha256ToBase64("secret", "apple"))
}
