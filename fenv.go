package fenv

import (
	"github.com/adammck/venv"
	venvMock "github.com/adammck/venv/mock"
	venvOs "github.com/adammck/venv/os"
	"github.com/lbrun25/fenv/domain"
	"github.com/lbrun25/fenv/mock"
	"github.com/lbrun25/fenv/os"
)

func OS() domain.Env {
	return os.NewOsEnv()
}

func Mock(env venv.Env) domain.Env {
	return mock.NewMockEnv(env)
}

// Auto returns the type of environment used
func Auto(env venv.Env) domain.Env {
	switch env.(type) {
	case *venvMock.MockEnv:
		return Mock(env)
	case *venvOs.OsEnv:
		return OS()
	default:
		return nil
	}
}