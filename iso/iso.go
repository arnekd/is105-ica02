package iso

import (
	//"encoding/hex"
	"fmt"
	//"log"
	"strings"
)

// Oppgave 2a
// Lag selv en "string literal" med utvidet ASCII
// Blir det kun en slik "string literal" eller trenger man flere
// for å representere utvidet ASCII?

const extendedAscii = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89" +
	"\x8A\x8B\x8C\x8D\x8E\x8F" +
	"\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99" +
	"\x9A\x9B\x9C\x9D\x9E\x9F" +
	"\xA0\xA1\xA2\xA3\xA4\xA5\xA6\xA7\xA8\xA9" +
	"\xAA\xAB\xAC\xAD\xAE\xAF" +
	"\xB0\xB1\xB2\xB3\xB4\xB5\xB6\xB7\xB8\xB9" +
	"\xBA\xBB\xBC\xBD\xBE\xBF" +
	"\xC0\xC1\xC2\xC3\xC4\xC5\xC6\xC7\xC8\xC9" +
	"\xCA\xCB\xCC\xCD\xCE\xCF" +
	"\xD0\xD1\xD2\xD3\xD4\xD5\xD6\xD7\xD8\xD9" +
	"\xDA\xDB\xDC\xDD\xDE\xDF" +
	"\xE0\xE1\xE2\xE3\xE4\xE5\xE6\xE7\xE8\xE9" +
	"\xEA\xEB\xEC\xED\xEE\xEF" +
	"\xF0\xF1\xF2\xF3\xF4\xF5\xF6\xF7\xF8\xF9" +
	"\xFA\xFB\xFC\xFD\xFE\xFF"

// IterateOverASCIIStringLiteral tar en "string literal" som INN-data
func IterateOverASCIIExtendedStringLiteral(stringLiteral string) {
	// Kode for Oppgave 2a

	for a := range stringLiteral { // itererer gjennom hele stringLiteral

		s := fmt.Sprintf("%c", a) // setter s lik ASCII-symbolet, hvis det finnes

		if len(strings.TrimSpace(s)) != 0 { // strings.TrimSpace(s) fjerner alle istanser av mellomrom(" ") https://stackoverflow.com/questions/18594330/what-is-the-best-way-to-test-for-an-empty-string-in-go
			fmt.Printf("%X %c %b\n", stringLiteral[a], stringLiteral[a], stringLiteral[a]) // %X gir heksades, %c gir symbol, %b gir den binære koden
		} else { // utføres hvis ASCII tegnet ikke har et symbol
			fmt.Printf("%X %b\n", stringLiteral[a], stringLiteral[a])
		}
	}
}

func GetExtendedASCIIStringLiteral() string {
	return extendedAscii
}

// GreetingExtendedASCII returnerer en tekst-streng i
// utvidet ASCII
// Kode for Oppgave 2c
func GreetingExtendedASCII() string {

	s := []byte("\"Salut, ça va °-) Ça coute €50\"")
	s2 := []byte("\"Salut, ça va °-) Κοστίζει €50\"")

	fmt.Printf("\"Salut, ça va °-) Κοστίζει ​€50​\" til heksadesimal verdi: % X\n", s2)

	return string(s[:])
}

// Skriver ut "Salut, ça va °-) Κοστίζει €50" fra den heksadesimale verdien
func PrintFromHex() {

	h := []byte("\x22\x53\x61\x6C\x75\x74\x2C\x20\x63\xCC\xA7" +
		"\x61\x20\x76\x61\x20\xC2\xB0\x2D\x29\x20\xCE\x9A\xCE\xBF" +
		"\xCF\x83\xCF\x84\xCE\xB9\xCC\x81\xCE\xB6\xCE\xB5\xCE\xB9\x20\xE2\x82\xAC\x35\x30\x22")

	fmt.Printf("%s\n", h)

}
