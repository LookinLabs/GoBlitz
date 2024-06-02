package helper

import "os"

func CheckIfFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
