package parser

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/adammck/venv"
	"github.com/spf13/afero"
)

var (
	baseFs   = afero.NewOsFs()
	roBaseFs = afero.NewReadOnlyFs(baseFs)
	Ufs      = afero.NewCopyOnWriteFs(roBaseFs, afero.NewMemMapFs())
)

// findAllOccurrences gets the first index of occurrences of a string contained within a byte array
func findAllOccurrences(data []byte, searches []string) map[string][]int {
	results := make(map[string][]int)
	for _, search := range searches {
		searchData := data
		term := []byte(search)
		for x, d := bytes.Index(searchData, term), 0; x > -1; x, d = bytes.Index(searchData, term), d + x + 1 {
			results[search] = append(results[search], x + d)
			searchData = searchData[x + 1:]
		}
	}
	return results
}

// EnvironmentVariables parse the file and replace any environment variable by its value
func EnvironmentVariables(filePath string, e venv.Env) ([]byte, error) {
	content, err := afero.ReadFile(Ufs, filePath); if err != nil {
		return nil, err
	}
	startIndexes := findAllOccurrences(content, []string{"$"})

	for j, startIndex := range startIndexes["$"] {
		var i int
		for i = startIndex; content[i] != '}'; i++ {
			if i >= len(content) - 1 {
				return nil, errors.New("error when parsing environment variables")
			}
		}

		if content[i] == '}' && i != len(content) {
			variable := string(content[startIndex+2:i])
			value := e.Getenv(variable)
			if len(value) == 0 {
				msg := fmt.Sprintf("'%s' mandatory env variable not found", variable)
				return nil, errors.New(msg)
			}
			enclosure := fmt.Sprintf("${%s}", variable)
			content = bytes.Replace(content, []byte(enclosure), []byte(value), 1)

			newOccurrences := findAllOccurrences(content, []string{"$"})
			if len(startIndexes["$"]) > j + 1 && len(newOccurrences["$"]) > 0 {
				startIndexes["$"][j+1] = newOccurrences["$"][0]
			}
		}
	}
	return content, nil
}