package main

import (
	"fmt"
	winc "github.com/danieljoos/wincred"
)

func main() {
	//g := winc.NewGenericCredential("targetname1")
	//g.UserName = "user1"
	//g.CredentialBlob = []byte("shhh1")
	//g.Persist = winc.PersistLocalMachine
	//g.Write()
	fmt.Println("hello")
	fmt.Println("list stuff ------------")
	err := winc.List()
	fmt.Println(err)
}


