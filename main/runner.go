package main

import (
	"fmt"
	winc "github.com/danieljoos/wincred"
)

func main() {
	g := winc.NewGenericCredential("targetname")
	g.UserName = "user"
	g.CredentialBlob = []byte("shhh")
	g.Persist = winc.PersistLocalMachine
	g.Write()
	fmt.Println("hello")
	fmt.Println("list stuff ------------")
	err := winc.List()
	fmt.Println(err)
}


