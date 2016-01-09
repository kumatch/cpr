package cpr

import (
	"io"
	"os"
)

func copyFilename(source, dest string) error {
	sFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sFile.Close()

	dFile, err := os.Create(dest)
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
