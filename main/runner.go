package main

import (
	"fmt"
	winc "github.com/danieljoos/wincred"
)

func main() {
	err := winc.List()
	fmt.Println(err)
}


