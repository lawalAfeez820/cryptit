package main 

import (
	"fmt"
	"github.com/lawalAfeez820/cryptit/encrypt"
	"github.com/lawalAfeez820/cryptit/decrypt"
)


func main(){

	encryptedstr := encrypt.Nimbus("Afeez")

	fmt.Println(encryptedstr)
	fmt.Println(decrypt.Nimbus(encryptedstr))
}
