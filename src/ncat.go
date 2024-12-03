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
	fileContents, err := os.ReadFile(options.FilePath)

	if err != nil {
		if os.IsNotExist(err) {
			Logger.Printf("ncat: file '%s' does not exist. \n", options.FilePath)
		}
		return NcatResult{out: ""}
	}

	return NcatResult{out: string(fileContents)}
}
