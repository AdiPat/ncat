package src

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNcat struct {
	mock.Mock
}

func (m *MockNcat) Ncat(options NcatOptions) NcatResult {
	args := m.Called(options)
	return args.Get(0).(NcatResult)
}

func TestCli(t *testing.T) {
	var (
		buff       bytes.Buffer
		MockLogger = log.New(&buff, "", log.LstdFlags)
	)

	t.Run("reads file and calls Ncat", func(t *testing.T) {
		testFilePath := "../test_data/test.txt"

		t.Run("accepts file path and returns the file text", func(t *testing.T) {
			mockNcat := new(MockNcat)
			options := NcatOptions{
				FilePath: testFilePath,
			}
			expected := NcatResult{
				out: "file contents",
			}
			mockNcat.On("Ncat", options).Return(expected)
			args := []string{"ncat", testFilePath}
			cli := NewCli(mockNcat)
			cli.Run(args)
			mockNcat.AssertCalled(t, "Ncat", options)
		})

		t.Run("logs the 'output' to stdout", func(t *testing.T) {
			Logger = MockLogger
			mockNcat := new(MockNcat)
			options := NcatOptions{
				FilePath: testFilePath,
			}
			expected := NcatResult{
				out: "file contents",
			}
			mockNcat.On("Ncat", options).Return(expected)
			args := []string{"ncat", testFilePath}
			cli := NewCli(mockNcat)
			cli.Run(args)
			actual := buff.String()
			expectedLog := "file contents\n"
			assert.Equal(t, expectedLog, actual)
		})
	})
}
