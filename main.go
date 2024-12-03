package main

import (
	"os"

	core "github.com/AdiPat/ncat/src"
)

func Ncat(options core.NcatOptions) core.NcatResult {
	ncatClient := core.NcatClient{}
	return ncatClient.Ncat(options)
}

func main() {
	ncatClient := core.NcatClient{}
	cli := core.NewCli(ncatClient)
	cli.Run(os.Args)
}
