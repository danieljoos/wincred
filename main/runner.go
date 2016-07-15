package main

import (
	"fmt"
	winc "github.com/danieljoos/wincred"
)

func main() {
	g, _ := winc.GetGenericCredential("targetname")
	if g == nil {
		fmt.Println("not found")
	}
	return g.UserName, string(g.CredentialBlob), nil
	fmt.Println("hello")
	fmt.Println("list stuff ------------")
	err := winc.List()
	fmt.Println(err)
}


