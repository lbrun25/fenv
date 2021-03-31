package os

import (
	"github.com/adammck/venv"
	"github.com/lbrun25/fenv/domain"
	"github.com/lbrun25/fenv/parser"
)

type Env struct {
	env venv.Env
}

func (e Env) Parse(filePath string) ([]byte, error) {
	return parser.EnvironmentVariables(filePath, e.env)
}

func NewOsEnv() domain.Env {
	return &Env{env: venv.OS()}
}