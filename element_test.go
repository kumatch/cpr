package cpr

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

var pwd, _ = os.Getwd()

func TestFileElementRelativePath(t *testing.T) {
	file := "element.go"
	e := createElement(file, pwd)

	assert.True(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.Equal(t, filepath.Join(pwd, file), e.Path())
	assert.Equal(t, pwd, e.Dirname())
}

func TestFileElementFullPath(t *testing.T) {
	file, _ := filepath.Abs("element.go")
	e := createElement(file, "/path/to/dummy")

	assert.True(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.Equal(t, file, e.Path())
	assert.Equal(t, pwd, e.Dirname())
}

func TestDirElementRelativePath(t *testing.T) {
	file := "."
	e := createElement(file, pwd)

	assert.True(t, e.IsExists())
	assert.True(t, e.IsDir())
	assert.Equal(t, pwd, e.Path())
	assert.Equal(t, filepath.Dir(pwd), e.Dirname())
}

func TestDirElementFullPath(t *testing.T) {
	e := createElement(pwd, "/path/to/dummy")

	assert.True(t, e.IsExists())
	assert.True(t, e.IsDir())
	assert.Equal(t, pwd, e.Path())
	assert.Equal(t, filepath.Dir(pwd), e.Dirname())
}

func TestNotExistsElementRelativePath(t *testing.T) {
	file := "invalid"
	e := createElement(file, pwd)

	assert.False(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.Equal(t, filepath.Join(pwd, file), e.Path())
	assert.Equal(t, pwd, e.Dirname())
}

func TestNotExistsElementFullPath(t *testing.T) {
	file := filepath.Join(pwd, "invalid")
	e := createElement(file, "/path/to/dummy")

	assert.False(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.Equal(t, file, e.Path())
	assert.Equal(t, pwd, e.Dirname())
}
