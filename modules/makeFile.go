package modules

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

const path string = "./img"

func MakeFile(file multipart.File, name string) error {

	os.MkdirAll(path, 0777)
	filePath := fmt.Sprintf("%s/%s", path, name)

	newFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(newFile, file)
	if err != nil {
		return err
	}

	return nil

}
