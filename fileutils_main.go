package main

import (
	"fmt"
	"github.com/arnekd/ICA02/fileutils"
	//"strings"
)

func main() {

	text1 := fileutils.FileToByteslice("files/lang01.wl")
	text2 := fileutils.FileToByteslice("files/lang02.wl")
	text3 := fileutils.FileToByteslice("files/lang03.wl")

	fmt.Printf("lang01.wl:\n% c\n", text1)
	fmt.Printf("lang02.wl:\n% c\n", text2)
	fmt.Printf("lang03.wl:\n% c\n", text3)

	fmt.Println()

	// Lagrer de 16 første bytes i hver fil
	const text1First16 = "\xEF\xDA\xA3\xD2\xD3\xCB\x0A\xEF\xDA\xA3\xD2\xD3\xCB\xC1\x0A\xEF"
	const text2First16 = "\xFE\xFD\x73\x6B\x61\x72\x0A\xFE\xFD\x73\x6B\x61\x72\x61\x6E\x61"
	const text3First16 = "\xF8\x79\x65\x73\x70\x65\x73\x69\x61\x6C\x69\x73\x74\x65\x6E\x0A"

	fmt.Println("De første 16 bytes\ni filen lang01.wl:")

	// Skriver ut informasjon om de 16 første bytes i lang01.wl
	for a := range text1First16 {
		fmt.Printf("%c %U\n", text1First16[a], text1First16[a])
	}

	fmt.Println("\nDe første 16 bytes\ni filen lang02.wl:")

	for b := range text2First16 {
		fmt.Printf("%c %U\n", text2First16[b], text2First16[b])
	}

	fmt.Println("\nDe første 16 bytes\ni filen lang03.wl:")

	for c := range text3First16 {
		fmt.Printf("%c %U\n", text3First16[c], text3First16[c])
	}

	fmt.Println()

	fmt.Printf("Original: %s\n", "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	fmt.Printf("UTF-8: %s\n", "\xc2\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")

}
