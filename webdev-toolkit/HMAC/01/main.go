package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

//Keyed-Hash Message Authentication Code
//Encrypts the message from the user

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)

	c = getCode("teste@exemplo.com")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
