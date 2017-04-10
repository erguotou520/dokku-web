package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

// GenerateSalt generate byte[] salt
func GenerateSalt() []byte {
	c := 16
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
	return b
}

// EncryptPassword encrypt password with salt
func EncryptPassword(pwd string, salt []byte) string {
	dk := pbkdf2.Key([]byte(pwd), salt, 4096, 32, sha1.New)
	fmt.Println(dk)
	return fmt.Sprintf("%x", dk)
}
