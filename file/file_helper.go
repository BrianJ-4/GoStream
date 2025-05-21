package file

import (
	"io"
	"os"
	"path/filepath"
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

func GetFileSize(file *os.File) (int64, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}
	size := fileInfo.Size()
	return size, err
}

func GetFileExtension(file *os.File) string {
	return filepath.Ext(file.Name())
}

func GetData(w io.Writer, file *os.File, start int64, length int64) error {
	_, err := file.Seek(start, 0)
	if err != nil {
		return err
	}
	_, err = io.CopyN(w, file, length)
	return err
}
