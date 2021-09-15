package main

import (
	"fmt"
	"github.com/wumansgy/goEncrypt"
)

func main(){
	hash := goEncrypt.Sha512Hex([]byte("7b60eed750ec5e9d0dbfac48004e6ebb2104287067727713mobile=2104287067727713&timestamp=2104287067727713&userId=2104287067727713&userRole=1"))
	fmt.Println(hash)
}
