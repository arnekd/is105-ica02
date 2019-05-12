package main

import (
	"fmt"
	"github.com/arnekd/ICA02/fileutils"
)

func main() {

	text1 := fileutils.FileToByteslice("files/lang01.wl")
	text2 := fileutils.FileToByteslice("files/lang02.wl")
	text3 := fileutils.FileToByteslice("files/lang03.wl")

	fmt.Printf("lang01.wl:\n% c\n", text1)
	fmt.Printf("lang02.wl:\n% c\n", text2)
	fmt.Printf("lang03.wl:\n% c\n", text3)

	fmt.Println()

	//text1first16 := "\xEF\xDA\xA3\xD2\xD3\xCB\x0A\xEF\xDA\xA3\xD2\xD3\xCB\xC1\x0A\xEF"
	//fmt.Printf("%c\n", byte[]("\xEF"))
	fmt.Printf("%s", "\xc3\xaf\xc3\x9a")

	//1110 1111 --> 11000011 10101111 --> C3 AF
	//1101 1010 --> 11000011 10011010 --> C3 9A

	fmt.Println()

	fmt.Printf("Original: %s\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Printf("UTF-8: %s\n", "\xc2\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")

}
