package main

import (
	"fmt"
	"github.com/arnekd/ICA02/unicode"
)

func main() {

	const islandsk = "\"norður og suður\""
	const japansk = "\"北と南\""

	fmt.Printf("\"norður og suður\" til heksadesimal: % X\n", islandsk)
	fmt.Printf("\"北と南\" til heksadesimal: % X\n", japansk)
	fmt.Println()

	jp := unicode.Translate("nord og sør", "jp")
	is := unicode.Translate("nord og sør", "is")

	fmt.Println(jp)
	fmt.Println(is)
	fmt.Println()

	unicode.UnicodeCodePointDemo()

}
