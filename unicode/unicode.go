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
				"\xe2\x80\x9c\xe5\x8c\x97\xe3\x81\xa8\xe5\x8d\x97\xe2\x80\x9d" //e2 80 9c e5 8c 97 e3 81 a8 e5 8d 97 e2 80 9d
		} else if language == "is" {
			uni = "\xe2\x80\x9c" + expression + "\xe2\x80\x9d" + " på islandsk: " +
				"\xe2\x80\x9c\x6e\x6f\x72\xc3\xb0\x75\x72\x20\x6f\x67\x20\x73\x75\xc3\xb0\x75\x72\xe2\x80\x9d"
			//e2 80 9c 6e 6f 72 c3 b0 75 72 20 6f 67 20 73 75 c3 b0 75 72 e2 80 9d
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
