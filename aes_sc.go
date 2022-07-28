package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
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

func main() {
	str := "payload"
	key := []byte("LeslieCheungKwok")
	src := EncyptogAES([]byte(str), key)
	base64Str := base64.StdEncoding.EncodeToString(src)
	fmt.Println("加密后的数据为:", base64Str)

}
