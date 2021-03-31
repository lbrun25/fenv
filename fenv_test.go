package fenv_test

import (
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
	assert.IsType(t, &mock.Env{}, fenv.Mock())
}
