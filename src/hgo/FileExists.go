package hgo

import "os"

func CheckFileExists(filePath string) bool {
	var _, err = os.Stat(filePath)
	return false == os.IsNotExist(err)
}
