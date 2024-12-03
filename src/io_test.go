package src

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockStdin(input string) func() {
	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = r

	go func() {
		defer w.Close()
		w.Write([]byte(input))
	}()

	return func() {
		os.Stdin = oldStdin
	}
}

func TestReadFromStdin(t *testing.T) {
	t.Run("reads the input from stdin", func(t *testing.T) {
		expectedInput := "Hello, ncat!\n"
		restoreStdin := mockStdin(expectedInput)
		actualInput, _ := ReadFromStdin()
		assert.Equal(t, expectedInput, actualInput)
		defer restoreStdin()
	})
}
