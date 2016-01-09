package cpr

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	FileOverwrite = iota
)

func Copy(source, dest string, options ...int) error {
	basePath, _ := os.Getwd()
	sourceElement := createElement(source, basePath)
	destElement := createElement(dest, basePath)

	err := checkFirst(sourceElement, destElement)
	if err != nil {
		return err
	}

	if sourceElement.IsDir() {
		return copyOnDir(sourceElement, destElement, options)
	} else {
		return copyOnFile(sourceElement, destElement, options)
	}

	return nil
}

func checkFirst(source, dest *element) error {
	if !dest.IsExists() {
		return nil
	}

	if source.IsDir() && !dest.IsDir() {
		return fmt.Errorf("Invalid copy task: source = dir, dest = filename")
	}

	return nil
}

func parseOptions(options []int) (isOverwrite bool) {
	for _, option := range options {
		if option == FileOverwrite {
			isOverwrite = true
		}
	}

	return
}

func copyOnFile(source, dest *element, options []int) error {
	isOverwrite := parseOptions(options)

	if dest.IsDir() {
		return copyFilename(source.Path(), source.CreateCopyToPath(dest))
	}

	if !dest.IsExists() {
		return copyFilename(source.Path(), dest.Path())
	}

	if isOverwrite {
		return copyFilename(source.Path(), dest.Path())
	}

	return nil
}

func copyOnDir(source, dest *element, options []int) error {
	if !dest.IsExists() {
		err := os.Mkdir(dest.Path(), 0755)
		if err != nil {
			return err
		}
	}

	copyToSlice, err := createCopyToEntries(source, dest)
	if err != nil {
		return err
	}

	for _, c := range *copyToSlice {
		fmt.Printf("%s => %s\n", c.source, c.dest)
		childSource := createElement(c.source, "")
		childDest := createElement(c.dest, "")

		if c.isDir {
			err := copyOnDir(childSource, childDest, options)
			if err != nil {
				return err
			}
		} else {
			err := copyOnFile(childSource, childDest, options)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type copyTo struct {
	source string
	dest   string
	isDir  bool
}

func createCopyToEntries(source, dest *element) (*[]*copyTo, error) {
	files, _ := ioutil.ReadDir(source.Path())
	copyToSlice := make([]*copyTo, 0, len(files))

	for _, file := range files {
		sourcePath := filepath.Join(source.Path(), file.Name())
		stat, err := os.Lstat(sourcePath)
		if err != nil {
			return nil, err
		}

		copyToSlice = append(copyToSlice, &copyTo{
			source: sourcePath,
			dest:   filepath.Join(dest.Path(), file.Name()),
			isDir:  stat.IsDir(),
		})
	}

	return &copyToSlice, nil
}

/*
func isWritable(filename string) bool {
	_, err := os.Lstat(filename)
	if err != nil {
		return false
	}

	return true
}
*/
