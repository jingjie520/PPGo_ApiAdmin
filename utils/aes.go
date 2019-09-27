package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

var AesKey string = "hgfedcba12345678"
var Gsm_IV string = "000010000000010000000100"

func AesEncrypt(text string) string {
	return encrypt(text)
}

func AesDecrypt(text string) string {
	return decrypt(text)
}

func getKeys() []byte {
	return []byte(AesKey)
}

func getBlock() cipher.Block {
	key := getKeys()
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	return block
}

// getNonce 加密用的nonce
func getNonce() []byte {
	nonce, err := hex.DecodeString(Gsm_IV)
	if err != nil {
		panic(err.Error())
	}
	return nonce
}

func encrypt(data string) string {
	block := getBlock()
	nonce := getNonce()
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	cipherText := aesgcm.Seal(nil, nonce, []byte(data), nil)
	return fmt.Sprintf("%x", cipherText)
}

func decrypt(data string) string {
	cipherText, _ := hex.DecodeString(data)

	nonce := getNonce()
	block := getBlock()

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%x", plaintext)
}
