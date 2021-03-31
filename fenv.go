package fenv

import (
	"github.com/lbrun25/fenv/domain"
	"github.com/lbrun25/fenv/mock"
	"github.com/lbrun25/fenv/os"
)

func OS() domain.Env {
	return os.NewOsEnv()
}

func Mock() domain.Env {
	return mock.NewMockEnv()
}
