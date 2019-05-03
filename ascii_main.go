package main

import (
	//"fmt"
	"github.com/arnekd/ICA02/ascii"
)

func main() {

	//ascii.IterateOverASCIIStringLiteral("ascii.ascii")

	strLiteral := ascii.GetASCIIStringLiteral()

	ascii.GreetingASCII()
	ascii.IterateOverASCIIStringLiteral(strLiteral)

}
