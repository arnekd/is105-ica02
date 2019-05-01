package main

import (
	"fmt"
	"github.com/arnekd/ICA02/unicode"
)

func main() {

	jp := unicode.Translate("nord og sør", "jp")
	is := unicode.Translate("nord og sør", "is")

	fmt.Println(jp)
	fmt.Println(is)

	unicode.UnicodeCodePointDemo()

}
