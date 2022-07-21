package wordwrap_test

import (
	"os"
	"testing"

	"code-katas/wordwrap"

	"github.com/alecthomas/assert/v2"
)

func TestWordWrap(t *testing.T) {
	t.Run("wrap the words to a max limit of chars", func(t *testing.T) {
		input, err := os.ReadFile("test_input.txt")
		assert.NoError(t, err)

		expectedOutput, err := os.ReadFile("test_output.txt")
		assert.NoError(t, err)

		gotOutput := wordwrap.Wrap(string(input), 58)
		assert.Equal(t, string(expectedOutput), gotOutput)
	})

	t.Run("if the max width is empty, return nothing", func(t *testing.T) {
		wrapped := wordwrap.Wrap("This kata is actually quite annoying.\n\nThis is another para that needs to be wrapped", 0)
		assert.Equal(t, "", wrapped)
	})

	t.Run("if the string is empty, return nothing", func(t *testing.T) {
		wrapped := wordwrap.Wrap("", 10)
		assert.Equal(t, "", wrapped)
	})
}
