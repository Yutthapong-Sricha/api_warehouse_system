package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	ba64 "encoding/base64"
	"fmt"
	"io"
)

func main() {
	id_core := "1001"
	id_core_enc := util_encode(id_core)
	id_core_dec := util_decode(id_core_enc)
	secret_key := "spcholding"
	fmt.Println("id_core = ", id_core)
	fmt.Println("id_core_enc = ", id_core_enc)
	fmt.Println("id_core_dec = ", id_core_dec)
	fmt.Println("secrect = ", secret_key)
	b := []byte(secret_key)
	secret_key_byte, err := secret(b)
	fmt.Println("func secrect = ", secret_key_byte)
	if err != nil {
		fmt.Println("func secrect err = ", err.Error())
	}
	id_encrypt, err := encrypt(id_core, secret_key_byte)
	fmt.Println("id_encrypt = ", id_encrypt)
	if err != nil {
		fmt.Println("id_encrypt err = ", err.Error())
	}

	id_decrypt, err := decrypt(id_encrypt, secret_key_byte)
	fmt.Println("id_decrypt = ", id_decrypt)
	if err != nil {
		fmt.Println("id_decrypt err = ", err.Error())
	}
}

func util_encode(id string) string {
	id_encode := ba64.StdEncoding.EncodeToString([]byte(id)) // string to byte
	return id_encode
}

func util_decode(id_enc string) string {
	id_decode, _ := ba64.StdEncoding.DecodeString(id_enc)
	return string(id_decode) // byte to string
}

// secret returns a 32 bytes AES key.
func secret() ([]byte, error) {
	key := make([]byte, 16)

	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}

// encrypt encrypts plain string with a secret key and returns encrypt string.
func encrypt(plainData string, secret []byte) (string, error) {
	cipherBlock, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(aead.Seal(nonce, nonce, []byte(plainData), nil)), nil
}

// decrypt decrypts encrypt string with a secret key and returns plain string.
func decrypt(encodedData string, secret []byte) (string, error) {
	encryptData, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	cipherBlock, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		return "", err
	}

	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainData), nil
}
