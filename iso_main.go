package main

import (
	"fmt"
	"github.com/arnekd/ICA02/iso"
)

func main() {

	iso.IterateOverASCIIExtendedStringLiteral(iso.GetExtendedASCIIStringLiteral())

	fmt.Println("Salut, ça va °-) Κοστίζει ​€50​")

	fmt.Printf("%s\n", "\xa5")

}
