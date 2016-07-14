package main

import (
	"fmt"
	winc "github.com/danieljoos/wincred"
)

func main() {
	fmt.Println("hello")
	g, err := winc.GetGenericCredential("credzzzz")
	fmt.Println(err)
	if g == nil {
		fmt.Println("error")
	}
	fmt.Println(g.UserName)
	fmt.Println(string(g.CredentialBlob))
	fmt.Println("list stuff ------------")
	err = winc.List()
	fmt.Println(err)
}


