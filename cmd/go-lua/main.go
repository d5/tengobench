package main

import (
	"os"

	"github.com/Shopify/go-lua"
)

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)

	if err := lua.DoFile(l, os.Args[1]); err != nil {
		panic(err)
	}
}
