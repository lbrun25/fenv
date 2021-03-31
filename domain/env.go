package domain

type Env interface {
	// Parse parse the file and replace any environment variable by its value
	Parse(filePath string) ([]byte, error)
}