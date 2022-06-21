package crypt_laravel

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"regexp"
)

type Token struct {
	Iv    string `json:"iv"`
	Value string `json:"value"`
	Mac   string `json:"mac"`
	Tag   string `json:"tag"`
}

func EncryptString(privateKey string, plaintext string) (token string) {
	var resultToken Token
	iv := randomString(16)
	value := aes256(plaintext, privateKey, iv, aes.BlockSize)

	ivBase64 := base64.StdEncoding.EncodeToString([]byte(iv))
	tagBase64 := base64.StdEncoding.EncodeToString([]byte(""))

	resultToken.Iv = ivBase64
	resultToken.Tag = tagBase64
	resultToken.Value = value
	mac := hash(ivBase64, value, privateKey)
	resultToken.Mac = mac
	jsonResult, err := json.Marshal(resultToken)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(jsonResult)
}

func hash(iv string, value string, key string) (hash string) {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(iv + value))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func aes256(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := pKCS5Padding([]byte(plaintext), blockSize)
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func DecryptString(privateKey string, token string) (str string) {
	key := []byte(privateKey)
	var decodedByte, _ = base64.StdEncoding.DecodeString(token)
	var decodedString = string(decodedByte)

	var resultToken Token
	jsonString := []byte(decodedString)
	err := json.Unmarshal(jsonString, &resultToken)

	if err != nil {
		return ""
	}

	var iv, _ = base64.StdEncoding.DecodeString(resultToken.Iv)

	ciphertext, _ := base64.StdEncoding.DecodeString(resultToken.Value)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	reg, _ := regexp.Compile("[^0-9]+")
	processedString := reg.ReplaceAllString(string(ciphertext), "")

	return processedString
}
