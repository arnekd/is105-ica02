package unicode

import (
	"fmt"
)

// Kode for Oppgave 4a
func Translate(expression string, language string) string {

	uni := ""

	if expression == "nord og sør" {
		if language == "jp" {
			uni = "\xe2\x80\x9c" + expression + "\xe2\x80\x9d" + " på japansk: " +
				"\x22\xe5\x8c\x97\xe3\x81\xa8\xe5\x8d\x97\x22"
			//22 E5 8C 97 E3 81 A8 E5 8D 97 22
		} else if language == "is" {
			uni = "\xe2\x80\x9c" + expression + "\xe2\x80\x9d" + " på islandsk: " +
				"\x22\x6E\x6F\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72\x22"
			//22 6E 6F 72 C3 B0 75 72 20 6F 67 20 73 75 C3 B0 75 72 22
		}
	}
	return uni
}

// Kode for Oppgave 4b
func UnicodeCodePointDemo() {
	// Hva er dette?
	// Er det likt på MS Windows og macOS?
	fmt.Println("\xf0\x9F\x98\x80")
	fmt.Println("\xf0\x9F\x98\x97")
	// Demonstrerer at deler av et tegn representert med flere bytes
	// kan ikke tolkes innenfor koden (unicode)
	fmt.Println("\xf0\x9F\x98")
	fmt.Println("\xf0\x9F")
	fmt.Println("\xf0")
}
