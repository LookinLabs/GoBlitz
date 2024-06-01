package helper

import "os"

func CheckFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func CheckFileNotExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}
