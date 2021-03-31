package parser_test

import (
	"github.com/adammck/venv"
	"github.com/lbrun25/fenv/parser"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseEnvironmentVariables(t *testing.T) {
	var err error
	content := []byte(`
variables:
	first: ${FIRST}
	second: ${SECOND}
	third:
		value: ${THIRD}
`)

	t.Run("success", func(t *testing.T) {
		expectedContent := []byte(`
variables:
	first: first_value
	second: second_value
	third:
		value: third_value
`)

		filePath := "a"
		err = afero.WriteFile(parser.Ufs, filePath, content, 0644)
		assert.NoError(t, err)

		// Mock environment variables
		e := venv.Mock()
		err = e.Setenv("FIRST", "first_value")
		assert.NoError(t, err)
		err = e.Setenv("SECOND", "second_value")
		assert.NoError(t, err)
		err = e.Setenv("THIRD", "third_value")
		assert.NoError(t, err)

		newContent, err := parser.EnvironmentVariables(filePath, e)
		assert.NoError(t, err)

		assert.Equal(t, string(expectedContent), string(newContent))
	})

	// No return line at the end of the file
	t.Run("success-2", func(t *testing.T) {
		originalContent := []byte(`
variables:
	first: ${FIRST}
	second: ${SECOND}
	third: ${THIRD}`)
		expectedContent := []byte(`
variables:
	first: first_value
	second: second_value
	third: third_value`)

		filePath := "a"
		err = afero.WriteFile(parser.Ufs, filePath, originalContent, 0644)
		assert.NoError(t, err)

		// Mock environment variables
		e := venv.Mock()
		err = e.Setenv("FIRST", "first_value")
		assert.NoError(t, err)
		err = e.Setenv("SECOND", "second_value")
		assert.NoError(t, err)
		err = e.Setenv("THIRD", "third_value")
		assert.NoError(t, err)

		newContent, err := parser.EnvironmentVariables(filePath, e)
		assert.NoError(t, err)

		assert.Equal(t, string(expectedContent), string(newContent))
	})

	t.Run("error-file", func(t *testing.T) {
		err = afero.WriteFile(parser.Ufs, "a", content, 0644)
		assert.NoError(t, err)

		_, err := parser.EnvironmentVariables("b", venv.Mock())
		assert.Error(t, err)
	})

	t.Run("error-env-variable", func(t *testing.T) {
		filePath := "a"
		err = afero.WriteFile(parser.Ufs, filePath, content, 0644)
		assert.NoError(t, err)

		_, err := parser.EnvironmentVariables(filePath, venv.Mock())
		assert.Error(t, err)
	})

	t.Run("error-parsing", func(t *testing.T) {
		contentError := []byte(`
variables:
	first: ${FIRST}
	second: ${SECOND
	third:
		value: ${THIRD}
`)

		filePath := "a"
		err = afero.WriteFile(parser.Ufs, filePath, contentError, 0644)
		assert.NoError(t, err)

		// Mock environment variables
		e := venv.Mock()
		err = e.Setenv("FIRST", "first_value")
		assert.NoError(t, err)
		err = e.Setenv("SECOND", "second_value")
		assert.NoError(t, err)
		err = e.Setenv("THIRD", "third_value")
		assert.NoError(t, err)

		_, err := parser.EnvironmentVariables(filePath, e)
		assert.Error(t, err)
	})
}