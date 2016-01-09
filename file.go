package cpr

import (
	"io"
	"os"
)

func copyFilename(source, dest *element) error {
	if source.IsSymLink() {
		return _createSymLink(source, dest)
	}

	sFile, err := os.Open(source.Path())
	if err != nil {
		return err
	}
	defer sFile.Close()

	dFile, err := os.Create(dest.Path())
	if err != nil {
		return err
	}
	defer dFile.Close()

	return _copyFile(sFile, dFile)
}

func _copyFile(sourceFile, destFile *os.File) error {
	_, err := io.Copy(destFile, sourceFile)

	return err
}

func _createSymLink(source, dest *element) error {
	path, err := os.Readlink(source.Path())
	if err != nil {
		return err
	}

	return os.Symlink(path, dest.Path())
}
