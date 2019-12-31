/***************************************************************************
 *
 * Copyright (c) 2019 github.com, Inc. All Rights Reserved
 * rc6 encrypt
 * Author scloudrun
 *
**************************************************************************/
package rc6lib

import (
	"bytes"
	"crypto/cipher"
	"errors"

	rc6 "github.com/dgryski/go-rc6"
)

func Rc6Enc(key, iv, plantText []byte, paddingStatus bool) ([]byte, error) {
	block, err := rc6.New(key)
	if err != nil {
		return nil, err
	}

	if paddingStatus {
		plantText = pKCS7Padding(plantText, block.BlockSize())
	}

	blockModel := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plantText))
	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func Rc6Dec(key, iv, ciphertext []byte, paddingStatus bool) ([]byte, error) {
	block, err := rc6.New(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < block.BlockSize() {
		return nil, errors.New("crypto/cipher: ciphertext too short")
	}

	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, errors.New("crypto/cipher: ciphertext is not a multiple of the block size")
	}

	blockModel := cipher.NewCBCDecrypter(block, iv)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	if paddingStatus {
		plantText, err = pKCS7UnPadding(plantText, block.BlockSize())
	}
	if err != nil {
		return nil, err
	}
	return plantText, nil
}

func pKCS7UnPadding(plantText []byte, blockSize int) ([]byte, error) {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	if length-unpadding < 0 || length-unpadding > length {
		return nil, errors.New("aes unpadding error")
	}
	return plantText[:(length - unpadding)], nil
}

func pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
