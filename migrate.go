package main

import (
	"cp_migrator/walk"
	"os"
)

func main() {
	args := os.Args[1:]

	walk.CssConverter(args[0])
	// css_converter.ConvertScssVarToCustomProps(cssFiles)
}
