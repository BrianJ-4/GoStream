package file

import (
	"os"
)

func CheckFileExists(fileName string) error {
	_, err := os.Open("videos/" + fileName)
	return err
}

func OpenFile(fileName string) (*os.File, error) {
	file, err := os.Open("videos/" + fileName)
	if err != nil {
		return nil, err
	}
	return file, err
}

func GetFileLength(file *os.File) (int64, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	size := fileInfo.Size()
	return size, err
}
