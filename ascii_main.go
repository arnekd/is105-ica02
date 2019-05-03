package main

import (
	"github.com/arnekd/ICA02/ascii"
)

func main() {

	strLiteral := ascii.GetASCIIStringLiteral()

	ascii.GreetingASCII()
	ascii.IterateOverASCIIStringLiteral(strLiteral)

}
