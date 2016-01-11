package cpr

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

var (
	pwd, _ = os.Getwd()

	filename          = "test-fixtures/basic/1.txt"
	dirname           = "test-fixtures"
	fileSymlink       = "test-fixtures/symlink/file-symlink"
	dirSymlink        = "test-fixtures/symlink/file-symlink"
	brokenFileSymlink = "test-fixtures/broken-symlink/file"
)

func TestFileElementRelativePath(t *testing.T) {
	e := createElement(filename, pwd)

	assert.True(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.False(t, e.IsSymLink())
	assert.Equal(t, filepath.Join(pwd, filename), e.Path())
	assert.Equal(t, filepath.Dir(filepath.Join(pwd, filename)), e.Dirname())
}

func TestFileElementFullPath(t *testing.T) {
	fullPath := filepath.Join(pwd, filename)
	e := createElement(fullPath, "/path/to/dummy")

	assert.True(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.False(t, e.IsSymLink())
	assert.Equal(t, fullPath, e.Path())
	assert.Equal(t, filepath.Dir(fullPath), e.Dirname())
}

func TestDirElementRelativePath(t *testing.T) {
	e := createElement(dirname, pwd)

	assert.True(t, e.IsExists())
	assert.True(t, e.IsDir())
	assert.False(t, e.IsSymLink())
	assert.Equal(t, filepath.Join(pwd, dirname), e.Path())
	assert.Equal(t, filepath.Dir(filepath.Join(pwd, dirname)), e.Dirname())
}

func TestDirElementFullPath(t *testing.T) {
	fullPath := filepath.Join(pwd, dirname)
	e := createElement(fullPath, "/path/to/dummy")

	assert.True(t, e.IsExists())
	assert.True(t, e.IsDir())
	assert.False(t, e.IsSymLink())
	assert.Equal(t, fullPath, e.Path())
	assert.Equal(t, filepath.Dir(fullPath), e.Dirname())
}

func TestFileSymlinkElementRelativePath(t *testing.T) {
	e := createElement(fileSymlink, pwd)

	assert.True(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.True(t, e.IsSymLink())
	assert.Equal(t, filepath.Join(pwd, fileSymlink), e.Path())
	assert.Equal(t, filepath.Dir(filepath.Join(pwd, fileSymlink)), e.Dirname())
}

func TestFileSymlinkElementFullPath(t *testing.T) {
	fullPath := filepath.Join(pwd, fileSymlink)
	e := createElement(fullPath, "/path/to/dummy")

	assert.True(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.True(t, e.IsSymLink())
	assert.Equal(t, fullPath, e.Path())
	assert.Equal(t, filepath.Dir(fullPath), e.Dirname())
}

func TestDirSymlinkElementRelativePath(t *testing.T) {
	e := createElement(dirSymlink, pwd)

	assert.True(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.True(t, e.IsSymLink())
	assert.Equal(t, filepath.Join(pwd, dirSymlink), e.Path())
	assert.Equal(t, filepath.Dir(filepath.Join(pwd, dirSymlink)), e.Dirname())
}

func TestDirSymlinkElementFullPath(t *testing.T) {
	fullPath := filepath.Join(pwd, dirSymlink)
	e := createElement(fullPath, "/path/to/dummy")

	assert.True(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.True(t, e.IsSymLink())
	assert.Equal(t, fullPath, e.Path())
	assert.Equal(t, filepath.Dir(fullPath), e.Dirname())
}

func TestNotExistsElementRelativePath(t *testing.T) {
	file := "invalid"
	e := createElement(file, pwd)

	assert.False(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.False(t, e.IsSymLink())
	assert.Equal(t, filepath.Join(pwd, file), e.Path())
	assert.Equal(t, pwd, e.Dirname())
}

func TestNotExistsElementFullPath(t *testing.T) {
	file := filepath.Join(pwd, "invalid")
	e := createElement(file, "/path/to/dummy")

	assert.False(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.False(t, e.IsSymLink())
	assert.Equal(t, file, e.Path())
	assert.Equal(t, pwd, e.Dirname())
}

func TestBrokenFileSymlinkElementRelativePath(t *testing.T) {
	e := createElement(brokenFileSymlink, pwd)

	assert.True(t, e.IsExists())
	assert.False(t, e.IsDir())
	assert.True(t, e.IsSymLink())
	assert.Equal(t, filepath.Join(pwd, brokenFileSymlink), e.Path())
	assert.Equal(t, filepath.Dir(filepath.Join(pwd, brokenFileSymlink)), e.Dirname())
}
