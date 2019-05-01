package main

import (
	//"fmt"
	"github.com/arnekd/ICA02/ascii"
)

func main() {

	//ascii.IterateOverASCIIStringLiteral("ascii.ascii")

	strLiteral := ascii.GetASCIIStringLiteral()

	//fmt.Println(strLiteral)
	ascii.IterateOverASCIIStringLiteral(strLiteral)

}
