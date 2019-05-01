package main

import (
	"fmt"
	"github.com/arnekd/ICA02/fileutils"
)

func main() {

	text1 := fileutils.FileToByteslice("files/lang01.wl")
	text2 := fileutils.FileToByteslice("files/lang02.wl")
	text3 := fileutils.FileToByteslice("files/lang03.wl")

	fmt.Printf("lang01:\n% X\n", text1)
	fmt.Printf("lang02:\n% X\n", text2)
	fmt.Printf("lang03:\n% X\n", text3)

}
