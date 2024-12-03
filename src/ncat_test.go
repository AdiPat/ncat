package src

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNcat(t *testing.T) {
	testFilePath := "../test_data/test.txt"
	var (
		buff       bytes.Buffer
		MockLogger = log.New(&buff, "", log.LstdFlags)
	)

	t.Run("ncat", func(t *testing.T) {

		t.Run("accepts file path and returns the file text", func(t *testing.T) {
			expectedLines := []string{
				`"Your heart is the size of an ocean. Go find yourself in its hidden depths."`,
				`"The Bay of Bengal is hit frequently by cyclones. The months of November and May, in particular, are dangerous in this regard."`,
				`"Thinking is the capital, Enterprise is the way, Hard Work is the solution."`,
				`"If You Can'T Make It Good, At Least Make It Look Good."`,
				`"Heart be brave. If you cannot be brave, just go. Love's glory is not a small thing."`,
				`"It is bad for a young man to sin; but it is worse for an old man to sin."`,
				`"If You Are Out To Describe The Truth, Leave Elegance To The Tailor."`,
				`"O man you are busy working for the world, and the world is busy trying to turn you out."`,
				`"While children are struggling to be unique, the world around them is trying all means to make them look like everybody else."`,
				`"These Capitalists Generally Act Harmoniously And In Concert, To Fleece The People."`,
			}
			expected := strings.Join(expectedLines, "\n") + "\n"

			ncatOptions := NcatOptions{
				FilePath: testFilePath,
			}

			ncatClient := NcatClient{}
			actual := ncatClient.Ncat(ncatOptions)

			assert.Equal(t, expected, actual.out)
		})

		t.Run("logs an error message if the file is not found", func(t *testing.T) {
			Logger = MockLogger
			ncatOptions := NcatOptions{
				FilePath: "non_existent_file.txt",
			}
			ncatClient := NcatClient{}
			ncatClient.Ncat(ncatOptions)
			actual := buff.String()
			expected := "ncat: file 'non_existent_file.txt' does not exist. \n"
			assert.Contains(t, actual, expected)
		})

	})

}
