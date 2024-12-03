package main

import (
	"os"

	core "github.com/AdiPat/ncat/src"
)

func main() {
	ncatClient := core.NcatClient{}
	cli := core.NewCli(ncatClient)
	cli.Run(os.Args)
}
