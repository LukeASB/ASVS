package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(val []byte, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)

	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return aead.Seal(nonce, nonce, val, nil), nil
}

func Decrypt(val []byte, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	size := aead.NonceSize()
	if len(val) < size {
		return nil, err
	}

	result, err := aead.Open(nil, val[:size], val[size:], nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Secret() ([]byte, error) {
	key := make([]byte, 16)

	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}
