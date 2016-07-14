package main

import (
	"fmt"
	winc "github.com/danieljoos/wincred"
)

func main() {
	fmt.Println("hello")
	g := winc.NewGenericCredential("credzzzz")
	g.UserName = "test"
	g.CredentialBlob = []byte("abcdefghijlmnop")
	g.Persist = winc.PersistLocalMachine
	g.Write()
}


