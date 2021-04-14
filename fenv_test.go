package fenv_test

import (
	"github.com/adammck/venv"
	"github.com/lbrun25/fenv"
	"github.com/lbrun25/fenv/mock"
	"github.com/lbrun25/fenv/os"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS(t *testing.T) {
	assert.IsType(t, &os.Env{}, fenv.OS())
}

func TestMock(t *testing.T) {
	assert.IsType(t, &mock.Env{}, fenv.Mock(nil))
}

func TestAuto(t *testing.T) {
	t.Run("os", func(t *testing.T) {
		env := venv.OS()
		assert.IsType(t, &os.Env{}, fenv.Auto(env))
	})

	t.Run("mock", func(t *testing.T) {
		env := venv.Mock()
		assert.IsType(t, &mock.Env{}, fenv.Auto(env))
	})
}