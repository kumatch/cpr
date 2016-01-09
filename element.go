package cpr

import (
	"os"
	"path/filepath"
)

type element struct {
	targetPath string
	basePath   string
	fileInfo   *os.FileInfo
}

func (e *element) Path() string {
	if filepath.IsAbs(e.targetPath) {
		return e.targetPath
	}

	return filepath.Join(e.basePath, e.targetPath)
}

func (e *element) IsExists() bool {
	stat, _ := e.stat()

	return stat != nil
}

func (e *element) IsDir() bool {
	stat, _ := e.stat()

	return stat != nil && (*stat).IsDir()
}

func (e *element) Dirname() string {
	return filepath.Dir(e.Path())
}

func (e *element) Basename() string {
	return filepath.Base(e.Path())
}

func (e *element) CreateCopyToPath(copyToElement *element) string {
	if copyToElement.IsDir() {
		return filepath.Join(copyToElement.Path(), e.Basename())
	}

	return filepath.Join(copyToElement.Dirname(), e.Basename())
}

func (e *element) stat() (*os.FileInfo, error) {
	if e.fileInfo != nil {
		return e.fileInfo, nil
	}

	fileInfo, err := os.Lstat(e.Path())
	if err != nil {
		return nil, err
	}

	e.fileInfo = &fileInfo

	return e.fileInfo, nil
}

func createElement(targetPath string, basePath string) *element {
	return &element{targetPath: targetPath, basePath: basePath}
}
