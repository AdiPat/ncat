package src

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
	options := NcatOptions{
		FilePath: args[1],
	}
	c.Ncat.Ncat(options)
}
