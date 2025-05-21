package file

import (
	"os"
)

func CheckVideoExists(fileName string) error {
	_, err := os.ReadFile("videos/" + fileName)
	return err
}
