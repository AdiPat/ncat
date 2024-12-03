package src

import (
	"os"
)

type NcatOptions struct {
	FilePath string
}

type NcatResult struct {
	out string
}

func Ncat(options NcatOptions) NcatResult {
	fileContents, _ := os.ReadFile(options.FilePath)
	return NcatResult{out: string(fileContents)}
}
