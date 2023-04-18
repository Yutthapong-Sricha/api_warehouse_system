package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"os"
)

var (
// We're using a 32 byte long secret key.
// This is probably something you generate first
// then put into and environment variable.
// secretKey string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
// secretKey string = "IamSreepreechaHolding1FebMay2023"
// secretKey string = "AmSpcHolding2023"
// secretKey string = os.Getenv("SECRETKEY")
)

func Encrypt(plaintext string) string {
	secretKey := os.Getenv("SECRETKEY")
	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	// ciphertext here is actually nonce+ciphertext
	// So that when we decrypt, just knowing the nonce size
	// is enough to separate it from the ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	//return string(ciphertext)
	return hex.EncodeToString(ciphertext)
}

func Decrypt(hex_ciphertext string) string {
	secretKey := os.Getenv("SECRETKEY")
	ciphertext, _ := hex.DecodeString(hex_ciphertext)
	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}

// func Hsession() {
// 	store := cookie.NewStore([]byte(secretKey))
// 	sessions.Sessions("SPCSession", store)
// 	return sessions
// }
