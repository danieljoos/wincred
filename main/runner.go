package main

import (
	"fmt"
	winc "github.com/danieljoos/wincred"
)

func main() {
	userNames, serverURLs, err := winc.List()
	fmt.Println(userNames)
	fmt.Println(serverURLs)
	fmt.Println(err)
}


