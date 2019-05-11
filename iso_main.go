package main

import (
	"fmt"
	"github.com/arnekd/ICA02/iso"
)

func main() {

	iso.IterateOverASCIIExtendedStringLiteral(iso.GetExtendedASCIIStringLiteral())

	fmt.Println("Salut, ça va °-) Κοστίζει ​€50​")
	iso.GreetingExtendedASCII()
	fmt.Println("Fra hexadesimal verdi:")
	iso.PrintFromHex()

}
