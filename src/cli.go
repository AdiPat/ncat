package src

import "strings"

type Cli struct {
	Ncat NcatInterface
}

type NcatInterface interface {
	Ncat(options NcatOptions) NcatResult
}

func NewCli(ncat NcatInterface) *Cli {
	return &Cli{Ncat: ncat}
}

func (c *Cli) Run(args []string) {
	Logger.SetFlags(0)
	options := NcatOptions{
		FilePath: args[1],
	}
	result := c.Ncat.Ncat(options)
	resultWithNoNewLine := strings.TrimSuffix(result.out, "\n")
	Logger.Printf("%s\n", resultWithNoNewLine)
}
