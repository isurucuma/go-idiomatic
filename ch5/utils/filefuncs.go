package utils

import (
	"fmt"
	"os"
)

func Filelen(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	defer func() { fmt.Println("Closing file"); file.Close() }()
	fileInfo, err := file.Stat()
	if err != nil {
		return -1, err
	}
	return int(fileInfo.Size()), nil
}
