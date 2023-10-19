package main

import (
	"fmt"
	"github.com/piyushpatel2005/cryptit/decrypt"
	"github.com/piyushpatel2005/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("ABC")
	decryptedStr := decrypt.Nimbus(encryptedStr)
	fmt.Println(encryptedStr, decryptedStr)
}
