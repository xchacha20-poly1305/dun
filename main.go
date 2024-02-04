package main

import (
	"fmt"

	"github.com/xchacha20-poly1305/dun/dunbox"
	"github.com/xchacha20-poly1305/dun/dunmain"
	 _ "github.com/xchacha20-poly1305/dun/dunapi"
)

func main() {
	fmt.Println("dun:", dunbox.Version)
	fmt.Println()

	// dunmain
	dunmain.Main()
}
