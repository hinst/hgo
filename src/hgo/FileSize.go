package hgo

import "os"

func GetFileSize(filePath string) (result int64) {
	var fileInfo, statResult = os.Stat(filePath)
	if statResult == nil {
		result = fileInfo.Size()
	} else {
		result = -1
	}
	return
}
