package src

import (
	"bytes"
	"io"
	"os"
)

func ReadFromStdin() (string, error) {
	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)
	return buf.String(), nil
}
