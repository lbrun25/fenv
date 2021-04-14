package mock

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

func NewMockEnv(env venv.Env) domain.Env {
	return &Env{env: env}
}