package main

import (
	"fmt"
	_ "unsafe"

	_ "github.com/xchacha20-poly1305/dun/distro/all"
	"github.com/xchacha20-poly1305/dun/dunbox"
	"github.com/xchacha20-poly1305/dun/dunmain"
)

func main() {
	fmt.Println("sing-box-extra:", dunbox.Version)
	fmt.Println()

	// sing-box
	dunmain.Main()
}
