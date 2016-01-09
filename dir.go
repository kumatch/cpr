package cpr

import "os"

func makeDirectoryIfNotExists(dirname string) error {
	_, err := os.Lstat(dirname)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		return os.Mkdir(dirname, 0755)
	}

	return nil
}
