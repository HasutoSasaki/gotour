package main

import "fmt"

func xorEncryptDecrypt(data []byte, key byte) []byte {
	result := make([]byte, len(data))
	for i, b := range data {
		result[i] = b ^ key
	}
	return result
}

func main() {
	message := []byte("hello world")
	key := byte(42) // 適当な1バイトの鍵

	encrypted := xorEncryptDecrypt(message, key)
	fmt.Printf("暗号化: %v\n", encrypted)

	decrypted := xorEncryptDecrypt(encrypted, key)
	fmt.Printf("復号: %s\n", decrypted)
}
