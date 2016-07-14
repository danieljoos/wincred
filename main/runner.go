package main

import (
	"fmt"
	winc "github.com/danieljoos/wincred"
)

func main() {
	fmt.Println("hello")
	fmt.Println("list stuff ------------")
	err := winc.List()
	fmt.Println(err)
}


