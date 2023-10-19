package main

import (
	"fmt"
	"github.com/pborman/uuid"
	"github.com/piyushpatel2005/cryptit/decrypt"
	"github.com/piyushpatel2005/cryptit/encrypt"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)

	encryptedStr := encrypt.Nimbus("ABC")
	decryptedStr := decrypt.Nimbus(encryptedStr)
	fmt.Println(encryptedStr, decryptedStr)
}
