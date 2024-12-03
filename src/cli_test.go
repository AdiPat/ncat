package src

import (
	"testing"

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
	})
}
